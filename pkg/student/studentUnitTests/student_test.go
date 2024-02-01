package studentServices

// Import necessary packages for testing
import (
	"github.com/golang/mock/gomock"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/student/studentRepository"
	"github.com/komalreddy3/Attendance-go/pkg/student/studentServices/studentServiceBean"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestStudentPunchIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock StudentRepo
	mockStudentRepo := studentRepository.NewMockStudentRepo(ctrl)

	// Create a StudentServiceImpl with the mock repository
	studentService := NewStudentServiceImpl(mockStudentRepo, zap.NewNop().Sugar())

	// Set up expectations and behaviors for EnrollCheck
	testUserID := "user1"
	testClass := "ClassA"
	mockStudentRepo.EXPECT().EnrollCheck(testUserID, testClass)

	// Set up expectations and behaviors for PunchCheck
	mockStudentRepo.EXPECT().PunchCheck(testUserID).Return("ClassB")

	// Set up expectations and behaviors for FetchAttendance
	mockStudentRepo.EXPECT().FetchAttendance(testUserID).Return(0, nil)

	// Set up expectations and behaviors for FetchClass
	mockStudentRepo.EXPECT().FetchClass("ClassB").Return(2)

	// Set up expectations and behaviors for PunchOutCheck
	mockStudentRepo.EXPECT().PunchOutCheck(testUserID, 2).Return(nil)

	// Set up expectations and behaviors for CreatePunchIn
	mockStudentRepo.EXPECT().CreatePunchIn(testUserID, gomock.Any(), 0, testClass).Return(nil).AnyTimes()

	// Call the method in your service
	studentService.StudentPunchIn(testUserID, testClass)
}

func TestStudentPunchOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock StudentRepo
	mockStudentRepo := studentRepository.NewMockStudentRepo(ctrl)

	// Create a StudentServiceImpl with the mock repository
	studentService := NewStudentServiceImpl(mockStudentRepo, zap.NewNop().Sugar())

	// Set up expectations and behaviors for EnrollCheck
	testUserID := "user1"
	testClass := "ClassA"
	mockStudentRepo.EXPECT().EnrollCheck(testUserID, testClass)

	// Set up expectations and behaviors for FetchAttendance
	mockStudentRepo.EXPECT().FetchAttendance(testUserID).Return(0, nil)

	// Set up expectations and behaviors for FetchClass
	mockStudentRepo.EXPECT().FetchClass(testClass).Return(1)

	// Set up expectations and behaviors for PunchOutCheck
	mockStudentRepo.EXPECT().PunchOutCheck(testUserID, 1).Return(nil)

	// Set up expectations and behaviors for UpdatePunchOut
	mockStudentRepo.EXPECT().UpdatePunchOut(0, 1, gomock.Any()).Return(nil)

	// Call the method in your service
	studentService.StudentPunchOut(testUserID, testClass)
}

func TestGetStudentAttendanceByMonth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock StudentRepo
	mockStudentRepo := studentRepository.NewMockStudentRepo(ctrl)

	// Create a StudentServiceImpl with the mock repository
	studentService := NewStudentServiceImpl(mockStudentRepo, zap.NewNop().Sugar())

	// Set up expectations and behaviors for FetchAttendanceWithMonth
	testUserID := "user1"
	testMonth := 1
	testYear := 2024
	mockStudentRepo.EXPECT().FetchAttendanceWithMonth(testUserID, testMonth, testYear).Return([]int{1, 2, 3})

	// Set up expectations and behaviors for FetchDay
	mockStudentRepo.EXPECT().FetchDay(1).Return(1)
	mockStudentRepo.EXPECT().FetchDay(2).Return(2)
	mockStudentRepo.EXPECT().FetchDay(3).Return(3)

	// Set up expectations and behaviors for FetchPunch
	mockStudentRepo.EXPECT().FetchPunch(1).Return([]attendanceServiceBean.PunchRecord{{ID: 11, PunchIn: "2022-01-01T09:00:00Z", PunchOut: "2022-01-01T17:00:00Z"}})
	mockStudentRepo.EXPECT().FetchPunch(2).Return([]attendanceServiceBean.PunchRecord{{ID: 21, PunchIn: "2022-01-02T09:00:00Z", PunchOut: "2022-01-02T17:00:00Z"}})
	mockStudentRepo.EXPECT().FetchPunch(3).Return([]attendanceServiceBean.PunchRecord{{ID: 31, PunchIn: "2022-01-03T09:00:00Z", PunchOut: "2022-01-03T17:00:00Z"}})

	// Set up expectations and behaviors for ClassMapAttendancePunch
	mockStudentRepo.EXPECT().ClassMapAttendancePunch(11).Return("ClassA")
	mockStudentRepo.EXPECT().ClassMapAttendancePunch(21).Return("ClassB")
	mockStudentRepo.EXPECT().ClassMapAttendancePunch(31).Return("ClassC")

	// Create your expected response
	expectedResponse := studentServiceBean.TeacherAttendanceResponse{
		ID:    testUserID,
		Month: testMonth,
		Year:  testYear,
		Attendance: map[int][]studentServiceBean.AttendanceEntry{
			1: {
				{Class: "ClassA", FirstPunchIn: "2022-01-01T09:00:00Z", LastPunchOut: "2022-01-01T17:00:00Z"},
			},
			2: {
				{Class: "ClassB", FirstPunchIn: "2022-01-02T09:00:00Z", LastPunchOut: "2022-01-02T17:00:00Z"},
			},
			3: {
				{Class: "ClassC", FirstPunchIn: "2022-01-03T09:00:00Z", LastPunchOut: "2022-01-03T17:00:00Z"},
			},
		},
	}

	// Call the method in your service
	response := studentService.GetStudentAttendanceByMonth(testUserID, testMonth, testYear)

	// Assert that the response matches the expected response
	assert.Equal(t, expectedResponse, response)
}
