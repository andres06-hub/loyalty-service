package models

type Purchases struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	BranchId       string `json:"branch_id"`
	PurchaseAmount string `json:"purchase_amount"`
	RewardEarned   string `json:"reward_earned"`
	RewardType     string `json:"reward_type"`
	CampaignId     string `json:"campaign_id"`
	CreatedAt      string `json:"created_at"`
}
