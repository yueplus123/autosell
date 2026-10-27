/**
 * 官网首屏和两个专区首屏的业务图标墙（错落排列，和之前的首屏图标墙同一种风格）。
 * 图标文件在 public/xinhuo/icons/ai、acg 下。
 *
 * 坐标单位都是百分比：
 *   x  砖块中心点距左边的位置（占图标墙宽度）
 *   y  砖块上边缘距顶部的位置（占图标墙高度）
 *   s  砖块边长（占图标墙宽度）
 * 首页图标墙宽高比：电脑 1200:430、手机 390:260，两套坐标（x/y/s 与 mx/my/ms）。
 * 专区图标墙是正方形，电脑和手机共用一套坐标。
 * row：back 后排小砖块（轻微模糊）；front 前排（带倒影）。z 越大越靠前。
 */
export interface BrandTile {
  key: string
  label: string
  src: string
  tone: 'light' | 'dark'
  row: 'back' | 'front'
  wide?: boolean
  glow?: string
  z?: number
  x: number
  y: number
  s: number
  mx: number
  my: number
  ms: number
}

const AI = (f: string) => `/xinhuo/icons/ai/${f}`
const ACG = (f: string) => `/xinhuo/icons/acg/${f}`

// 首页首屏：AI 图标，ChatGPT 在最前面正中间
export const HOME_AI_TILES: BrandTile[] = [
  // 后排
  { key: 'google-play', label: 'Google Play', src: AI('google-play.png'), tone: 'light', row: 'back', glow: 'rgba(52,168,83,0.4)', x: 50, y: 3, s: 6, mx: 50, my: 5, ms: 11 },
  { key: 'apple', label: 'Apple', src: AI('apple.png'), tone: 'light', row: 'back', glow: 'rgba(148,163,184,0.45)', x: 29, y: 8, s: 7.5, mx: 22, my: 7, ms: 12 },
  { key: 'visa', label: 'Visa', src: AI('visa.png'), tone: 'light', row: 'back', wide: true, glow: 'rgba(37,99,235,0.5)', x: 71, y: 8, s: 7.5, mx: 78, my: 7, ms: 12 },
  { key: 'openrouter', label: 'OpenRouter', src: AI('openrouter.png'), tone: 'light', row: 'back', glow: 'rgba(99,102,241,0.45)', x: 13, y: 21, s: 6.5, mx: 8, my: 25, ms: 10 },
  { key: 'code', label: 'Code', src: AI('code.png'), tone: 'light', row: 'back', glow: 'rgba(99,102,241,0.4)', x: 87, y: 21, s: 6.5, mx: 92, my: 25, ms: 10 },
  // 前排
  { key: 'gmail', label: 'Gmail', src: AI('gmail.png'), tone: 'light', row: 'front', glow: 'rgba(239,68,68,0.35)', z: 2, x: 7.5, y: 53, s: 8, mx: 36, my: 13, ms: 10 },
  { key: 'mastercard', label: 'Mastercard', src: AI('mastercard.png'), tone: 'light', row: 'front', wide: true, glow: 'rgba(249,115,22,0.4)', z: 2, x: 92.5, y: 53, s: 8, mx: 64, my: 13, ms: 10 },
  { key: 'google', label: 'Google', src: AI('google.png'), tone: 'light', row: 'front', glow: 'rgba(66,133,244,0.4)', z: 3, x: 20, y: 46, s: 10.5, mx: 11, my: 47, ms: 14 },
  { key: 'youtube', label: 'YouTube', src: AI('youtube.png'), tone: 'light', row: 'front', glow: 'rgba(239,68,68,0.4)', z: 3, x: 80, y: 46, s: 10.5, mx: 89, my: 47, ms: 14 },
  { key: 'claude', label: 'Claude', src: AI('claude.png'), tone: 'light', row: 'front', glow: 'rgba(217,119,87,0.45)', z: 4, x: 34, y: 38, s: 13, mx: 29, my: 40, ms: 18 },
  { key: 'grok', label: 'Grok', src: AI('grok.png'), tone: 'dark', row: 'front', glow: 'rgba(168,85,247,0.3)', z: 4, x: 66, y: 38, s: 13, mx: 71, my: 40, ms: 18 },
  { key: 'chatgpt', label: 'ChatGPT', src: AI('chatgpt.png'), tone: 'dark', row: 'front', glow: 'rgba(34,211,238,0.28)', z: 5, x: 50, y: 28, s: 17, mx: 50, my: 30, ms: 25 },
]

const square = (t: Omit<BrandTile, 'mx' | 'my' | 'ms'>): BrandTile => ({ ...t, mx: t.x, my: t.y, ms: t.s })

// AI 专区首屏：同一组 AI 图标，ChatGPT 在最前面
export const ZONE_AI_TILES: BrandTile[] = [
  square({ key: 'apple', label: 'Apple', src: AI('apple.png'), tone: 'light', row: 'back', glow: 'rgba(148,163,184,0.45)', x: 10, y: 4, s: 12 }),
  square({ key: 'gmail', label: 'Gmail', src: AI('gmail.png'), tone: 'light', row: 'back', glow: 'rgba(239,68,68,0.35)', x: 30, y: 4, s: 12 }),
  square({ key: 'openrouter', label: 'OpenRouter', src: AI('openrouter.png'), tone: 'light', row: 'back', glow: 'rgba(99,102,241,0.45)', x: 50, y: 4, s: 12 }),
  square({ key: 'mastercard', label: 'Mastercard', src: AI('mastercard.png'), tone: 'light', row: 'back', wide: true, glow: 'rgba(249,115,22,0.4)', x: 70, y: 4, s: 12 }),
  square({ key: 'visa', label: 'Visa', src: AI('visa.png'), tone: 'light', row: 'back', wide: true, glow: 'rgba(37,99,235,0.5)', x: 90, y: 4, s: 12 }),
  square({ key: 'google', label: 'Google', src: AI('google.png'), tone: 'light', row: 'back', glow: 'rgba(66,133,244,0.4)', x: 14, y: 21, s: 13 }),
  square({ key: 'google-play', label: 'Google Play', src: AI('google-play.png'), tone: 'light', row: 'back', glow: 'rgba(52,168,83,0.4)', x: 35, y: 22, s: 11 }),
  square({ key: 'code', label: 'Code', src: AI('code.png'), tone: 'light', row: 'back', glow: 'rgba(99,102,241,0.4)', x: 65, y: 22, s: 11 }),
  square({ key: 'youtube', label: 'YouTube', src: AI('youtube.png'), tone: 'light', row: 'back', glow: 'rgba(239,68,68,0.4)', x: 86, y: 21, s: 13 }),
  square({ key: 'claude', label: 'Claude', src: AI('claude.png'), tone: 'light', row: 'front', glow: 'rgba(217,119,87,0.45)', z: 4, x: 23, y: 44, s: 22 }),
  square({ key: 'grok', label: 'Grok', src: AI('grok.png'), tone: 'dark', row: 'front', glow: 'rgba(168,85,247,0.3)', z: 4, x: 77, y: 44, s: 22 }),
  square({ key: 'chatgpt', label: 'ChatGPT', src: AI('chatgpt.png'), tone: 'dark', row: 'front', glow: 'rgba(34,211,238,0.28)', z: 5, x: 50, y: 36, s: 30 }),
]

// 二次元专区首屏：二次元图标，pixiv 在最前面
export const ZONE_ACG_TILES: BrandTile[] = [
  square({ key: 'visa', label: 'Visa', src: ACG('visa.png'), tone: 'light', row: 'back', wide: true, glow: 'rgba(37,99,235,0.5)', x: 22, y: 6, s: 15 }),
  square({ key: 'apple', label: 'Apple', src: ACG('apple.png'), tone: 'light', row: 'back', glow: 'rgba(148,163,184,0.45)', x: 50, y: 4, s: 14 }),
  square({ key: 'mastercard', label: 'Mastercard', src: ACG('mastercard.png'), tone: 'light', row: 'back', wide: true, glow: 'rgba(249,115,22,0.4)', x: 78, y: 6, s: 15 }),
  square({ key: 'gmail', label: 'Gmail', src: ACG('gmail.png'), tone: 'light', row: 'back', glow: 'rgba(239,68,68,0.35)', x: 9, y: 26, s: 13 }),
  square({ key: 'outlook', label: 'Outlook', src: ACG('outlook.png'), tone: 'light', row: 'back', glow: 'rgba(37,99,235,0.4)', x: 91, y: 26, s: 13 }),
  square({ key: 'patreon', label: 'Patreon', src: ACG('patreon.png'), tone: 'light', row: 'front', glow: 'rgba(255,66,77,0.4)', z: 4, x: 23, y: 44, s: 22 }),
  square({ key: 'paypal', label: 'PayPal', src: ACG('paypal.png'), tone: 'light', row: 'front', glow: 'rgba(0,112,224,0.4)', z: 4, x: 77, y: 44, s: 22 }),
  square({ key: 'pixiv', label: 'pixiv', src: ACG('pixiv.png'), tone: 'light', row: 'front', wide: true, glow: 'rgba(14,165,233,0.5)', z: 5, x: 50, y: 36, s: 30 }),
]
