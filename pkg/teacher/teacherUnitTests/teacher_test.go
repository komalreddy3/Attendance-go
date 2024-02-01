package teacherServices

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/teacher/teacherRepository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestTeacherPunchIn_SuccessfulPunchIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock TeacherRepo
	mockTeacherRepo := teacherRepository.NewMockTeacherRepo(ctrl)

	// Set up expectations and behaviors
	testUserID := "user1"
	testClass := "ClassA"
	mockTeacherRepo.EXPECT().EnrollCheckTeacher(testUserID, testClass)
	mockTeacherRepo.EXPECT().PunchCheckTeacher(testUserID).Return([]string{"ClassB"})
	mockTeacherRepo.EXPECT().FetchAttendance(testUserID).Return(1, nil)
	mockTeacherRepo.EXPECT().FetchClass(gomock.Any()).AnyTimes().Return(2)
	mockTeacherRepo.EXPECT().PunchOutCheck(testUserID, gomock.Any()).Return(nil)
	mockTeacherRepo.EXPECT().CreatePunchIn(testUserID, gomock.Any(), 1, testClass).Return(nil).AnyTimes()

	// Create your TeacherServiceImpl with the mock repository
	teacherService := NewTeacherServiceImpl(mockTeacherRepo, zap.NewNop().Sugar())

	// Call the method in your service
	teacherService.TeacherPunchIn(testUserID, testClass)
}

func TestTeacherPunchIn_FailurePunchOutCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock TeacherRepo
	mockTeacherRepo := teacherRepository.NewMockTeacherRepo(ctrl)

	// Set up expectations and behaviors
	testUserID := "user1"
	testClass := "ClassA"
	// EnrollCheckTeacher is expected to be called
	mockTeacherRepo.EXPECT().EnrollCheckTeacher(testUserID, testClass)

	// PunchCheckTeacher is expected to be called and return some enrolled classes
	mockTeacherRepo.EXPECT().PunchCheckTeacher(testUserID).Return([]string{"ClassB"})

	// FetchAttendance is expected to be called and return an error
	mockTeacherRepo.EXPECT().FetchAttendance(testUserID).Return(0, errors.New("some error"))

	// Create your TeacherServiceImpl with the mock repository
	teacherService := NewTeacherServiceImpl(mockTeacherRepo, zap.NewNop().Sugar())

	// Call the method in your service
	teacherService.TeacherPunchIn(testUserID, testClass)
	// Add assertions for the expected error
}

func TestTeacherPunchOut_SuccessfulPunchOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock TeacherRepo
	mockTeacherRepo := teacherRepository.NewMockTeacherRepo(ctrl)

	// Set up expectations and behaviors
	userID := "testUserID"
	class := "testClass"
	mockTeacherRepo.EXPECT().EnrollCheckTeacher(userID, class)
	mockTeacherRepo.EXPECT().FetchAttendance(userID).Return(1, nil)
	mockTeacherRepo.EXPECT().FetchClass(class).Return(123)
	mockTeacherRepo.EXPECT().PunchOutCheck(userID, 123).Return(nil)
	mockTeacherRepo.EXPECT().UpdatePunchOut(1, 123, gomock.Any())

	// Create your TeacherServiceImpl with the mock repository
	teacherService := NewTeacherServiceImpl(mockTeacherRepo, zap.NewNop().Sugar())

	// Call the method in your service
	teacherService.TeacherPunchOut(userID, class)
}

func TestTeacherPunchOut_FailurePunchOutCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock TeacherRepo
	mockTeacherRepo := teacherRepository.NewMockTeacherRepo(ctrl)

	// Set up expectations and behaviors
	userID := "testUserID"
	class := "testClass"
	mockTeacherRepo.EXPECT().EnrollCheckTeacher(userID, class)
	mockTeacherRepo.EXPECT().FetchAttendance(userID).Return(1, nil)
	mockTeacherRepo.EXPECT().FetchClass(class).Return(123)
	expectedError := errors.New("PunchOutCheck error")
	mockTeacherRepo.EXPECT().PunchOutCheck(userID, 123).Return(expectedError)

	// Create your TeacherServiceImpl with the mock repository
	teacherService := NewTeacherServiceImpl(mockTeacherRepo, zap.NewNop().Sugar())

	// Call the method in your service
	teacherService.TeacherPunchOut(userID, class)
	// Add assertions for the expected error
}
func TestGetClassAttendance_Successful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock TeacherRepo
	mockTeacherRepo := teacherRepository.NewMockTeacherRepo(ctrl)

	// Set up expectations and behaviors
	class := "testClass"
	day := 1
	month := 1
	year := 2022
	mockTeacherRepo.EXPECT().FetchStudentAttendance(day, month, year).Return([]int{1, 2, 3})
	mockTeacherRepo.EXPECT().FetchDay(gomock.Any()).AnyTimes().Return(1)
	mockTeacherRepo.EXPECT().FetchPunch(gomock.Any()).AnyTimes().Return([]attendanceServiceBean.PunchRecord{
		{ID: 1, UserID: "user1", PunchIn: "2022-01-01T08:00:00Z", PunchOut: "2022-01-01T09:00:00Z"},
		{ID: 2, UserID: "user2", PunchIn: "2022-01-01T09:30:00Z", PunchOut: "2022-01-01T11:30:00Z"},
		{ID: 3, UserID: "user1", PunchIn: "2022-01-01T08:15:00Z", PunchOut: "2022-01-01T11:00:00Z"},
	})
	mockTeacherRepo.EXPECT().FetchStudent(gomock.Any()).AnyTimes().Return("John Doe")
	mockTeacherRepo.EXPECT().ClassMapAttendancePunch(gomock.Any()).AnyTimes().Return("testClass")

	// Create your TeacherServiceImpl with the mock repository
	teacherService := NewTeacherServiceImpl(mockTeacherRepo, zap.NewNop().Sugar())

	// Call the method in your service
	response := teacherService.GetClassAttendance(class, day, month, year)
	fmt.Println(response)
	// Add assertions for the expected response
	assert.NotNil(t, response)
	// Add more assertions based on the expected behavior of your service
}

func TestGetTeacherAttendanceByMonth_Successful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock TeacherRepo
	mockTeacherRepo := teacherRepository.NewMockTeacherRepo(ctrl)

	// Set up expectations and behaviors
	ID := "testID"
	month := 1
	year := 2022
	mockTeacherRepo.EXPECT().FetchAttendanceWithMonth(ID, month, year).Return([]int{1, 2, 3})
	mockTeacherRepo.EXPECT().FetchDay(gomock.Any()).AnyTimes().Return(1)
	mockTeacherRepo.EXPECT().FetchPunch(gomock.Any()).AnyTimes().Return([]attendanceServiceBean.PunchRecord{
		{ID: 1, PunchIn: "2022-01-01T08:00:00Z", PunchOut: "2022-01-01T09:00:00Z"},
		{ID: 2, PunchIn: "2022-01-01T09:30:00Z", PunchOut: "2022-01-01T11:30:00Z"},
		{ID: 3, PunchIn: "2022-01-01T08:15:00Z", PunchOut: "2022-01-01T11:00:00Z"},
	})
	mockTeacherRepo.EXPECT().ClassMapAttendancePunch(gomock.Any()).AnyTimes().Return("testClass")

	// Create your TeacherServiceImpl with the mock repository
	teacherService := NewTeacherServiceImpl(mockTeacherRepo, zap.NewNop().Sugar())

	// Call the method in your service
	response := teacherService.GetTeacherAttendanceByMonth(ID, month, year)
	fmt.Println(response)
	// Add assertions for the expected response
	assert.NotNil(t, response)
	// Add more assertions based on the expected behavior of your service
}
