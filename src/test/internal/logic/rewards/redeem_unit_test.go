package rewards

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/andres06-hub/loyalty-service/src/internal/config"
	"github.com/andres06-hub/loyalty-service/src/internal/svc"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	rw "github.com/andres06-hub/loyalty-service/src/internal/logic/rewards/infrastructure/handlers"
)

func TestRedeemRewards__Unit(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	t.Run("RedeemRewardHandler | Success", func(t *testing.T) {
		// Mocking for Rewards
		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "branch_id", "reward_type", "reward_value", "created_at"}).
				AddRow("1", "2", "3", "points", 100, "2024-09-16"))

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE \"rewards\" SET \"reward_value\"=$1 WHERE id = $2")).
			WithArgs(90.0, "1").
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO \"redemption_transactions\" (\"id\",\"user_id\",\"branch_id\",\"reward_type\",\"reward_value\",\"created_at\") VALUES ($1,$2,$3,$4,$5,$6)")).
			WithArgs(
				sqlmock.AnyArg(),
				"1",
				"1",
				"points",
				10.0,
				sqlmock.AnyArg(),
			).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		assr := assert.New(t)
		svcCtx := svc.NewServiceContext(config.Config{}, db)

		body := map[string]interface{}{
			"userId":      "1",
			"branchId":    "1",
			"rewardType":  "points",
			"rewardValue": 10000,
		}

		bodyBytes, _ := json.Marshal(body)

		mockReq := httptest.NewRequest("POST", "/rewards", bytes.NewReader(bodyBytes))
		rr := httptest.NewRecorder()

		rw.RedeemRewardHandler(svcCtx).ServeHTTP(rr, mockReq)

		assr.Equal(http.StatusOK, rr.Code)

		var response map[string]interface{}
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		assr.NoError(err)

		assr.Contains(response, "timestamp")
		assr.Contains(response, "data")
		assr.Contains(response, "message")
		assr.Equal("points redeemd", response["message"])

		rewards, ok := response["data"]
		assr.True(ok)

		dataExpected := map[string]interface{}{
			"rewardType":     "points",
			"rewardRedeemed": 10.0,
			"currentReward":  90.0,
		}

		assr.Equal(dataExpected, rewards)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
