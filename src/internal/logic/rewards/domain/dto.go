package domain

type RewardsDto struct {
	UserID         string `json:"userId" validate:"required,uuid"`
	BranchID       string `json:"branchId" validate:"required,uuid"`
	PurchaseAmount int64  `json:"purchaseAmount" validate:"required,number"`
}
