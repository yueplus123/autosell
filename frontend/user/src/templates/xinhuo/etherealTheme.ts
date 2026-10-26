import { ref, watch } from 'vue'

/**
 * 官网首页（Ethereal 风格）的访客显示设置：主题色相、深浅色、文章列表/网格。
 * 只影响访客自己的浏览器（存在 localStorage），存储不可用时用默认值。
 */
export type EtherealLayout = 'list' | 'grid'

const STORAGE_KEY = 'xh-ethereal-theme'
export const DEFAULT_HUE = 225

interface Saved { hue?: number; dark?: boolean; layout?: EtherealLayout }

const read = (): Saved => {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    return raw ? (JSON.parse(raw) as Saved) : {}
  } catch {
    return {}
  }
}

const saved = read()
const systemDark = () => {
  try {
    return typeof window !== 'undefined' && window.matchMedia('(prefers-color-scheme: dark)').matches
  } catch {
    return false
  }
}

export const ethHue = ref<number>(typeof saved.hue === 'number' ? saved.hue : DEFAULT_HUE)
export const ethDark = ref<boolean>(typeof saved.dark === 'boolean' ? saved.dark : systemDark())
export const ethLayout = ref<EtherealLayout>(saved.layout === 'grid' ? 'grid' : 'list')

watch([ethHue, ethDark, ethLayout], () => {
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify({ hue: ethHue.value, dark: ethDark.value, layout: ethLayout.value }))
  } catch {
    // 存储不可用时只在本次浏览生效
  }
})

export const resetEtherealTheme = () => {
  ethHue.value = DEFAULT_HUE
  ethLayout.value = 'list'
}

/** 根据色相与深浅色生成页面配色变量 */
export const etherealVars = (hue: number, dark: boolean): Record<string, string> =>
  dark
    ? {
        '--eth-bg': `hsl(${hue} 18% 9%)`,
        '--eth-card': `hsl(${hue} 16% 13%)`,
        '--eth-ink': `hsl(${hue} 20% 92%)`,
        '--eth-muted': `hsl(${hue} 10% 64%)`,
        '--eth-line': `hsl(${hue} 14% 20%)`,
        '--eth-accent': `hsl(${hue} 85% 68%)`,
        '--eth-accent-soft': `hsl(${hue} 45% 22%)`,
        '--eth-hover-line': `hsl(${hue} 40% 34%)`,
        '--eth-shadow': 'rgba(0, 0, 0, 0.35)',
      }
    : {
        '--eth-bg': `hsl(${hue} 40% 95%)`,
        '--eth-card': '#ffffff',
        '--eth-ink': `hsl(${hue} 22% 16%)`,
        '--eth-muted': `hsl(${hue} 10% 47%)`,
        '--eth-line': `hsl(${hue} 30% 91%)`,
        '--eth-accent': `hsl(${hue} 82% 62%)`,
        '--eth-accent-soft': `hsl(${hue} 90% 95%)`,
        '--eth-hover-line': `hsl(${hue} 70% 85%)`,
        '--eth-shadow': `hsla(${hue} 60% 40% / 0.12)`,
      }
