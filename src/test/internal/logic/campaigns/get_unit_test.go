package campaigns

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/andres06-hub/loyalty-service/src/internal/config"
	cm "github.com/andres06-hub/loyalty-service/src/internal/logic/campaigns/infrastructure/handlers"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetCampaigns__Unit(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	startDate, _ := time.Parse("2006-01-02", "2024-09-15")
	endDate, _ := time.Parse("2006-01-02", "2024-09-20")

	t.Run("GetCampaignsHandler | BranchId - Success", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "branch_id", "start_date", "end_date", "bonus_type", "bonus_value", "min_purchase", "created_at"}).AddRow("1", "2", startDate, endDate, "double", 2.0, 0.0, "2024-09-14"))
		assr := assert.New(t)
		svcCtx := svc.NewServiceContext(config.Config{}, db)

		mockReq := httptest.NewRequest("GET", "/campaigns?branchId=1", nil)
		rr := httptest.NewRecorder()

		cm.GetCampaignsHandler(svcCtx).ServeHTTP(rr, mockReq)

		assr.Equal(http.StatusOK, rr.Code)

		var response map[string]interface{}
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		assr.NoError(err)

		assr.Contains(response, "timestamp")
		assr.Contains(response, "data")
		assr.Contains(response, "message")
		assr.Equal("campaigns found", response["message"])

		campaigns, ok := response["data"].([]interface{})
		assr.True(ok)

		dataExpected := map[string]interface{}{
			"id":          "1",
			"branchId":    "2",
			"startDate":   "2024-09-15T00:00:00Z",
			"endDate":     "2024-09-20T00:00:00Z",
			"bonusType":   "double",
			"bonusValue":  2.0,
			"createdAt":   "2024-09-14",
			"minPurchase": 0.0,
		}
		assr.Equal(dataExpected, campaigns[0])
	})

	t.Run("GetCampaignsHandler | No BranchId - Success", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "branch_id", "start_date", "end_date", "bonus_type", "bonus_value", "min_purchase", "created_at"}).AddRow("1", "2", startDate, endDate, "double", 2.0, 0.0, "2024-09-14"))
		assr := assert.New(t)
		svcCtx := svc.NewServiceContext(config.Config{}, db)

		mockReq := httptest.NewRequest("GET", "/campaigns", nil)
		rr := httptest.NewRecorder()

		cm.GetCampaignsHandler(svcCtx).ServeHTTP(rr, mockReq)

		assr.Equal(http.StatusOK, rr.Code)

		var response map[string]interface{}
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		assr.NoError(err)

		assr.Contains(response, "timestamp")
		assr.Contains(response, "data")
		assr.Contains(response, "message")
		assr.Equal("campaigns found", response["message"])

		campaigns, ok := response["data"].([]interface{})
		assr.True(ok)

		dataExpected := map[string]interface{}{
			"id":          "1",
			"branchId":    "2",
			"startDate":   "2024-09-15T00:00:00Z",
			"endDate":     "2024-09-20T00:00:00Z",
			"bonusType":   "double",
			"bonusValue":  2.0,
			"createdAt":   "2024-09-14",
			"minPurchase": 0.0,
		}
		assr.Equal(dataExpected, campaigns[0])
	})
}
