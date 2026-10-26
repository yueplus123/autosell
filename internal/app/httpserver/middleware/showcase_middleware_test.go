package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	resellercontract "github.com/dujiao-next/internal/modules/reseller/contract"

	"github.com/gin-gonic/gin"
)

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

func TestResellerContentIsolation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	newRouter := func(tenant resellercontract.TenantContext) *gin.Engine {
		r := gin.New()
		r.Use(func(c *gin.Context) {
			c.Request = c.Request.WithContext(resellercontract.WithTenantContext(c.Request.Context(), tenant))
			c.Next()
		})
		r.Use(ResellerContentIsolation())
		r.GET("/api/v1/public/posts", func(c *gin.Context) { c.String(http.StatusOK, "real-list") })
		r.GET("/api/v1/public/posts/:slug", func(c *gin.Context) { c.String(http.StatusOK, "real-detail") })
		r.GET("/api/v1/public/products", func(c *gin.Context) { c.String(http.StatusOK, "products") })
		return r
	}
	get := func(r *gin.Engine, path string) string {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
		return w.Body.String()
	}

	main := newRouter(resellercontract.MainTenantContext("shop.example.com"))
	if got := get(main, "/api/v1/public/posts"); got != "real-list" {
		t.Fatalf("main site list should pass through, got %q", got)
	}
	if got := get(main, "/api/v1/public/posts/hello"); got != "real-detail" {
		t.Fatalf("main site detail should pass through, got %q", got)
	}

	reseller := newRouter(resellercontract.ResellerTenantContext("a.example.com", 1, 2, "a.example.com"))
	if got := get(reseller, "/api/v1/public/posts?page=1&page_size=3"); !strings.Contains(got, `"data":[]`) || !strings.Contains(got, `"total":0`) {
		t.Fatalf("reseller list should be empty, got %q", got)
	}
	if got := get(reseller, "/api/v1/public/posts/hello"); !strings.Contains(got, `"status_code":404`) {
		t.Fatalf("reseller detail should be not found, got %q", got)
	}
	if got := get(reseller, "/api/v1/public/products"); got != "products" {
		t.Fatalf("other routes must pass through, got %q", got)
	}
}
