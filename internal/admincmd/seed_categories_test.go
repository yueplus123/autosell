package admincmd

import (
	"fmt"
	"testing"
	"time"

	categorydomain "github.com/dujiao-next/internal/modules/catalog/category/domain"
	"github.com/dujiao-next/internal/shared/jsonmap"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func newSeedTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	dsn := fmt.Sprintf("file:seed_categories_%d?mode=memory&cache=shared", time.Now().UnixNano())
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&categorydomain.Category{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func countSeedSlugs(seeds []seedCategory) int {
	n := 0
	for _, s := range seeds {
		n += 1 + len(s.Children)
	}
	return n
}

func TestSeedBusinessCategoriesCreatesTreeAndIsIdempotent(t *testing.T) {
	db := newSeedTestDB(t)
	total := countSeedSlugs(businessCategorySeed)

	first, err := seedBusinessCategories(db, businessCategorySeed)
	if err != nil {
		t.Fatalf("first seed: %v", err)
	}
	if len(first.Created) != total || len(first.Skipped) != 0 {
		t.Fatalf("first run: created=%v skipped=%v", first.Created, first.Skipped)
	}

	var ai, chatgpt, game, genshin categorydomain.Category
	db.Where("slug = ?", "ai").First(&ai)
	db.Where("slug = ?", "chatgpt").First(&chatgpt)
	db.Where("slug = ?", "game-topup").First(&game)
	db.Where("slug = ?", "genshin-impact").First(&genshin)
	if ai.ParentID != 0 || game.ParentID != 0 {
		t.Fatalf("roots must be top-level: ai=%d game=%d", ai.ParentID, game.ParentID)
	}
	if chatgpt.ParentID != ai.ID || genshin.ParentID != game.ID {
		t.Fatalf("children attached to wrong parents")
	}
	if !ai.IsActive || ai.NameJSON["zh-CN"] != "AI 服务" {
		t.Fatalf("unexpected ai row: %+v", ai)
	}

	second, err := seedBusinessCategories(db, businessCategorySeed)
	if err != nil {
		t.Fatalf("second seed: %v", err)
	}
	if len(second.Created) != 0 || len(second.Skipped) != total {
		t.Fatalf("second run should skip all: %+v", second)
	}
	var count int64
	db.Model(&categorydomain.Category{}).Count(&count)
	if int(count) != total {
		t.Fatalf("row count = %d, want %d", count, total)
	}
}

func TestSeedBusinessCategoriesKeepsAdminEditsAndRestoresDeleted(t *testing.T) {
	db := newSeedTestDB(t)
	deletedAt := time.Now()
	edited := categorydomain.Category{Slug: "ai", NameJSON: jsonmap.JSON{"zh-CN": "人工智能"}, SortOrder: 999, IsActive: true}
	deleted := categorydomain.Category{Slug: "game-topup", NameJSON: jsonmap.JSON{"zh-CN": "旧"}, IsActive: false, DeletedAt: &deletedAt}
	db.Create(&edited)
	db.Create(&deleted)

	result, err := seedBusinessCategories(db, businessCategorySeed)
	if err != nil {
		t.Fatalf("seed: %v", err)
	}
	if len(result.Restored) != 1 || result.Restored[0] != "game-topup" {
		t.Fatalf("restored = %v", result.Restored)
	}

	var ai, game categorydomain.Category
	db.Where("slug = ?", "ai").First(&ai)
	db.Where("slug = ?", "game-topup").First(&game)
	if ai.NameJSON["zh-CN"] != "人工智能" || ai.SortOrder != 999 {
		t.Fatalf("admin edits overwritten: %+v", ai)
	}
	if game.DeletedAt != nil || !game.IsActive || game.ParentID != 0 {
		t.Fatalf("deleted root not restored: %+v", game)
	}
}

func TestSeedBusinessCategoriesWarnsWhenRootIsNested(t *testing.T) {
	db := newSeedTestDB(t)
	parent := categorydomain.Category{Slug: "misc", NameJSON: jsonmap.JSON{"zh-CN": "杂项"}, IsActive: true}
	db.Create(&parent)
	db.Create(&categorydomain.Category{Slug: "ai", ParentID: parent.ID, NameJSON: jsonmap.JSON{"zh-CN": "AI"}, IsActive: true})

	result, err := seedBusinessCategories(db, businessCategorySeed)
	if err != nil {
		t.Fatalf("seed: %v", err)
	}
	if len(result.Warnings) == 0 {
		t.Fatalf("expected a warning for nested root")
	}
}
