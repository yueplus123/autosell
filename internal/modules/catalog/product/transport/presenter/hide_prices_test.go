package presenter

import (
	"testing"

	"github.com/dujiao-next/internal/shared/money"

	"github.com/shopspring/decimal"
)

func TestProductHidePricesClearsAllPriceFields(t *testing.T) {
	promo := money.FromDecimal(decimal.NewFromInt(8))
	member := money.FromDecimal(decimal.NewFromInt(7))
	id := uint(3)
	product := Product{
		PriceAmount:          money.FromDecimal(decimal.NewFromInt(10)),
		WholesalePrices:      []WholesalePrice{{MinQuantity: 5, UnitPrice: "9.00"}},
		PromotionID:          &id,
		PromotionName:        "sale",
		PromotionPriceAmount: &promo,
		PromotionRules:       []PromotionRule{{Type: "percent"}},
		MemberPrices:         []MemberLevelPrice{{}},
		SKUs: []SKU{{
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
