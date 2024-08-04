package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	test_utils "go-payments/tests"

	"go-payments/internal/expenses/api"
	"go-payments/internal/expenses/entities"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ListResponse struct {
	Data []entities.ExpenseEntityResponse
}

type TestListExpensesSuite struct {
	suite.Suite
}

func TestExpensesApisSuite(test *testing.T) {
	suite.Run(test, &TestListExpensesSuite{})
}

func (suite *TestListExpensesSuite) SetupSuite() {
	test_utils.InitTestDb()
	AddTestExpenses(3)
}

func (suite *TestListExpensesSuite) SetupRouter() *gin.Engine {
	r := gin.Default()
	router := r.Group("api/v1/")
	{
		api.ExpensesRoutes(router)
	}
	return r
}

func (suite *TestListExpensesSuite) SendRequest() *httptest.ResponseRecorder {
	router := suite.SetupRouter()
	resp_writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/expense", nil)
	router.ServeHTTP(resp_writer, req)
	return resp_writer
}

func (suite *TestListExpensesSuite) TestListExpenses() {
	resp := suite.SendRequest()
	var resp_body ListResponse
	err := json.Unmarshal(resp.Body.Bytes(), &resp_body)
	if err != nil {
		log.Fatalf("Error on unmarshal response %v", err)
	}
	log.Printf("resp body %+v", resp_body)
	suite.Equal(200, resp.Result().StatusCode)
	suite.Equal(3, len(resp_body.Data))
	suite.Equal(int64(3), resp_body.Data[0].Id)
	suite.Equal(int64(2), resp_body.Data[1].Id)
	suite.Equal(int64(1), resp_body.Data[2].Id)
}
