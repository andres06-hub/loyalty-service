package models

type RedemptionTransactions struct {
	Id          string  `gorm:"column:id;primaryKey" json:"id"`
	UserID      string  `gorm:"column:user_id" json:"userId"`
	BranchID    string  `gorm:"column:branch_id" json:"branchId"`
	RewardType  string  `gorm:"column:reward_type" json:"rewardType"`
	RewardValue float64 `gorm:"column:reward_value" json:"rewardValue"`
	CreatedAt   string  `gorm:"column:created_at" json:"createdAt"`
}
