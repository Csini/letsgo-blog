package impl

import (
	"github.com/stretchr/testify/assert"
	sut "impl"
	"testing"

	"entity"
	openapi "generated/openapi"

	model "impl/model"
)

func TestApiStatisticsService_HandleUsers_empty(t *testing.T) {
	var expected []openapi.Statistic
	calculated := sut.HandleUsers(nil, nil, nil)
	assert.Equal(t, expected, calculated)
}

func TestApiStatisticsService_HandleUsers_with_data(t *testing.T) {
	expected := []openapi.Statistic{
		openapi.Statistic{
			Userid: "unittest",
			Amount: openapi.Amount{
				Blog:    2,
				Comment: 1,
			},
		},
		openapi.Statistic{
			Userid: "otheruser",
			Amount: openapi.Amount{
				Blog:    1,
				Comment: 0,
			},
		},
	}
	users := []entity.User{
		entity.User{Username: "unittest"},
		entity.User{Username: "otheruser"},
	}

	blogs := []model.CustomRawResult{
		model.CustomRawResult{Username: "unittest", Amount: 2},
		model.CustomRawResult{Username: "otheruser", Amount: 1},
	}
	comments := []model.CustomRawResult{
		model.CustomRawResult{Username: "unittest", Amount: 1},
		model.CustomRawResult{Username: "supriseuser", Amount: 1},
	}
	calculated := sut.HandleUsers(users, blogs, comments)
	assert.Equal(t, expected, calculated)
}
