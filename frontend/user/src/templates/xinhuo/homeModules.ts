import type { BusinessLineKey } from '../../utils/businessLines'
import { zoneSearchPath } from '../../utils/storefrontZone'

/**
 * 薪火首页商品展示模块配置（「热门商品」的替代）。
 *
 * 每个模块是一组方块（square）或一组横向长条（banner）。
 * 每个方块的跳转：
 *   - product：填商品的 slug（商品详情页地址 /products/<slug> 里的那段），点击直接进入该商品；
 *   - 否则用 search：点击进入所属专区（zone，默认 AI 专区）并按这个关键词搜索。
 * 图标文件放在 frontend/user/public/xinhuo/icons/ 下，按文件名引用。
 * 文案在语言包 xinhuo.modules.* 下（titleKey / subKey 指向的键）。
 */

export type ModuleTileTone = 'light' | 'dark'

export interface ModuleTile {
  key: string
  /** 方块上的大字（如 Plus / 5x），不翻译 */
  label?: string
  /** 方块上的大字，取自语言包（与 label 二选一） */
  labelKey?: string
  /** 小字说明，取自语言包 */
  subKey?: string
  icon: string
  tone: ModuleTileTone
  /** 横向 logo（如 VISA）用更宽的图片区域 */
  wide?: boolean
  /** 图标是否反色。默认黑色方块反色（适合 ChatGPT、Grok 这类黑色线稿 logo），彩色 logo 设为 false */
  invertIcon?: boolean
  /** 背景光晕颜色 */
  glow: string
  product?: string
  search?: string
  /** 所属专区，默认 ai */
  zone?: BusinessLineKey
}

export interface HomeModule {
  key: string
  layout: 'square' | 'banner'
  /** 方块布局时电脑端每行几个（默认 4；2 时每个占一半，方块改为横向比例） */
  desktopColumns?: 2 | 4
  titleKey: string
  descKey?: string
  tiles: ModuleTile[]
}

export const homeModules: HomeModule[] = [
  {
    key: 'chatgpt',
    layout: 'square',
    titleKey: 'xinhuo.modules.chatgpt.title',
    descKey: 'xinhuo.modules.chatgpt.desc',
    tiles: [
      { key: 'plus', label: 'Plus', subKey: 'xinhuo.modules.chatgpt.sub', icon: 'chatgpt.png', tone: 'dark', glow: 'rgba(16,163,127,0.45)', search: 'ChatGPT' },
      { key: '5x', label: '5x', subKey: 'xinhuo.modules.chatgpt.sub', icon: 'chatgpt.png', tone: 'dark', glow: 'rgba(34,211,238,0.4)', search: 'ChatGPT' },
      { key: '20x', label: '20x', subKey: 'xinhuo.modules.chatgpt.sub', icon: 'chatgpt.png', tone: 'dark', glow: 'rgba(99,102,241,0.45)', search: 'ChatGPT' },
      { key: '50x', label: '50x', subKey: 'xinhuo.modules.chatgpt.sub', icon: 'chatgpt.png', tone: 'dark', glow: 'rgba(217,70,239,0.45)', search: 'ChatGPT' },
    ],
  },
  {
    key: 'global',
    layout: 'banner',
    titleKey: 'xinhuo.modules.global.title',
    tiles: [
      { key: 'apple', labelKey: 'xinhuo.modules.global.apple', subKey: 'xinhuo.modules.global.appleSub', icon: 'apple.png', tone: 'light', glow: 'rgba(148,163,184,0.45)', search: 'Apple' },
      { key: 'card', labelKey: 'xinhuo.modules.global.card', subKey: 'xinhuo.modules.global.cardSub', icon: 'visa.png', tone: 'light', wide: true, glow: 'rgba(37,99,235,0.45)', search: 'Visa' },
    ],
  },
  {
    key: 'more',
    layout: 'square',
    desktopColumns: 2,
    titleKey: 'xinhuo.modules.more.title',
    tiles: [
      { key: 'grok', labelKey: 'xinhuo.modules.more.grok', subKey: 'xinhuo.modules.more.member', icon: 'grok.png', tone: 'dark', glow: 'rgba(168,85,247,0.4)', search: 'Grok' },
      { key: 'gemini', labelKey: 'xinhuo.modules.more.gemini', subKey: 'xinhuo.modules.more.member', icon: 'gemini.png', tone: 'light', glow: 'rgba(59,130,246,0.45)', search: 'Gemini' },
      { key: 'card', labelKey: 'xinhuo.modules.more.card', subKey: 'xinhuo.modules.more.cardSub', icon: 'visa.png', tone: 'light', wide: true, glow: 'rgba(37,99,235,0.45)', search: 'Visa' },
      { key: 'gmail', labelKey: 'xinhuo.modules.more.gmail', subKey: 'xinhuo.modules.more.gmailSub', icon: 'gmail.png', tone: 'dark', invertIcon: false, glow: 'rgba(239,68,68,0.4)', search: 'Gmail' },
    ],
  },
]

export const moduleTileLink = (tile: ModuleTile) => {
  if (tile.product) return `/products/${encodeURIComponent(tile.product)}`
  return zoneSearchPath(tile.zone ?? 'ai', tile.search ?? '')
}
