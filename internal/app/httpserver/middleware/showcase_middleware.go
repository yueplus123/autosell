package middleware

import (
	"net/http"
	"strings"

	"github.com/dujiao-next/internal/i18n"
	resellercontract "github.com/dujiao-next/internal/modules/reseller/contract"
	"github.com/dujiao-next/internal/platform/http/response"

	"github.com/gin-gonic/gin"
)

// ShowcaseModeReader 读取主站展示模式开关。
type ShowcaseModeReader interface {
	IsShowcaseMode() bool
}

// showcaseBlockedSuffixes 是主站展示模式下拒绝的下单类接口（均为 POST）：
// 加入购物车、订单预览、创建订单、创建并支付（用户与游客）。
var showcaseBlockedSuffixes = []string{
	"/cart/items",
	"/orders/preview",
	"/orders",
	"/orders/create-and-pay",
}

func isShowcaseBlockedRoute(method, fullPath string) bool {
	if method != http.MethodPost || fullPath == "" {
		return false
	}
	for _, suffix := range showcaseBlockedSuffixes {
		if strings.HasSuffix(fullPath, suffix) {
			return true
		}
	}
	return false
}

// ShowcaseModeOrderGuard 在主站展示模式下拒绝主站的购物车与下单请求；分销站不受影响。
func ShowcaseModeOrderGuard(reader ShowcaseModeReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		if reader == nil || !isShowcaseBlockedRoute(c.Request.Method, c.FullPath()) {
			c.Next()
			return
		}
		tenant, ok := resellercontract.TenantFromContext(c.Request.Context())
		if ok && tenant.IsReseller() {
			c.Next()
			return
		}
		if !reader.IsShowcaseMode() {
			c.Next()
			return
		}
		response.Forbidden(c, i18n.T(i18n.ResolveLocale(c), "error.showcase_mode"))
		c.Abort()
	}
}
