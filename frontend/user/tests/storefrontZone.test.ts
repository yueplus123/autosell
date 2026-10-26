import test from 'node:test'
import assert from 'node:assert/strict'

// sessionStorage 在 node 下不存在，按浏览器行为模拟
const store = new Map<string, string>()
;(globalThis as any).sessionStorage = {
  getItem: (k: string) => (store.has(k) ? store.get(k)! : null),
  setItem: (k: string, v: string) => void store.set(k, v),
  removeItem: (k: string) => void store.delete(k),
}

const { tenantLockedZone, storefrontHomePath, storefrontShopPath, rememberActiveZone, zoneSearchPath, isMainShowcase } = await import('../src/utils/storefrontZone.ts')

test('main site without an active zone keeps normal home and shop paths', () => {
  store.clear()
  const config = { tenant: { mode: 'main' }, showcase_mode: false }
  assert.equal(tenantLockedZone(config), null)
  assert.equal(storefrontHomePath(config), '/')
  assert.equal(storefrontShopPath(config), '/products')
  // 展示模式下没有全部商品页，回到首页
  assert.equal(storefrontShopPath({ tenant: { mode: 'main' } }), '/')
})

test('after entering a zone, home and shop links point back into it', () => {
  store.clear()
  rememberActiveZone('game')
  const config = { tenant: { mode: 'main' }, showcase_mode: false }
  assert.equal(storefrontHomePath(config), '/zone/game')
  assert.equal(storefrontShopPath(config), '/zone/game')
  // 官网展示模式：首页链接始终回到文章首页
  assert.equal(storefrontHomePath({ tenant: { mode: 'main' } }), '/')
})

test('single-zone reseller is locked to its zone regardless of session', () => {
  store.clear()
  rememberActiveZone('game')
  const config = { tenant: { mode: 'reseller', zone_scope: 'ai' } }
  assert.equal(tenantLockedZone(config), 'ai')
  assert.equal(storefrontHomePath(config), '/zone/ai')
})

test('all-zone reseller is not locked', () => {
  store.clear()
  assert.equal(tenantLockedZone({ tenant: { mode: 'reseller', zone_scope: 'all' } }), null)
  assert.equal(tenantLockedZone({ tenant: { mode: 'reseller' } }), null)
})

test('zone search path encodes keyword', () => {
  assert.equal(zoneSearchPath('ai', ' ChatGPT Plus '), '/zone/ai?q=ChatGPT%20Plus')
  assert.equal(zoneSearchPath('ai', ''), '/zone/ai')
})

test('主站展示模式：主站默认开启，分销站永远关闭', () => {
  assert.equal(isMainShowcase(null), false)
  assert.equal(isMainShowcase({ tenant: { mode: 'main' } }), true)
  assert.equal(isMainShowcase({ tenant: { mode: 'main' }, showcase_mode: true }), true)
  assert.equal(isMainShowcase({ tenant: { mode: 'main' }, showcase_mode: false }), false)
  assert.equal(isMainShowcase({}), true)
  assert.equal(isMainShowcase({ tenant: { mode: 'reseller' }, showcase_mode: true }), false)
})
