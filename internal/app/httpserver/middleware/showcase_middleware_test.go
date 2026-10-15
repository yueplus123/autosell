package middleware

import "testing"

func TestIsShowcaseBlockedRoute(t *testing.T) {
	cases := []struct {
		method, path string
		want         bool
	}{
		{"POST", "/api/v1/cart/items", true},
		{"POST", "/api/v1/orders", true},
		{"POST", "/api/v1/orders/preview", true},
		{"POST", "/api/v1/orders/create-and-pay", true},
		{"POST", "/api/v1/guest/orders", true},
		{"POST", "/api/v1/guest/orders/preview", true},
		{"POST", "/api/v1/guest/orders/create-and-pay", true},
		{"GET", "/api/v1/orders", false},
		{"GET", "/api/v1/guest/orders", false},
		{"POST", "/api/v1/orders/:order_no/cancel", false},
		{"POST", "/api/v1/order/payment-channels", false},
		{"DELETE", "/api/v1/cart/items/:product_id", false},
		{"POST", "/api/v1/reseller/orders/export", false},
		{"POST", "", false},
	}
	for _, tc := range cases {
		if got := isShowcaseBlockedRoute(tc.method, tc.path); got != tc.want {
			t.Errorf("%s %s: got %v want %v", tc.method, tc.path, got, tc.want)
		}
	}
}
