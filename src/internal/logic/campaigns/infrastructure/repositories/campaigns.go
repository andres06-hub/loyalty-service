package repositories

import (
	"fmt"

	"github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/models"
	campRpt "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/domain/repositories"
	"gorm.io/gorm"
)

type CampaignsRepository struct {
	db *gorm.DB
}

func NewCampaignsRepository(dbConnection *gorm.DB) campRpt.CampaignsRepository {
	return &CampaignsRepository{
		db: dbConnection,
	}
}

func (c *CampaignsRepository) FindAll() (res []*models.Campaigns, err error) {
	err = c.db.Raw("SELECT * FROM campaigns").Scan(&res).Error
	if err != nil {
		return nil, fmt.Errorf("error finding campaigns")
	}
	return res, nil
}

func (c *CampaignsRepository) FindAllByBranchId(campaingId string) (res []*models.Campaigns, err error) {
	err = c.db.Raw("SELECT * FROM campaigns WHERE branch_id = ?", campaingId).Scan(&res).Error
	if err != nil {
		return nil, fmt.Errorf("error finding campaign with branch_id: %s", campaingId)
	}

	if res == nil {
		return nil, fmt.Errorf("campaign not found with branch_id: %s", campaingId)
	}

	return res, nil
}

func (c *CampaignsRepository) FindOneByBranchId(branchId string) (res *models.Campaigns, err error) {
	err = c.db.Raw("SELECT * FROM campaigns WHERE branch_id = ?", branchId).Scan(&res).Error
	if err != nil {
		return nil, fmt.Errorf("error finding campaign with branch_id: %s", branchId)
	}

	if res == nil {
		fmt.Printf("campaign not found with branch_id: %s", branchId)
		return nil, nil
	}

	return res, nil
}

func (c *CampaignsRepository) FindOneByBranchIdAndDates(branchID, nowDate string) (res *models.Campaigns, err error) {
	err = c.db.Where("branch_id = ? AND start_date <= ? AND end_date >= ?", branchID, nowDate, nowDate).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("error finding campaign with branch_id: %s", branchID)
	}
	return res, nil
}

func (c *CampaignsRepository) FindOneById(id string) (res *models.Campaigns, err error) {
	err = c.db.Raw("SELECT * FROM campaigns WHERE id = ?", id).Scan(&res).Error
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("error finding campaign")
	}

	if res == nil {
		return nil, fmt.Errorf("campaign not found")
	}

	return res, nil
}

func (c *CampaignsRepository) Update(data *models.Campaigns) (res *models.Campaigns, err error) {
	err = c.db.Model(&models.Campaigns{}).Where("id = ?", data.Id).Updates(map[string]interface{}{
		"branch_id":    data.BranchId,
		"start_date":   data.StartDate,
		"end_date":     data.EndDate,
		"bonus_type":   data.BonusType,
		"bonus_value":  data.BonusValue,
		"min_purchase": data.MinPurchase,
	}).Error
	if err != nil {
		return nil, fmt.Errorf("error updating campaign")
	}

	return data, nil
}
