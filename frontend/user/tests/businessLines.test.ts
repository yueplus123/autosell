import test from 'node:test'
import assert from 'node:assert/strict'
import { resolveBusinessLines } from '../src/utils/businessLines.ts'

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
  assert.equal(ai.link, '/categories/AI')
  assert.equal(game.category?.id, 4)
  assert.deepEqual(game.children.map((c) => c.id), [5, 6])
})

test('business lines ignore child categories with a matching slug', () => {
  const [ai] = resolveBusinessLines([{ id: 6, slug: 'ai', parent_id: 4 }])
  assert.equal(ai.category, null)
})

test('unmatched business line falls back to all products', () => {
  const lines = resolveBusinessLines([])
  assert.equal(lines.length, 2)
  for (const line of lines) {
    assert.equal(line.category, null)
    assert.deepEqual(line.children, [])
    assert.equal(line.link, '/products')
  }
})

test('children are capped', () => {
  const many = [{ id: 1, slug: 'ai', parent_id: 0 }]
  for (let i = 0; i < 10; i++) many.push({ id: 10 + i, slug: `c${i}`, parent_id: 1 })
  const [ai] = resolveBusinessLines(many, undefined, 4)
  assert.equal(ai.children.length, 4)
})
