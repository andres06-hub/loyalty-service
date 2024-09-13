package models

type Rewards struct {
	Id          string `json:"id"`
	UserID      string `json:"user_id"`
	BrachID     string `json:"branch_id"`
	RewardType  string `json:"reward_type"`
	RewardValue string `json:"reward_value"`
	CreatedAt   string `json:"created_at"`
}
