/**
 * 薪火模板首屏图标墙配置。
 *
 * 图标文件放在 frontend/user/public/xinhuo/icons/ 下，这里按文件名引用。
 * 换图标：替换同名文件即可；增删砖块：增删下面的条目。
 *
 * 坐标单位都是百分比：
 *   x  砖块中心点距左边的位置（占图标墙宽度）
 *   y  砖块上边缘距顶部的位置（占图标墙高度）
 *   s  砖块边长（占图标墙宽度）
 * 桌面版图标墙宽高比 1200:430，手机版（≤640px）390:260，所以两套坐标分开写（mx / my / ms）。
 */

export type HeroTileTone = 'light' | 'dark'

export interface HeroTile {
  /** 图标文件名（public/xinhuo/icons/ 下） */
  icon: string
  /** 无障碍名称，读屏时朗读 */
  label: string
  /** light：浅色砖块；dark：黑色砖块（黑色线稿 logo 会自动反白） */
  tone: HeroTileTone
  /** back：后排小砖块，带轻微模糊；front：前排，带地面倒影 */
  row: 'back' | 'front'
  /** 横向 logo（如 VISA）用更宽的图片区域 */
  wide?: boolean
  /** 砖块外发光颜色 */
  glow?: string
  /** 叠放层级，越大越靠前 */
  z?: number
  x: number
  y: number
  s: number
  mx: number
  my: number
  ms: number
}

export const HERO_ICON_BASE = '/xinhuo/icons/'

export const heroTiles: HeroTile[] = [
  // 后排
  { icon: 'pixiv.png', label: 'pixiv', tone: 'light', row: 'back', glow: 'rgba(14,165,233,0.45)', x: 27, y: 9, s: 8, mx: 22, my: 8, ms: 13 },
  { icon: 'visa.png', label: 'Visa', tone: 'light', row: 'back', wide: true, glow: 'rgba(37,99,235,0.45)', x: 73, y: 9, s: 8, mx: 78, my: 8, ms: 13 },
  // 前排（中间最大的是主推）
  { icon: 'gmail.png', label: 'Gmail', tone: 'light', row: 'front', glow: 'rgba(239,68,68,0.25)', x: 19, y: 45, s: 11, mx: 11, my: 48, ms: 16 },
  { icon: 'gemini.png', label: 'Gemini', tone: 'light', row: 'front', glow: 'rgba(59,130,246,0.3)', z: 2, x: 33.5, y: 40, s: 13, mx: 29, my: 43.5, ms: 19 },
  { icon: 'chatgpt.png', label: 'ChatGPT', tone: 'dark', row: 'front', z: 3, x: 50, y: 28, s: 17, mx: 50, my: 34.5, ms: 25 },
  { icon: 'grok.png', label: 'Grok', tone: 'dark', row: 'front', z: 2, x: 66.5, y: 40, s: 13, mx: 71, my: 43.5, ms: 19 },
  { icon: 'apple.png', label: 'Apple', tone: 'light', row: 'front', glow: 'rgba(255,255,255,0.12)', x: 81, y: 45, s: 11, mx: 89, my: 48, ms: 16 },
]
