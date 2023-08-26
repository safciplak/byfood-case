package users

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestEntitySuite struct {
	suite.Suite

	userEntity User
}

// SetupTest sets up often used objects
func (test *TestEntitySuite) SetupTest() {

	test.userEntity = User{
		ID:   1,
		Name: "test user",
	}
}

// TestEntityTestSuite Runs the testsuite
func TestEntityTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestEntitySuite))
}

func (test *TestEntitySuite) TestGetUserProperties() {
	result := test.userEntity

	test.Equal(result.ID, int64(1))
	test.Equal(result.Name, "test user")
}

func (test *TestEntitySuite) TestGetCampaignPropertiesFailure() {
	result := test.userEntity

	test.NotEqual(result.ID, int64(2))
}
