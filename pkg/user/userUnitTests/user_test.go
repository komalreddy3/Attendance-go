package userServices

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices/userServiceBean"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestCheckEnrollment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	expectedError := errors.New("some error")
	mockUserRepo.EXPECT().CheckEnrollment("testUserID", "testClass").Return(expectedError)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	err := userService.CheckEnrollment("testUserID", "testClass")

	// Assertions
	assert.EqualError(t, err, expectedError.Error())
}

func TestCheckPunchOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	expectedEnrolledClasses := []string{"ClassA", "ClassB"}
	mockUserRepo.EXPECT().CheckPunchOut("testUserID").Return(expectedEnrolledClasses, nil)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	enrolledClasses, err := userService.CheckPunchOut("testUserID")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedEnrolledClasses, enrolledClasses)
}

func TestFetchClass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	mockUserRepo.EXPECT().FetchClass("ClassA").Return(123)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	classID := userService.FetchClass("ClassA")

	// Assertions
	assert.Equal(t, 123, classID)
}

func TestFetchStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	expectedUserID := "testUserID"
	mockUserRepo.EXPECT().FetchStudent("testUserID").Return(expectedUserID)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	userID := userService.FetchStudent("testUserID")

	// Assertions
	assert.Equal(t, expectedUserID, userID)
}

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
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	users := userService.FetchUser("student")

	// Assertions
	assert.Equal(t, len(expectedUsers), len(users))
	assert.Equal(t, expectedUsers, users)
}

func TestInsertingStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	expectedError := errors.New("some error")
	mockUserRepo.EXPECT().InsertingStudent("testID", "testUsername", "testPassword").Return(expectedError)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	err := userService.InsertingStudent("testID", "testUsername", "testPassword")

	// Assertions
	assert.EqualError(t, err, expectedError.Error())
}

func TestInsertingTeacher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	expectedError := errors.New("some error")
	mockUserRepo.EXPECT().InsertingTeacher("testID", "testUsername", "testPassword").Return(expectedError)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	err := userService.InsertingTeacher("testID", "testUsername", "testPassword")

	// Assertions
	assert.EqualError(t, err, expectedError.Error())
}

func TestInsertClass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	expectedClassID := 123
	mockUserRepo.EXPECT().InsertClass("TestClass").Return(expectedClassID, nil)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	classID, err := userService.InsertClass("TestClass")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedClassID, classID)
}

func TestInsertClassMap(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	expectedError := errors.New("some error")
	mockUserRepo.EXPECT().InsertClassMap("testID", 123).Return(expectedError)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	err := userService.InsertClassMap("testID", 123)

	// Assertions
	assert.EqualError(t, err, expectedError.Error())
}

func TestClassMappingTeacher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock UserRepository
	mockUserRepo := userRepository.NewMockUserRepo(ctrl)

	// Set up expectations and behaviors
	expectedClassIDMap := map[string]int{"ClassA": 1, "ClassB": 2}
	mockUserRepo.EXPECT().ClassMappingTeacher([]string{"ClassA", "ClassB"}).Return(expectedClassIDMap, nil)

	// Create your UserServiceImpl with the mock repository
	userService := NewUserServiceImpl(mockUserRepo, zap.NewNop().Sugar())

	// Call the method in your service
	classIDMap, err := userService.ClassMappingTeacher([]string{"ClassA", "ClassB"})

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedClassIDMap, classIDMap)
}
