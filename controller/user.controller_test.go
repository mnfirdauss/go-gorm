package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/mnfirdauss/go-gorm/model"
	"github.com/mnfirdauss/go-gorm/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(t *testing.M) {

}

// func TestGetUser(t *testing.T) {
// 	mock, _ := conf.InitMockDB()
// 	mockDB := *mock
// 	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")
// 	mockRow := sqlmock.NewRows([]string{
// 		"username", "password", "address",
// 	})
// 	mockRow.AddRow("daus", "test", "jakarta")
// 	mockDB.ExpectQuery(query).WillReturnRows(mockRow)

// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
// 	q := req.URL.Query()
// 	q.Add("id", "1")
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	controller := Controller{}
// 	controller.GetUser(c)
// 	var users []model.User
// 	json.Unmarshal(rec.Body.Bytes(), &users)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, "daus", users[0].Username)
// }

func TestCreateUser(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "localhost:9000/book" // 3rd party API
			httpmock.RegisterResponder(http.MethodGet, url, httpmock.NewStringResponder(http.StatusOK, `{"message":"success}`))

			data := model.User{
				Username: "daus",
			}
			userRepository.Mock.On("CreateUser", &data).Return(errors.New("error"))

			e := echo.New()

			bData, _ := json.Marshal(data)
			req := httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(bData))
			req.Header.Set("content-type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			controller := Controller{}
			controller.CreateUser(c)
			assert.Equal(t, http.StatusOK, rec.Code)

			var resultJSON map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &resultJSON)

			expectResult := map[string]interface{}{
				"message": "success",
			}
			assert.Equal(t, expectResult, resultJSON)

		})
	}
}
