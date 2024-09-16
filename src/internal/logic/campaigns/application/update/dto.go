package update

type UpdateCampaignDto struct {
	StartDate  string  `json:"startDate" validate:"required"`
	EndDate    string  `json:"endDate" validate:"required"`
	BonusType  string  `json:"bonusType" validate:"required"`
	BonusValue float64 `json:"bonusValue" validate:"required,number"`
}
