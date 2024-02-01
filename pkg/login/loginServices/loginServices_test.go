package loginServices

import (
	"github.com/golang/mock/gomock"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginRepository"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginRepository/loginModels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

func TestAuthenticateRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoginRepo := loginRepository.NewMockLoginRepo(ctrl)
	mockLoginService := NewMockLoginService(ctrl)
	impl := &LoginServiceImpl{mockLoginRepo, zap.NewNop().Sugar()}

	// Set up expectations for AuthenticateRole
	mockCookie := &http.Cookie{Value: "mocked_cookie_value"}
	mockCheck := "teacher"

	// Mock the NewJWT method
	mockLoginService.EXPECT().NewJWT(gomock.Any(), gomock.Any()).Return(loginModels.JWT{}).AnyTimes()

	// Mock the Validate method
	mockLoginRepo.EXPECT().CheckRole(gomock.Any()).Return(loginModels.Teacher)

	// Call the method in your service
	result := impl.AuthenticateRole(mockCookie, mockCheck)

	// Assert that the result is true as expected
	assert.True(t, result)
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoginRepo := loginRepository.NewMockLoginRepo(ctrl)
	mockLoginService := NewMockLoginService(ctrl)
	impl := &LoginServiceImpl{mockLoginRepo, zap.NewNop().Sugar()}

	// Set up expectations for Login
	mockUsername := "test_user"
	mockPassword := "test_password"
	mockUserRole := "student"

	// Mock the CheckCreds method
	mockLoginRepo.EXPECT().CheckCreds(mockUsername, mockUserRole).Return("user_id", nil)

	// Mock the AuthenticateUser method
	mockLoginRepo.EXPECT().AuthenticateUser(mockUsername, mockPassword, mockUserRole).Return(true)

	// Mock the NewJWT method
	mockLoginService.EXPECT().NewJWT(gomock.Any(), gomock.Any()).Return(loginModels.JWT{}).AnyTimes()

	// Call the method in your service
	token := impl.Login(mockUsername, mockPassword, mockUserRole)

	// Assert that the returned token is not nil
	assert.NotNil(t, token)

}
