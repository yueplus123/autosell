package admincmd

import (
	"errors"
	"fmt"
	"os"

	categorydomain "github.com/dujiao-next/internal/modules/catalog/category/domain"
	"github.com/dujiao-next/internal/platform/database/gormdb"
	"github.com/dujiao-next/internal/shared/jsonmap"

	"gorm.io/gorm"
)

// seedCategory 描述一个待初始化的分类节点。
type seedCategory struct {
	Slug      string
	Name      map[string]interface{}
	SortOrder int
	Children  []seedCategory
}

func seedName(zhCN, zhTW, enUS string) map[string]interface{} {
	return map[string]interface{}{"zh-CN": zhCN, "zh-TW": zhTW, "en-US": enUS}
}

// businessCategorySeed 是首页「双业务入口」与 /zone/* 专区页依赖的分类结构。
// 一级分类 slug 必须与前端 frontend/user/src/utils/businessLines.ts 中的配置一致。
// 前台按 sort_order 倒序展示，数值越大越靠前。
var businessCategorySeed = []seedCategory{
	{
		Slug: "ai", Name: seedName("AI 服务", "AI 服務", "AI Services"), SortOrder: 200,
		Children: []seedCategory{
			{Slug: "chatgpt", Name: seedName("ChatGPT", "ChatGPT", "ChatGPT"), SortOrder: 40},
			{Slug: "claude", Name: seedName("Claude", "Claude", "Claude"), SortOrder: 30},
			{Slug: "midjourney", Name: seedName("Midjourney", "Midjourney", "Midjourney"), SortOrder: 20},
			{Slug: "ai-api", Name: seedName("API 额度", "API 額度", "API Credits"), SortOrder: 10},
			{Slug: "grok", Name: seedName("Grok", "Grok", "Grok"), SortOrder: 9},
			{Slug: "gemini", Name: seedName("Gemini", "Gemini", "Gemini"), SortOrder: 8},
			{Slug: "apple-id", Name: seedName("海外 Apple ID", "海外 Apple ID", "Overseas Apple ID"), SortOrder: 7},
			{Slug: "gift-card", Name: seedName("礼品卡", "禮品卡", "Gift Cards"), SortOrder: 6},
			{Slug: "visa-mastercard", Name: seedName("Visa / 万事达", "Visa / 萬事達", "Visa / Mastercard"), SortOrder: 5},
			{Slug: "gmail", Name: seedName("Gmail 邮箱", "Gmail 信箱", "Gmail"), SortOrder: 4},
		},
	},
	{
		Slug: "game-topup", Name: seedName("二次元赞助", "二次元贊助", "Anime Support"), SortOrder: 190,
		Children: []seedCategory{
			{Slug: "genshin-impact", Name: seedName("原神", "原神", "Genshin Impact"), SortOrder: 40},
			{Slug: "honkai-star-rail", Name: seedName("崩坏：星穹铁道", "崩壞：星穹鐵道", "Honkai: Star Rail"), SortOrder: 30},
			{Slug: "zenless-zone-zero", Name: seedName("绝区零", "絕區零", "Zenless Zone Zero"), SortOrder: 20},
			{Slug: "fgo", Name: seedName("Fate/Grand Order", "Fate/Grand Order", "Fate/Grand Order"), SortOrder: 10},
		},
	},
}

// seedResult 汇总一次初始化的结果，便于打印与测试。
type seedResult struct {
	Created  []string
	Restored []string
	Skipped  []string
	Warnings []string
}

// seedBusinessCategories 幂等地创建业务分类：
//   - slug 不存在 → 新建；
//   - slug 存在且未删除 → 原样保留（不覆盖站长在后台改过的名称、图标、排序）；
//   - slug 存在但已软删除 → 恢复并挂到正确的父分类下（slug 有唯一索引，无法重建）。
//
// 整个过程在一个事务内完成，任一步失败全部回滚。
func seedBusinessCategories(db *gorm.DB, seeds []seedCategory) (*seedResult, error) {
	result := &seedResult{}
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, root := range seeds {
			rootID, err := ensureCategory(tx, root, 0, result)
			if err != nil {
				return err
			}
			for _, child := range root.Children {
				if _, err := ensureCategory(tx, child, rootID, result); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ensureCategory(tx *gorm.DB, seed seedCategory, parentID uint, result *seedResult) (uint, error) {
	var existing categorydomain.Category
	err := tx.Where("slug = ?", seed.Slug).First(&existing).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		category := categorydomain.Category{
			ParentID:  parentID,
			Slug:      seed.Slug,
			NameJSON:  jsonmap.JSON(seed.Name),
			SortOrder: seed.SortOrder,
			IsActive:  true,
		}
		if err := tx.Create(&category).Error; err != nil {
			return 0, fmt.Errorf("create category %q: %w", seed.Slug, err)
		}
		result.Created = append(result.Created, seed.Slug)
		return category.ID, nil
	case err != nil:
		return 0, fmt.Errorf("lookup category %q: %w", seed.Slug, err)
	}

	if existing.DeletedAt != nil {
		updates := map[string]interface{}{
			"deleted_at": nil,
			"parent_id":  parentID,
			"is_active":  true,
		}
		if err := tx.Model(&categorydomain.Category{}).Where("id = ?", existing.ID).Updates(updates).Error; err != nil {
			return 0, fmt.Errorf("restore category %q: %w", seed.Slug, err)
		}
		result.Restored = append(result.Restored, seed.Slug)
		return existing.ID, nil
	}

	result.Skipped = append(result.Skipped, seed.Slug)
	if existing.ParentID != parentID {
		if parentID == 0 {
			result.Warnings = append(result.Warnings, fmt.Sprintf("分类 %q 已存在但不是一级分类，首页业务入口无法关联到它；请在后台把它的上级分类改为「无」", seed.Slug))
		} else {
			result.Warnings = append(result.Warnings, fmt.Sprintf("分类 %q 已存在但挂在其他上级分类下，未做改动", seed.Slug))
		}
	}
	return existing.ID, nil
}

func runSeedBusinessCategories() {
	result, err := seedBusinessCategories(gormdb.DB, businessCategorySeed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "seed-categories: %v\n", err)
		os.Exit(1)
	}
	printSlugs := func(label string, slugs []string) {
		if len(slugs) == 0 {
			return
		}
		fmt.Printf("%s (%d):\n", label, len(slugs))
		for _, slug := range slugs {
			fmt.Printf("  - %s\n", slug)
		}
	}
	printSlugs("新建", result.Created)
	printSlugs("已恢复（之前被删除）", result.Restored)
	printSlugs("已存在，保持不变", result.Skipped)
	for _, warning := range result.Warnings {
		fmt.Fprintf(os.Stderr, "注意: %s\n", warning)
	}
	fmt.Println("完成。可在后台「分类管理」中修改名称、上传图标、调整排序。")
}
