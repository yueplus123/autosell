package producthttp

import (
	"errors"
	"testing"

	productcontract "github.com/dujiao-next/internal/modules/catalog/product/contract"
	productdomain "github.com/dujiao-next/internal/modules/catalog/product/domain"

	productpresenter "github.com/dujiao-next/internal/modules/catalog/product/transport/presenter"
	reseller "github.com/dujiao-next/internal/modules/reseller/contract"
	"github.com/dujiao-next/internal/shared/money"
	"github.com/shopspring/decimal"
)

type stubResellerDisplayPricer struct {
	result *reseller.DisplayPriceResult
	err    error
}

func (s stubResellerDisplayPricer) LoadDisplayPricingBatch(tenant reseller.TenantContext, products []productdomain.Product) (*reseller.DisplayPricingBatch, error) {
	return &reseller.DisplayPricingBatch{Tenant: tenant}, nil
}

func (s stubResellerDisplayPricer) ResolveDisplayPrices(tenant reseller.TenantContext, product *productdomain.Product, batch *reseller.DisplayPricingBatch) (*reseller.DisplayPriceResult, error) {
	return s.result, s.err
}

func TestDecoratePublicProductDisplayPricePrefersFirstActiveSKU(t *testing.T) {
	h := &PublicHandler{}
	product := &productdomain.Product{
		ID:          1,
		PriceAmount: money.FromDecimal(decimal.RequireFromString("59.90")),
		SKUs: []productdomain.ProductSKU{
			{
				ID:          11,
				IsActive:    true,
				SortOrder:   100,
				PriceAmount: money.FromDecimal(decimal.RequireFromString("89.90")),
			},
			{
				ID:          12,
				IsActive:    true,
				SortOrder:   10,
				PriceAmount: money.FromDecimal(decimal.RequireFromString("49.90")),
			},
		},
	}

	item, err := h.decoratePublicProduct(product, nil)
	if err != nil {
		t.Fatalf("decoratePublicProduct failed: %v", err)
	}
	expected := decimal.RequireFromString("89.90")
	if !item.PriceAmount.Decimal.Equal(expected) {
		t.Fatalf("expected display price %s, got: %s", expected.String(), item.PriceAmount.String())
	}
}

func TestDecoratePublicProductForTenantUsesResellerPricesAndHidesMainDiscounts(t *testing.T) {
	product := &productdomain.Product{
		ID:          1,
		Slug:        "reseller-display",
		PriceAmount: money.FromDecimal(decimal.NewFromInt(100)),
		WholesalePrices: productdomain.WholesalePriceTiers{
			{MinQuantity: 5, UnitPrice: money.FromDecimal(decimal.NewFromInt(80))},
		},
		SKUs: []productdomain.ProductSKU{
			{
				ID:          11,
				ProductID:   1,
				SKUCode:     "VISIBLE",
				IsActive:    true,
				SortOrder:   100,
				PriceAmount: money.FromDecimal(decimal.NewFromInt(100)),
			},
			{
				ID:          12,
				ProductID:   1,
				SKUCode:     "HIDDEN",
				IsActive:    true,
				SortOrder:   10,
				PriceAmount: money.FromDecimal(decimal.NewFromInt(90)),
			},
		},
	}
	h := &PublicHandler{pricer: stubResellerDisplayPricer{result: &reseller.DisplayPriceResult{
		Visible:      true,
		ProductID:    1,
		DisplaySKUID: 11,
		DisplayPrice: money.FromDecimal(decimal.NewFromInt(130)),
		SKUPrices: map[uint]money.Amount{
			11: money.FromDecimal(decimal.NewFromInt(130)),
		},
		HiddenSKUIDs: map[uint]bool{12: true},
	}}}
	tenant := reseller.ResellerTenantContext("shop.example.test", 10, 99, "shop.example.test")

	item, err := h.decoratePublicProductForTenant(product, nil, tenant, &reseller.DisplayPricingBatch{})
	if err != nil {
		t.Fatalf("decoratePublicProductForTenant failed: %v", err)
	}
	if !item.PriceAmount.Decimal.Equal(decimal.NewFromInt(130)) {
		t.Fatalf("expected reseller product price 130, got %s", item.PriceAmount.String())
	}
	if len(item.SKUs) != 1 || item.SKUs[0].ID != 11 {
		t.Fatalf("expected only visible sku 11, got %+v", item.SKUs)
	}
	if !item.SKUs[0].PriceAmount.Decimal.Equal(decimal.NewFromInt(130)) {
		t.Fatalf("expected reseller sku price 130, got %s", item.SKUs[0].PriceAmount.String())
	}
	if item.PromotionPriceAmount != nil || item.PromotionID != nil || len(item.PromotionRules) > 0 {
		t.Fatalf("reseller display must not expose main promotion fields: %+v", item)
	}
	if len(item.WholesalePrices) > 0 {
		t.Fatalf("reseller display must not expose main wholesale prices: %+v", item.WholesalePrices)
	}
}

func TestDecoratePublicProductForTenantHiddenProduct(t *testing.T) {
	h := &PublicHandler{pricer: stubResellerDisplayPricer{result: &reseller.DisplayPriceResult{Visible: false}}}
	tenant := reseller.ResellerTenantContext("shop.example.test", 10, 99, "shop.example.test")
	product := &productdomain.Product{
		ID:          1,
		PriceAmount: money.FromDecimal(decimal.NewFromInt(100)),
		SKUs: []productdomain.ProductSKU{
			{ID: 11, ProductID: 1, IsActive: true, PriceAmount: money.FromDecimal(decimal.NewFromInt(100))},
		},
	}

	_, err := h.decoratePublicProductForTenant(product, nil, tenant, &reseller.DisplayPricingBatch{})
	if !errors.Is(err, productcontract.ErrResellerProductNotListed) {
		t.Fatalf("expected ErrResellerProductNotListed, got %v", err)
	}
}

func TestDecoratePublicProductForTenantAllHiddenSKUsReturnsNotListed(t *testing.T) {
	h := &PublicHandler{pricer: stubResellerDisplayPricer{result: &reseller.DisplayPriceResult{
		Visible:      true,
		DisplayPrice: money.FromDecimal(decimal.NewFromInt(100)),
		HiddenSKUIDs: map[uint]bool{11: true, 12: true},
	}}}
	tenant := reseller.ResellerTenantContext("shop.example.test", 10, 99, "shop.example.test")
	product := &productdomain.Product{
		ID:          1,
		PriceAmount: money.FromDecimal(decimal.NewFromInt(100)),
		SKUs: []productdomain.ProductSKU{
			{ID: 11, ProductID: 1, IsActive: true, PriceAmount: money.FromDecimal(decimal.NewFromInt(100))},
			{ID: 12, ProductID: 1, IsActive: true, PriceAmount: money.FromDecimal(decimal.NewFromInt(120))},
		},
	}

	_, err := h.decoratePublicProductForTenant(product, nil, tenant, &reseller.DisplayPricingBatch{})
	if !errors.Is(err, productcontract.ErrResellerProductNotListed) {
		t.Fatalf("expected ErrResellerProductNotListed, got %v", err)
	}
}

func TestDecoratePublicProductForTenantInvalidDisplayPricingIsHidden(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{name: "below base", err: reseller.ErrPriceBelowBase},
		{name: "markup exceeds max", err: reseller.ErrMarkupExceeded},
		{name: "unknown pricing mode", err: reseller.ErrPricingModeInvalid},
		{name: "already not listed", err: productcontract.ErrResellerProductNotListed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PublicHandler{pricer: stubResellerDisplayPricer{err: tt.err}}
			tenant := reseller.ResellerTenantContext("shop.example.test", 10, 99, "shop.example.test")
			product := &productdomain.Product{
				ID:          1,
				PriceAmount: money.FromDecimal(decimal.NewFromInt(100)),
				SKUs: []productdomain.ProductSKU{
					{ID: 11, ProductID: 1, IsActive: true, PriceAmount: money.FromDecimal(decimal.NewFromInt(100))},
				},
			}
			_, err := h.decoratePublicProductForTenant(product, nil, tenant, &reseller.DisplayPricingBatch{})
			if !errors.Is(err, productcontract.ErrResellerProductNotListed) {
				t.Fatalf("expected ErrResellerProductNotListed, got %v", err)
			}
		})
	}
}

func TestProductHidePricesClearsAllPriceFields(t *testing.T) {
	promo := money.FromDecimal(decimal.NewFromInt(8))
	member := money.FromDecimal(decimal.NewFromInt(7))
	id := uint(3)
	product := productpresenter.Product{
		PriceAmount:          money.FromDecimal(decimal.NewFromInt(10)),
		WholesalePrices:      []productpresenter.WholesalePrice{{MinQuantity: 5, UnitPrice: "9.00"}},
		PromotionID:          &id,
		PromotionName:        "sale",
		PromotionPriceAmount: &promo,
		PromotionRules:       []productpresenter.PromotionRule{{Type: "percent"}},
		MemberPrices:         []productpresenter.MemberLevelPrice{{}},
		SKUs: []productpresenter.SKU{{
			PriceAmount:          money.FromDecimal(decimal.NewFromInt(10)),
			PromotionPriceAmount: &promo,
			MemberPriceAmount:    &member,
		}},
	}
	product.HidePrices()
	if !product.PriceAmount.IsZero() || product.WholesalePrices != nil || product.PromotionID != nil ||
		product.PromotionName != "" || product.PromotionPriceAmount != nil || product.PromotionRules != nil ||
		product.MemberPrices != nil {
		t.Fatalf("product price fields not cleared: %+v", product)
	}
	sku := product.SKUs[0]
	if !sku.PriceAmount.IsZero() || sku.PromotionPriceAmount != nil || sku.MemberPriceAmount != nil {
		t.Fatalf("sku price fields not cleared: %+v", sku)
	}
	if !product.PriceHidden {
		t.Fatal("PriceHidden should be true")
	}
}

type stubShowcase bool

func (s stubShowcase) IsShowcaseMode() bool { return bool(s) }

func TestShouldHidePricesOnlyOnMainTenantInShowcaseMode(t *testing.T) {
	h := &PublicHandler{}
	main := reseller.MainTenantContext("shop.example.com")
	resellerTenant := reseller.ResellerTenantContext("a.example.com", 1, 2, "a.example.com")
	if h.shouldHidePrices(main) {
		t.Fatal("no showcase reader: prices must stay visible")
	}
	h.SetShowcaseModeReader(stubShowcase(true))
	if !h.shouldHidePrices(main) {
		t.Fatal("main tenant in showcase mode must hide prices")
	}
	if h.shouldHidePrices(resellerTenant) {
		t.Fatal("reseller tenant must never hide prices")
	}
	h.SetShowcaseModeReader(stubShowcase(false))
	if h.shouldHidePrices(main) {
		t.Fatal("showcase off: prices must stay visible")
	}
}
