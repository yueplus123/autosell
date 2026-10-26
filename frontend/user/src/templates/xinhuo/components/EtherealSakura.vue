<template>
  <canvas ref="canvas" class="pointer-events-none fixed inset-0 z-[40]" aria-hidden="true"></canvas>
</template>

<script setup lang="ts">
// 樱花飘落特效（参考 Ethereal 主题）。系统开启「减少动态效果」时不显示。
import { onMounted, onUnmounted, ref } from 'vue'

const canvas = ref<HTMLCanvasElement | null>(null)
let frame = 0
let resizeHandler: (() => void) | null = null

interface Petal { x: number; y: number; r: number; speed: number; drift: number; angle: number; spin: number; alpha: number }

onMounted(() => {
  const el = canvas.value
  if (!el) return
  if (typeof window.matchMedia === 'function' && window.matchMedia('(prefers-reduced-motion: reduce)').matches) return
  const ctx = el.getContext('2d')
  if (!ctx) return

  let width = 0
  let height = 0
  const dpr = Math.min(window.devicePixelRatio || 1, 2)
  const resize = () => {
    width = window.innerWidth
    height = window.innerHeight
    el.width = width * dpr
    el.height = height * dpr
    el.style.width = `${width}px`
    el.style.height = `${height}px`
    ctx.setTransform(dpr, 0, 0, dpr, 0, 0)
  }
  resize()
  resizeHandler = resize
  window.addEventListener('resize', resize)

  const count = width < 640 ? 14 : 26
  const spawn = (initial: boolean): Petal => ({
    x: Math.random() * width,
    y: initial ? Math.random() * height : -20,
    r: 5 + Math.random() * 6,
    speed: 0.4 + Math.random() * 0.8,
    drift: -0.3 + Math.random() * 0.9,
    angle: Math.random() * Math.PI * 2,
    spin: -0.02 + Math.random() * 0.04,
    alpha: 0.55 + Math.random() * 0.35,
  })
  const petals: Petal[] = Array.from({ length: count }, () => spawn(true))

  const draw = (p: Petal) => {
    ctx.save()
    ctx.translate(p.x, p.y)
    ctx.rotate(p.angle)
    ctx.globalAlpha = p.alpha
    ctx.fillStyle = '#ffc0d3'
    ctx.beginPath()
    ctx.moveTo(0, -p.r)
    ctx.bezierCurveTo(p.r, -p.r, p.r, p.r * 0.6, 0, p.r)
    ctx.bezierCurveTo(-p.r, p.r * 0.6, -p.r, -p.r, 0, -p.r)
    ctx.fill()
    ctx.restore()
  }

  const tick = () => {
    ctx.clearRect(0, 0, width, height)
    for (let i = 0; i < petals.length; i++) {
      const p = petals[i]!
      p.y += p.speed
      p.x += p.drift + Math.sin(p.y / 60) * 0.4
      p.angle += p.spin
      if (p.y > height + 20 || p.x < -30 || p.x > width + 30) petals[i] = spawn(false)
      draw(petals[i]!)
    }
    frame = requestAnimationFrame(tick)
  }
  frame = requestAnimationFrame(tick)
})

onUnmounted(() => {
  cancelAnimationFrame(frame)
  if (resizeHandler) window.removeEventListener('resize', resizeHandler)
})
</script>
