import type { PublicCategory } from './category'

/**
 * 首页「双业务入口」配置。
 *
 * 每条业务线通过一级分类的 slug 关联到后台分类：在后台把对应一级分类的 slug
 * 设置为 slugs 中的任意一个即可自动关联（推荐 `ai` 与 `game-topup`）。
 * 每条业务线都有独立的专区页 /zone/<key>，首页卡片点击进入专区；
 * 专区页按关联分类（含子分类）展示商品，未关联分类时显示"即将上线"。
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

export const isBusinessLineKey = (value: unknown): value is BusinessLineKey =>
  BUSINESS_LINES.some((def) => def.key === value)

export const businessZonePath = (key: BusinessLineKey, childSlug?: string) =>
  childSlug ? `/zone/${key}?c=${encodeURIComponent(childSlug)}` : `/zone/${key}`

const normalizeSlug = (value: unknown) => String(value ?? '').trim().toLowerCase()

const toPositiveInt = (value: unknown) => {
  const n = Number(value)
  return Number.isFinite(n) && n > 0 ? Math.floor(n) : 0
}

export const resolveBusinessLines = (
  categories: PublicCategory[],
  definitions: BusinessLineDefinition[] = BUSINESS_LINES,
  maxChildren = Number.POSITIVE_INFINITY,
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
    return { key: def.key, category, children, link: businessZonePath(def.key) }
  })
}

export interface HeroTile {
  /** 关联的子分类；为 null 时用占位图标 */
  category: PublicCategory | null
  /** 占位图标序号（仅 category 为空或没有上传图标时使用） */
  fallbackIndex: number
}

/**
 * 生成首页 AI 主视觉的图标方块：优先使用上传了图标的子分类，其次是没有图标的子分类，
 * 不足 count 个时用占位图标补齐。第 0 个放在正中间（最大的方块）。
 */
export const buildHeroTiles = (children: PublicCategory[], count = 7): HeroTile[] => {
  const list = Array.isArray(children) ? children : []
  const withIcon = list.filter((c) => typeof c.icon === 'string' && c.icon.trim() !== '')
  const withoutIcon = list.filter((c) => !(typeof c.icon === 'string' && c.icon.trim() !== ''))
  const ordered = [...withIcon, ...withoutIcon].slice(0, count)
  const tiles: HeroTile[] = ordered.map((category, i) => ({ category, fallbackIndex: i }))
  for (let i = tiles.length; i < count; i++) tiles.push({ category: null, fallbackIndex: i })
  return tiles
}
