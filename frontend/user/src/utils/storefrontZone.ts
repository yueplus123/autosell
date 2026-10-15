import { businessZonePath, isBusinessLineKey, type BusinessLineKey } from './businessLines.ts'

/**
 * 专区「锁定」逻辑：
 * - 分销商只代理一个专区（tenant.zone_scope = ai / game）时，整个分站就是这个专区：
 *   首页、全部商品、分类页都重定向到该专区，其他专区不可访问；
 * - 主站上，顾客一旦进入某个专区，站内所有「返回首页 / 全部商品」的入口都改为回到这个专区，
 *   不再出现能回到总首页的链接（在浏览器里手动输入网址不受影响）。
 */

const ACTIVE_ZONE_KEY = 'xh-active-zone'

interface ConfigLike {
  tenant?: { mode?: unknown; zone_scope?: unknown } | null
}

/** 分站只代理一个专区时返回该专区，否则返回 null（主站或代理全部专区） */
export const tenantLockedZone = (config: ConfigLike | null | undefined): BusinessLineKey | null => {
  const tenant = config?.tenant
  if (!tenant || String(tenant.mode || '') !== 'reseller') return null
  const scope = String(tenant.zone_scope || '')
  return isBusinessLineKey(scope) ? scope : null
}

export const readActiveZone = (): BusinessLineKey | null => {
  try {
    const value = sessionStorage.getItem(ACTIVE_ZONE_KEY)
    return isBusinessLineKey(value) ? value : null
  } catch {
    return null
  }
}

export const rememberActiveZone = (zone: BusinessLineKey) => {
  try {
    sessionStorage.setItem(ACTIVE_ZONE_KEY, zone)
  } catch {
    // 存储不可用时只影响「回到专区」的落点，不影响浏览
  }
}

/** 当前顾客所在的专区：分站锁定的专区优先，其次是本次会话进入过的专区 */
export const currentZone = (config: ConfigLike | null | undefined): BusinessLineKey | null =>
  tenantLockedZone(config) ?? readActiveZone()

/** 站内「首页 / 继续购物」应指向的地址 */
export const storefrontHomePath = (config: ConfigLike | null | undefined): string => {
  const zone = currentZone(config)
  return zone ? businessZonePath(zone) : '/'
}

/** 「全部商品」类入口应指向的地址：在专区内时回到专区，否则为全部商品页 */
export const storefrontShopPath = (config: ConfigLike | null | undefined): string => {
  const zone = currentZone(config)
  return zone ? businessZonePath(zone) : '/products'
}

/** 专区内的搜索地址 */
export const zoneSearchPath = (zone: BusinessLineKey, keyword: string) => {
  const q = keyword.trim()
  return q ? `${businessZonePath(zone)}?q=${encodeURIComponent(q)}` : businessZonePath(zone)
}
