package domain

import "strings"

// 分销商可代理的专区范围。
// 专区由一级分类的 slug 界定，必须与前台 frontend/user/src/utils/businessLines.ts 保持一致。
const (
	ZoneScopeAll  = "all"
	ZoneScopeAI   = "ai"
	ZoneScopeGame = "game"
)

// ZoneRootSlugs 每个专区对应的一级分类 slug（小写）。
var ZoneRootSlugs = map[string][]string{
	ZoneScopeAI:   {"ai", "ai-services", "ai-tools"},
	ZoneScopeGame: {"game-topup", "topup", "acg", "game"},
}

// NormalizeZoneScope 归一化专区范围；空值视为 all，非法值返回 false。
func NormalizeZoneScope(raw string) (string, bool) {
	value := strings.ToLower(strings.TrimSpace(raw))
	switch value {
	case "", ZoneScopeAll:
		return ZoneScopeAll, true
	case ZoneScopeAI, ZoneScopeGame:
		return value, true
	default:
		return "", false
	}
}

// EffectiveZoneScope 返回分销商的专区范围，历史数据为空时按 all 处理。
func (p *Profile) EffectiveZoneScope() string {
	if p == nil {
		return ZoneScopeAll
	}
	scope, ok := NormalizeZoneScope(p.ZoneScope)
	if !ok {
		return ZoneScopeAll
	}
	return scope
}
