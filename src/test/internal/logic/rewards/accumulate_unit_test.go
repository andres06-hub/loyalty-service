package rewards

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/andres06-hub/loyalty-service/src/internal/config"
	rw "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/infrastructure/handlers"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAccumulateRewards__Unit(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	startDate, _ := time.Parse("2006-01-02", "2024-09-15")
	endDate, _ := time.Parse("2006-01-02", "2024-09-20")

	t.Run("AccumulateRewardHandler | Success - Update", func(t *testing.T) {
		// Mocking for Branchs
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id", "merchant_id", "name", "location", "created_at"}).
				AddRow("1", "2", "Texaco Sucursal 1", "Calle 10 #45-67, Bogota", "2024-09-16"))

		// Mocking for Campaign
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id", "branch_id", "start_date", "end_date", "bonus_type", "bonus_value", "min_purchase", "created_at"}).
				AddRow("1", "2", startDate, endDate, "double", 2.0, 1.0, "2024-09-14"))

		// Mocking for Rewards
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "branch_id", "reward_type", "reward_value", "created_at"}).
				AddRow("1", "2", "3", "points", 100, "2024-09-16"))

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE \"rewards\" SET \"reward_value\"=$1 WHERE id = $2")).
			WithArgs(120.0, "1").
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"purchases\" (\"id\",\"user_id\",\"branch_id\",\"purchase_amount\",\"reward_earned\",\"reward_type\",\"campaign_id\",\"created_at\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)")).
			WithArgs(
				sqlmock.AnyArg(),
				"1",
				"1",
				10000.0,
				20.0,
				"points",
				"1",
				sqlmock.AnyArg(),
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		assr := assert.New(t)
		svcCtx := svc.NewServiceContext(config.Config{}, db)

		body := map[string]interface{}{
			"userId":         "1",
			"branchId":       "1",
			"purchaseAmount": 10000,
		}

		bodyBytes, _ := json.Marshal(body)

		mockReq := httptest.NewRequest("POST", "/rewards", bytes.NewReader(bodyBytes))
		rr := httptest.NewRecorder()

		rw.AccumulateRewardHandler(svcCtx).ServeHTTP(rr, mockReq)

		assr.Equal(http.StatusOK, rr.Code)

		var response map[string]interface{}
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		assr.NoError(err)

		assr.Contains(response, "timestamp")
		assr.Contains(response, "data")
		assr.Contains(response, "message")
		assr.Equal("Reward successfully accumulated", response["message"])

		rewards, ok := response["data"]
		assr.True(ok)

		dataExpected := map[string]interface{}{
			"rewardEarned":    20.0,
			"rewardType":      "points",
			"campaignApplied": true,
		}
		assr.Equal(dataExpected, rewards)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("AccumulateRewardHandler | Success - Create", func(t *testing.T) {
		// Mocking for Branchs
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id", "merchant_id", "name", "location", "created_at"}).
				AddRow("1", "2", "Texaco Sucursal 1", "Calle 10 #45-67, Bogota", "2024-09-16"))

		// Mocking for Campaign
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id", "branch_id", "start_date", "end_date", "bonus_type", "bonus_value", "min_purchase", "created_at"}).
				AddRow("1", "2", startDate, endDate, "double", 2.0, 1.0, "2024-09-14"))

			// Mocking for Rewards
		mock.ExpectQuery("SELECT").
			WillReturnError(sql.ErrNoRows)

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"rewards\" (\"id\",\"user_id\",\"branch_id\",\"reward_type\",\"reward_value\",\"created_at\") VALUES ($1,$2,$3,$4,$5,$6)")).
			WithArgs(
				sqlmock.AnyArg(),
				"1",
				"1",
				"points",
				20.0,
				sqlmock.AnyArg(),
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"purchases\" (\"id\",\"user_id\",\"branch_id\",\"purchase_amount\",\"reward_earned\",\"reward_type\",\"campaign_id\",\"created_at\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)")).
			WithArgs(
				sqlmock.AnyArg(),
				"1",
				"1",
				10000.0,
				20.0,
				"points",
				"1",
				sqlmock.AnyArg(),
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		assr := assert.New(t)
		svcCtx := svc.NewServiceContext(config.Config{}, db)

		body := map[string]interface{}{
			"userId":         "1",
			"branchId":       "1",
			"purchaseAmount": 10000,
		}

		bodyBytes, _ := json.Marshal(body)

		mockReq := httptest.NewRequest("POST", "/rewards", bytes.NewReader(bodyBytes))
		rr := httptest.NewRecorder()

		rw.AccumulateRewardHandler(svcCtx).ServeHTTP(rr, mockReq)

		assr.Equal(http.StatusOK, rr.Code)

		var response map[string]interface{}
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		assr.NoError(err)

		assr.Contains(response, "timestamp")
		assr.Contains(response, "data")
		assr.Contains(response, "message")
		assr.Equal("Reward successfully accumulated", response["message"])

		rewards, ok := response["data"]
		assr.True(ok)

		dataExpected := map[string]interface{}{
			"rewardEarned":    20.0,
			"rewardType":      "points",
			"campaignApplied": true,
		}
		assr.Equal(dataExpected, rewards)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
