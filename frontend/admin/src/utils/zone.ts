import type { AdminCategory } from '@/api/types'

/**
 * 前台专区（AI / 二次元）由一级分类的 slug 界定。
 * 必须与后端 internal/modules/reseller/domain/zone.go 和前台 utils/businessLines.ts 保持一致。
 */
export type StoreZone = 'ai' | 'game'

export const STORE_ZONES: StoreZone[] = ['ai', 'game']

export const ZONE_ROOT_SLUGS: Record<StoreZone, string[]> = {
  ai: ['ai', 'ai-services', 'ai-tools'],
  game: ['game-topup', 'topup', 'acg', 'game'],
}

const normalize = (value: unknown) => String(value ?? '').trim().toLowerCase()

/** 专区对应的一级分类（找不到时返回 null，说明还没建分类） */
export const findZoneRoot = (zone: StoreZone, categories: AdminCategory[]): AdminCategory | null => {
  for (const slug of ZONE_ROOT_SLUGS[zone]) {
    const hit = categories.find((c) => !(Number(c.parent_id) > 0) && normalize(c.slug) === slug)
    if (hit) return hit
  }
  return null
}

/** 某个分类属于哪个专区（一级分类本身或其子分类），都不属于时返回 null */
export const zoneOfCategory = (categoryId: number | null | undefined, categories: AdminCategory[]): StoreZone | null => {
  const id = Number(categoryId || 0)
  if (!id) return null
  const category = categories.find((c) => c.id === id)
  if (!category) return null
  const rootId = Number(category.parent_id) > 0 ? Number(category.parent_id) : category.id
  for (const zone of STORE_ZONES) {
    if (findZoneRoot(zone, categories)?.id === rootId) return zone
  }
  return null
}
