package models

import "database/sql"

type Purchases struct {
	Id             string         `json:"id"`
	UserId         string         `json:"user_id"`
	BranchId       string         `json:"branch_id"`
	PurchaseAmount float64        `json:"purchase_amount"`
	RewardEarned   float64        `json:"reward_earned"`
	RewardType     string         `json:"reward_type"`
	CampaignId     sql.NullString `json:"campaign_id"`
	CreatedAt      string         `json:"created_at"`
}
