package update

type UpdateCampaignDto struct {
	StartDate  string  `json:"startDate"`
	EndDate    string  `json:"endDate"`
	BonusType  string  `json:"bonusType"`
	BonusValue float64 `json:"bonusValue"`
}

type Params struct {
	CampaignId string `path:"campaignId"`
}
