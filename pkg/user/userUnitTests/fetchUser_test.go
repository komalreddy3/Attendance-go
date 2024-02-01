package userUnitTests

import (
	"github.com/golang/mock/gomock"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices/userServiceBean"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestFetchUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	//expectedUsers := []userModels.User{{ID: "1", Username: "user1"}, {ID: "2", Username: "user2"}}
	expectedUsers := []userServiceBean.CustomUserInfo{{Username: "user1"}, {Username: "user2"}}
	mockUserRepo.EXPECT().FetchUser("student").Return([]userModels.User{
		{ID: "1", Username: "user1"},
		{ID: "2", Username: "user2"},
	})

	// Create your UserServiceImpl with the mock repository
	userService := userServices.NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	users := userService.FetchUser("student")

	// Assertions
	assert.Equal(t, len(expectedUsers), len(users))
	assert.Equal(t, expectedUsers, users)

}
