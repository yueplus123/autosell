package integrationtest

import (
	"fmt"
	"sort"
	"testing"
	"time"

	categorydomain "github.com/dujiao-next/internal/modules/catalog/category/domain"
	productdomain "github.com/dujiao-next/internal/modules/catalog/product/domain"
	resellerdomain "github.com/dujiao-next/internal/modules/reseller/domain"
	resellergormstore "github.com/dujiao-next/internal/modules/reseller/infrastructure/gormstore"
	"github.com/dujiao-next/internal/shared/jsonmap"
	"github.com/dujiao-next/internal/shared/money"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func TestNormalizeZoneScope(t *testing.T) {
	cases := map[string]struct {
		want string
		ok   bool
	}{
		"":       {resellerdomain.ZoneScopeAll, true},
		" ALL ":  {resellerdomain.ZoneScopeAll, true},
		"ai":     {resellerdomain.ZoneScopeAI, true},
		"Game":   {resellerdomain.ZoneScopeGame, true},
		"anime":  {"", false},
		"ai,all": {"", false},
	}
	for raw, tc := range cases {
		got, ok := resellerdomain.NormalizeZoneScope(raw)
		if got != tc.want || ok != tc.ok {
			t.Fatalf("NormalizeZoneScope(%q) = %q,%v want %q,%v", raw, got, ok, tc.want, tc.ok)
		}
	}
	var nilProfile *resellerdomain.Profile
	if nilProfile.EffectiveZoneScope() != resellerdomain.ZoneScopeAll {
		t.Fatalf("nil profile should default to all")
	}
}

type zoneFixture struct {
	aiProduct, aiChildProduct, gameProduct, otherProduct productdomain.Product
}

func seedZoneCatalog(t *testing.T, db *gorm.DB) zoneFixture {
	t.Helper()
	mkCat := func(slug string, parent uint) categorydomain.Category {
		c := categorydomain.Category{Slug: slug, ParentID: parent, NameJSON: jsonmap.JSON{"zh-CN": slug}, IsActive: true}
		if err := db.Create(&c).Error; err != nil {
			t.Fatalf("create category %s: %v", slug, err)
		}
		return c
	}
	mkProduct := func(slug string, categoryID uint) productdomain.Product {
		p := productdomain.Product{
			CategoryID:  categoryID,
			Slug:        slug,
			TitleJSON:   jsonmap.JSON{"zh-CN": slug},
			PriceAmount: money.FromDecimal(decimal.RequireFromString("10.00")),
			IsActive:    true,
		}
		if err := db.Create(&p).Error; err != nil {
			t.Fatalf("create product %s: %v", slug, err)
		}
		return p
	}
	ai := mkCat("ai", 0)
	chatgpt := mkCat("chatgpt", ai.ID)
	game := mkCat("game-topup", 0)
	genshin := mkCat("genshin-impact", game.ID)
	other := mkCat("misc", 0)
	return zoneFixture{
		aiProduct:      mkProduct("ai-root-product", ai.ID),
		aiChildProduct: mkProduct("chatgpt-plus", chatgpt.ID),
		gameProduct:    mkProduct("genshin-6480", genshin.ID),
		otherProduct:   mkProduct("misc-product", other.ID),
	}
}

func sortedIDs(ids []uint) []uint {
	out := append([]uint(nil), ids...)
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

func TestResellerZoneScopeHidesProductsOutsideZone(t *testing.T) {
	db := openResellerProductSettingServiceTestDB(t)
	store := resellergormstore.New(db)
	_, profile, _, _ := seedResellerProductSettingServiceData(t, db)
	// 种子数据自带一个不在任何专区的商品，这里另建专区目录
	var seededProducts []uint
	if err := db.Model(&productdomain.Product{}).Pluck("id", &seededProducts).Error; err != nil {
		t.Fatalf("pluck seeded: %v", err)
	}
	fx := seedZoneCatalog(t, db)

	// 默认（历史数据）= 全部专区，不隐藏任何商品
	ids, err := store.ListOutOfZoneProductIDs(profile.ID)
	if err != nil {
		t.Fatalf("list out of zone: %v", err)
	}
	if len(ids) != 0 {
		t.Fatalf("all-zone reseller should hide nothing, got %v", ids)
	}

	// 只代理 AI：二次元、其他分类商品都隐藏，AI 一级与子分类商品可见
	if err := db.Model(&resellerdomain.Profile{}).Where("id = ?", profile.ID).Update("zone_scope", resellerdomain.ZoneScopeAI).Error; err != nil {
		t.Fatalf("update scope: %v", err)
	}
	ids, err = store.ListOutOfZoneProductIDs(profile.ID)
	if err != nil {
		t.Fatalf("list out of zone: %v", err)
	}
	want := sortedIDs(append([]uint{fx.gameProduct.ID, fx.otherProduct.ID}, seededProducts...))
	if fmt.Sprint(sortedIDs(ids)) != fmt.Sprint(want) {
		t.Fatalf("ai reseller hidden = %v, want %v", sortedIDs(ids), want)
	}

	hidden, err := store.ListHiddenProductIDs(profile.ID)
	if err != nil {
		t.Fatalf("list hidden: %v", err)
	}
	hiddenSet := map[uint]bool{}
	for _, id := range hidden {
		hiddenSet[id] = true
	}
	if !hiddenSet[fx.gameProduct.ID] || hiddenSet[fx.aiProduct.ID] || hiddenSet[fx.aiChildProduct.ID] {
		t.Fatalf("hidden list must include out-of-zone and exclude in-zone products: %v", hidden)
	}

	// 只代理二次元
	if err := db.Model(&resellerdomain.Profile{}).Where("id = ?", profile.ID).Update("zone_scope", resellerdomain.ZoneScopeGame).Error; err != nil {
		t.Fatalf("update scope: %v", err)
	}
	ids, err = store.ListOutOfZoneProductIDs(profile.ID)
	if err != nil {
		t.Fatalf("list out of zone: %v", err)
	}
	for _, id := range ids {
		if id == fx.gameProduct.ID {
			t.Fatalf("game product must stay visible for game reseller")
		}
	}
	if len(ids) != len(seededProducts)+3 {
		t.Fatalf("game reseller should hide ai(2) + other(1) + seeded, got %v", ids)
	}
}

func TestResellerApplyAndApproveZoneScope(t *testing.T) {
	db := openResellerManagementServiceTestDB(t)
	svc := newResellerManagementServiceForTest(db)
	user := seedResellerManagementUser(t, db, fmt.Sprintf("zone-%d@example.test", time.Now().UnixNano())).ID

	if _, err := svc.ApplyUserReseller(user, ResellerApplyInput{Reason: "x", ZoneScope: "bogus"}); err == nil {
		t.Fatalf("invalid zone scope must be rejected")
	}
	profile, err := svc.ApplyUserReseller(user, ResellerApplyInput{Reason: "only anime", ZoneScope: "game"})
	if err != nil {
		t.Fatalf("apply: %v", err)
	}
	if profile.EffectiveZoneScope() != resellerdomain.ZoneScopeGame {
		t.Fatalf("apply should store requested scope, got %q", profile.ZoneScope)
	}
}
