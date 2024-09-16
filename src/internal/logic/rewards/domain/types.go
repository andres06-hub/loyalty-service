package domain

type RewardsAccumulateResponse struct {
	RewardEarned    float64 `json:"rewardEarned"`
	RewardType      string  `json:"rewardType"`
	CampaignApplied bool    `json:"campaignApplied"`
}

type RewardRedemptionResponse struct {
	RewardType     string  `json:"rewardType"`
	RewardRedeemed float64 `json:"rewardRedeemed"`
	CurrentReward  float64 `json:"currentReward"`
}
