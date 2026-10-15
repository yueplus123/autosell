package domain

import (
	"strings"
	"time"

	userdomain "github.com/dujiao-next/internal/modules/identity/user/domain"
	"github.com/dujiao-next/internal/shared/money"
)

// Profile 分销商资料。
type Profile struct {
	ID                   uint         `gorm:"primarykey" json:"id"`
	UserID               uint         `gorm:"not null;uniqueIndex" json:"user_id"`
	Status               string       `gorm:"type:varchar(32);index;not null;default:'pending_review'" json:"status"`
	ApplyReason          string       `gorm:"type:text" json:"apply_reason,omitempty"`
	RejectReason         string       `gorm:"type:text" json:"reject_reason,omitempty"`
	DefaultMarkupPercent money.Amount `gorm:"type:decimal(10,2);not null;default:0" json:"default_markup_percent"`
	MaxMarkupPercent     money.Amount `gorm:"type:decimal(10,2);not null;default:0" json:"max_markup_percent"`
	SettlementStatus     string       `gorm:"type:varchar(32);index;not null;default:'normal'" json:"settlement_status"`
	// ZoneScope 可代理的专区：all（全部）/ ai / game，见下方 ZoneScope* 常量
	ZoneScope  string     `gorm:"type:varchar(16);not null;default:'all'" json:"zone_scope"`
	ReviewedBy *uint      `gorm:"index" json:"reviewed_by,omitempty"`
	ReviewedAt *time.Time `gorm:"index" json:"reviewed_at,omitempty"`
	CreatedAt  time.Time  `gorm:"index" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"index" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"index" json:"-"`

	User *userdomain.User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Profile) TableName() string { return "reseller_profiles" }

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
