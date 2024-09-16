package domain

type RewardsAccumulateResponse struct {
	RewardsEarned   float64
	RewardsType     string
	CampaignApplied bool
}

type RewardRedemptionResponse struct {
	RewardType         string  `json:"rewardType"`
	RewardRedeemed     float64 `json:"rewardRedeemed"`
	CurrentRewardValue float64 `json:"currentReward"`
}
