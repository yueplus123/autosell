import type { PublicCategory } from './category'

/**
 * 首页「双业务入口」配置。
 *
 * 每条业务线通过一级分类的 slug 关联到后台分类：在后台把对应一级分类的 slug
 * 设置为 slugs 中的任意一个即可自动关联（推荐 `ai` 与 `game-topup`）。
 * 关联后卡片跳转到该分类页，并展示其子分类标签与热门商品；
 * 未关联时卡片仍然显示，点击跳转到全部商品。
 */
export type BusinessLineKey = 'ai' | 'game'

export interface BusinessLineDefinition {
  key: BusinessLineKey
  slugs: string[]
}

export const BUSINESS_LINES: BusinessLineDefinition[] = [
  { key: 'ai', slugs: ['ai', 'ai-services', 'ai-tools'] },
  { key: 'game', slugs: ['game-topup', 'topup', 'acg', 'game'] },
]

export interface ResolvedBusinessLine {
  key: BusinessLineKey
  category: PublicCategory | null
  children: PublicCategory[]
  link: string
}

const normalizeSlug = (value: unknown) => String(value ?? '').trim().toLowerCase()

const toPositiveInt = (value: unknown) => {
  const n = Number(value)
  return Number.isFinite(n) && n > 0 ? Math.floor(n) : 0
}

export const resolveBusinessLines = (
  categories: PublicCategory[],
  definitions: BusinessLineDefinition[] = BUSINESS_LINES,
  maxChildren = 6,
): ResolvedBusinessLine[] => {
  const list = Array.isArray(categories) ? categories : []
  const roots = list.filter((c) => toPositiveInt(c.parent_id) === 0)

  return definitions.map((def) => {
    const wanted = def.slugs.map(normalizeSlug)
    let category: PublicCategory | null = null
    for (const slug of wanted) {
      category = roots.find((c) => normalizeSlug(c.slug) === slug) ?? null
      if (category) break
    }
    const children = category
      ? list
          .filter((c) => toPositiveInt(c.parent_id) === toPositiveInt(category!.id) && c.slug)
          .slice(0, maxChildren)
      : []
    const link = category?.slug ? `/categories/${encodeURIComponent(String(category.slug))}` : '/products'
    return { key: def.key, category, children, link }
  })
}
