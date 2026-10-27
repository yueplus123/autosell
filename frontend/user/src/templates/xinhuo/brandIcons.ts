/**
 * 官网首屏和两个专区首屏展示的业务图标（public/xinhuo/icons/ai、acg 下）。
 * dark：放在黑色砖块上（黑色线稿 logo 会反白）；wide：横向 logo（VISA、万事达、pixiv）用更宽的显示区域。
 */
export interface BrandIcon {
  key: string
  label: string
  src: string
  dark?: boolean
  wide?: boolean
}

const AI_BASE = '/xinhuo/icons/ai/'
const ACG_BASE = '/xinhuo/icons/acg/'

export const AI_ICONS: BrandIcon[] = [
  { key: 'chatgpt', label: 'ChatGPT', src: `${AI_BASE}chatgpt.png`, dark: true },
  { key: 'claude', label: 'Claude', src: `${AI_BASE}claude.png` },
  { key: 'grok', label: 'Grok', src: `${AI_BASE}grok.png`, dark: true },
  { key: 'google', label: 'Google', src: `${AI_BASE}google.png` },
  { key: 'openrouter', label: 'OpenRouter', src: `${AI_BASE}openrouter.png` },
  { key: 'code', label: 'Code', src: `${AI_BASE}code.png` },
  { key: 'apple', label: 'Apple', src: `${AI_BASE}apple.png` },
  { key: 'google-play', label: 'Google Play', src: `${AI_BASE}google-play.png` },
  { key: 'youtube', label: 'YouTube', src: `${AI_BASE}youtube.png` },
  { key: 'gmail', label: 'Gmail', src: `${AI_BASE}gmail.png` },
  { key: 'visa', label: 'Visa', src: `${AI_BASE}visa.png`, wide: true },
  { key: 'mastercard', label: 'Mastercard', src: `${AI_BASE}mastercard.png`, wide: true },
]

export const ACG_ICONS: BrandIcon[] = [
  { key: 'pixiv', label: 'pixiv', src: `${ACG_BASE}pixiv.png`, wide: true },
  { key: 'patreon', label: 'Patreon', src: `${ACG_BASE}patreon.png` },
  { key: 'paypal', label: 'PayPal', src: `${ACG_BASE}paypal.png` },
  { key: 'visa', label: 'Visa', src: `${ACG_BASE}visa.png`, wide: true },
  { key: 'mastercard', label: 'Mastercard', src: `${ACG_BASE}mastercard.png`, wide: true },
  { key: 'apple', label: 'Apple', src: `${ACG_BASE}apple.png` },
  { key: 'gmail', label: 'Gmail', src: `${ACG_BASE}gmail.png` },
  { key: 'outlook', label: 'Outlook', src: `${ACG_BASE}outlook.png` },
]
