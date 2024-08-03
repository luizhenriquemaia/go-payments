package tests

import (
	"log"
	"net/http"
	"testing"

	test_utils "go-payments/tests"

	"github.com/stretchr/testify/suite"
)

type TestListExpensesSuite struct {
	suite.Suite
}

func TestExpensesApisSuite(test *testing.T) {
	suite.Run(test, &TestListExpensesSuite{})
	result := AddTestExpenses(3)
	log.Printf("added expenses %+v", result)
}

func (suite *TestListExpensesSuite) SetupSuite() {
	test_utils.InitTestDb()
}

func (suite *TestListExpensesSuite) TestListExpenses() {
	resp, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Printf("request error %v", err)
	}
	log.Printf("response %+v", resp)

	suite.Equal(1, 2)
}
