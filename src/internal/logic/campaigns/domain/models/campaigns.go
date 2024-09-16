package models

import "time"

type Campaigns struct {
	Id          string    `gorm:"column:id;primaryKey" json:"id"`
	BranchId    string    `gorm:"column:branch_id" json:"branchId"`
	StartDate   time.Time `gorm:"column:start_date" json:"startDate"`
	EndDate     time.Time `gorm:"column:end_date" json:"endDate"`
	BonusType   string    `gorm:"column:bonus_type" json:"bonusType"`
	BonusValue  float64   `gorm:"column:bonus_value" json:"bonusValue"`
	MinPurchase float64   `gorm:"column:min_purchase" json:"minPurchase"`
	CreatedAt   string    `gorm:"column:created_at" json:"createdAt"`
}
