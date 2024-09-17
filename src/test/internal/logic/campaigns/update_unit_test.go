package campaigns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/andres06-hub/loyalty-service/src/internal/config"
	uc "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/application/update"
	cm "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/infrastructure/handlers"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUpdateCampaign__Unit(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	startDate, _ := time.Parse("2006-01-02", "2024-09-15")
	endDate, _ := time.Parse("2006-01-02", "2024-09-20")

	t.Run("GetCampaignsHandler | BranchId - Success", func(t *testing.T) {
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id", "branch_id", "start_date", "end_date", "bonus_type", "bonus_value", "min_purchase", "created_at"}).
				AddRow("1", "2", startDate, endDate, "double", 2.0, 1.0, "2024-09-14"))

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE \"campaigns\" SET \"bonus_type\"=$1,\"bonus_value\"=$2,\"branch_id\"=$3,\"end_date\"=$4,\"min_purchase\"=$5,\"start_date\"=$6 WHERE id = $7")).
			WithArgs("double", 2.0, "2", endDate, 1.0, startDate, "1").
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		assr := assert.New(t)
		svcCtx := svc.NewServiceContext(config.Config{}, db)

		var body uc.UpdateCampaignDto = uc.UpdateCampaignDto{
			StartDate:  "2024-09-15",
			EndDate:    "2024-09-20",
			BonusType:  "double",
			BonusValue: 2.0,
		}

		bodyBytes, _ := json.Marshal(&body)

		mockReq := httptest.NewRequest("PUT", "/api/campaigns/bdbc9716-217b-427e-80d0-a1e5f09bd3c4", bytes.NewReader(bodyBytes))
		rr := httptest.NewRecorder()

		cm.UpdateCampaignHandler(svcCtx).ServeHTTP(rr, mockReq)

		assr.Equal(http.StatusOK, rr.Code)

		var response map[string]interface{}
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		assr.NoError(err)

		assr.Contains(response, "timestamp")
		assr.Contains(response, "data")
		assr.Contains(response, "message")
		assr.Equal("campaign update", response["message"])

		campaigns, ok := response["data"]
		assr.True(ok)

		dataExpected := map[string]interface{}{
			"id":          "1",
			"branchId":    "2",
			"startDate":   "2024-09-15T00:00:00Z",
			"endDate":     "2024-09-20T00:00:00Z",
			"bonusType":   "double",
			"bonusValue":  2.0,
			"createdAt":   "2024-09-14",
			"minPurchase": 1.0,
		}
		fmt.Println("=>", campaigns)
		assr.Equal(dataExpected, campaigns)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
