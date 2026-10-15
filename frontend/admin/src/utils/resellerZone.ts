export type ResellerZoneScope = 'all' | 'ai' | 'game'

export const normalizeResellerZoneScope = (value: unknown): ResellerZoneScope =>
  value === 'ai' || value === 'game' ? value : 'all'
