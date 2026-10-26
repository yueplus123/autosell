package admincmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dujiao-next/internal/constants"
	contentdomain "github.com/dujiao-next/internal/modules/content/domain"
	"github.com/dujiao-next/internal/platform/database/gormdb"
	"github.com/dujiao-next/internal/shared/jsonmap"

	"gorm.io/gorm"
)

// demoPostSlugPrefix 示例文章的 slug 前缀，删除示例文章时只按这个前缀删，不会碰到真实文章。
const demoPostSlugPrefix = "demo-"

type demoPost struct {
	Slug    string
	Title   string
	Summary string
	Content string
	Cover   string
	DaysAgo int
}

// demoPosts 用来预览官网首页文章列表的样式，内容为示例。
var demoPosts = []demoPost{
	{
		Slug: "demo-chatgpt-plus-vs-pro", Cover: "/xinhuo/demo/cover-1.svg", DaysAgo: 1,
		Title:   "ChatGPT Plus 和 Pro 有什么区别？一篇看懂怎么选",
		Summary: "日常写作、翻译选 Plus 就够；重度使用、需要更高额度和更强模型再考虑 Pro。",
		Content: "<p>很多朋友第一次开会员时都会纠结 Plus 和 Pro 选哪个。简单说：</p><ul><li><b>Plus</b>：适合日常写作、翻译、学习，性价比最高。</li><li><b>Pro</b>：额度更高，适合每天大量使用或做开发的用户。</li></ul><p>如果拿不准，可以先开一个月 Plus，用得不够再升级。</p>",
	},
	{
		Slug: "demo-apple-id-notice", Cover: "/xinhuo/demo/cover-2.svg", DaysAgo: 3,
		Title:   "海外 Apple ID 使用须知：只在 App Store 登录",
		Summary: "海外账号只用来下载应用和订阅，千万不要在设置里登录 iCloud，否则可能锁机。",
		Content: "<p>拿到海外 Apple ID 后，请按下面的方式使用：</p><ol><li>打开 App Store，点右上角头像，退出原账号。</li><li>登录新的海外 Apple ID。</li><li>下载完应用后可以切回自己的账号。</li></ol><p><b>注意：</b>不要在「设置」里登录 iCloud。</p>",
	},
	{
		Slug: "demo-visa-subscribe", Cover: "/xinhuo/demo/cover-3.svg", DaysAgo: 6,
		Title:   "用虚拟 Visa 卡订阅海外服务的完整流程",
		Summary: "从拿到卡号到绑卡扣款，三步完成 ChatGPT、Claude 等服务的订阅。",
		Content: "<p>虚拟 Visa 卡适合一次性订阅海外服务。流程如下：</p><ol><li>收到卡号、有效期和安全码。</li><li>在订阅页面填写卡片信息，账单地址填卡片对应地址。</li><li>确认扣款成功后即可使用。</li></ol><p>同一张卡建议只绑定一个平台。</p>",
	},
	{
		Slug: "demo-claude-pro-max", Cover: "/xinhuo/demo/cover-4.svg", DaysAgo: 9,
		Title:   "Claude Pro 和 Max 适合哪些人？",
		Summary: "写长文、读文档用 Pro；需要整天高强度使用的团队和开发者可以看看 Max。",
		Content: "<p>Claude 擅长长文本和代码。选择建议：</p><ul><li><b>Pro</b>：个人日常使用足够。</li><li><b>Max</b>：额度更高，适合重度用户。</li></ul><p>首次登录建议使用稳定的网络环境。</p>",
	},
	{
		Slug: "demo-acg-topup-faq", Cover: "/xinhuo/demo/cover-5.svg", DaysAgo: 13,
		Title:   "原神、星穹铁道代充常见问题汇总",
		Summary: "代充要提供什么、多久到账、首充双倍怎么算，这里一次说清楚。",
		Content: "<p>整理了大家最常问的几个问题：</p><ul><li><b>需要提供什么？</b> UID、服务器和登录方式。</li><li><b>多久到账？</b> 一般 10～30 分钟。</li><li><b>首充双倍？</b> 以游戏内显示为准。</li></ul><p>代充期间请不要登录游戏。</p>",
	},
	{
		Slug: "demo-fanbox-support", Cover: "/xinhuo/demo/cover-6.svg", DaysAgo: 18,
		Title:   "Fanbox 代赞助是什么？怎么下单",
		Summary: "支持喜欢的画师又没有海外支付方式？代赞助帮你按方案按月赞助。",
		Content: "<p>Fanbox 代赞助就是由我们代为赞助你喜欢的创作者。</p><ol><li>发送创作者主页链接。</li><li>选择要赞助的方案和月数。</li><li>赞助成功后按约定方式交付。</li></ol><p>部分创作者禁止转载，请遵守创作者的规定。</p>",
	},
}

func demoLocalized(text string) jsonmap.JSON {
	return jsonmap.JSON{"zh-CN": text, "zh-TW": text, "en-US": text}
}

// seedDemoPosts 幂等地写入示例文章：slug 已存在（含已删除）则跳过。
func seedDemoPosts(db *gorm.DB, posts []demoPost, now time.Time) (created, skipped int, err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, p := range posts {
			if !strings.HasPrefix(p.Slug, demoPostSlugPrefix) {
				return fmt.Errorf("示例文章 slug 必须以 %q 开头: %s", demoPostSlugPrefix, p.Slug)
			}
			var existing contentdomain.Post
			lookupErr := tx.Where("slug = ?", p.Slug).First(&existing).Error
			if lookupErr == nil {
				skipped++
				continue
			}
			if !errors.Is(lookupErr, gorm.ErrRecordNotFound) {
				return fmt.Errorf("lookup post %q: %w", p.Slug, lookupErr)
			}
			publishedAt := now.AddDate(0, 0, -p.DaysAgo)
			post := contentdomain.Post{
				Slug:        p.Slug,
				Type:        constants.PostTypeBlog,
				TitleJSON:   demoLocalized(p.Title),
				SummaryJSON: demoLocalized(p.Summary),
				ContentJSON: demoLocalized(p.Content),
				Thumbnail:   p.Cover,
				IsPublished: true,
				PublishedAt: &publishedAt,
				CreatedAt:   publishedAt,
			}
			if createErr := tx.Create(&post).Error; createErr != nil {
				return fmt.Errorf("create post %q: %w", p.Slug, createErr)
			}
			created++
		}
		return nil
	})
	return created, skipped, err
}

// removeDemoPosts 彻底删除 slug 以 demo- 开头的示例文章（真实文章不受影响）。
func removeDemoPosts(db *gorm.DB) (int64, error) {
	result := db.Unscoped().Where("slug LIKE ?", demoPostSlugPrefix+"%").Delete(&contentdomain.Post{})
	return result.RowsAffected, result.Error
}

func runSeedDemoPosts() {
	created, skipped, err := seedDemoPosts(gormdb.DB, demoPosts, time.Now())
	if err != nil {
		fmt.Fprintf(os.Stderr, "seed-demo-posts: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("示例文章：新建 %d 篇，已存在跳过 %d 篇。\n", created, skipped)
	fmt.Println("看完样式后可执行 admin remove-demo-posts 一键删除。")
}

func runRemoveDemoPosts() {
	n, err := removeDemoPosts(gormdb.DB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remove-demo-posts: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("已删除示例文章 %d 篇。\n", n)
}
