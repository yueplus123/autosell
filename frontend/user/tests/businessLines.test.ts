import test from 'node:test'
import assert from 'node:assert/strict'
import { buildHeroTiles, businessZonePath, isBusinessLineKey, resolveBusinessLines } from '../src/utils/businessLines.ts'

const categories = [
  { id: 1, slug: 'AI', parent_id: 0 },
  { id: 2, slug: 'chatgpt', parent_id: 1 },
  { id: 3, slug: 'claude', parent_id: 1 },
  { id: 4, slug: 'game-topup', parent_id: 0 },
  { id: 5, slug: 'genshin', parent_id: 4 },
  { id: 6, slug: 'ai', parent_id: 4 },
]

test('business lines match root categories by slug, case-insensitive', () => {
  const [ai, game] = resolveBusinessLines(categories)
  assert.equal(ai.category?.id, 1)
  assert.deepEqual(ai.children.map((c) => c.id), [2, 3])
  assert.equal(ai.link, '/zone/ai')
  assert.equal(game.link, '/zone/game')
  assert.equal(game.category?.id, 4)
  assert.deepEqual(game.children.map((c) => c.id), [5, 6])
})

test('business lines ignore child categories with a matching slug', () => {
  const [ai] = resolveBusinessLines([{ id: 6, slug: 'ai', parent_id: 4 }])
  assert.equal(ai.category, null)
})

test('unmatched business line still links to its own zone', () => {
  const lines = resolveBusinessLines([])
  assert.equal(lines.length, 2)
  for (const line of lines) {
    assert.equal(line.category, null)
    assert.deepEqual(line.children, [])
    assert.equal(line.link, `/zone/${line.key}`)
  }
})

test('children are capped', () => {
  const many = [{ id: 1, slug: 'ai', parent_id: 0 }]
  for (let i = 0; i < 10; i++) many.push({ id: 10 + i, slug: `c${i}`, parent_id: 1 })
  const [ai] = resolveBusinessLines(many, undefined, 4)
  assert.equal(ai.children.length, 4)
})

test('zone helpers', () => {
  assert.equal(businessZonePath('game', '原神'), '/zone/game?c=%E5%8E%9F%E7%A5%9E')
  assert.equal(isBusinessLineKey('ai'), true)
  assert.equal(isBusinessLineKey('other'), false)
})

test('hero tiles prefer categories with icons and pad with placeholders', () => {
  const tiles = buildHeroTiles([
    { id: 1, slug: 'a' },
    { id: 2, slug: 'b', icon: '/uploads/b.png' },
    { id: 3, slug: 'c', icon: '  ' },
  ], 5)
  assert.deepEqual(tiles.map((t) => t.category?.id ?? null), [2, 1, 3, null, null])
  assert.deepEqual(tiles.map((t) => t.fallbackIndex), [0, 1, 2, 3, 4])
})

test('hero tiles cap at count', () => {
  const many = Array.from({ length: 10 }, (_, i) => ({ id: i + 1, slug: `s${i}` }))
  assert.equal(buildHeroTiles(many, 7).length, 7)
  assert.equal(buildHeroTiles([], 7).every((t) => t.category === null), true)
})
