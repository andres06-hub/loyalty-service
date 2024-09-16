package domain

type RewardsDto struct {
	UserID         string `json:"userId" validate:"required,uuid"`
	BranchID       string `json:"branchId" validate:"required,uuid"`
	PurchaseAmount int64  `json:"purchaseAmount" validate:"required,number"`
}

type RewardsRedeemDto struct {
	UserID      string  `json:"userId" validate:"required,uuid"`
	BranchID    string  `json:"branchId" validate:"required,uuid"`
	RewardType  string  `json:"rewardType" validate:"required"`
	RewardValue float64 `json:"rewardValue" validate:"required,number"`
}
