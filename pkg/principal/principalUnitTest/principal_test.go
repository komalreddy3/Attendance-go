package principalUnitTest

import (
	"github.com/golang/mock/gomock"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/principal/principalRepository"
	"github.com/komalreddy3/Attendance-go/pkg/principal/principalServices"
	"github.com/komalreddy3/Attendance-go/pkg/principal/principalServices/principalServiceBean"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestGetTeacherAttendance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPrincipalRepo := principalRepository.NewMockPrincipalRepo(ctrl)
	//impl := principalServices.PrincipalServiceImpl{
	//	//principalRepository: mockPrincipalRepo,
	//	//logger: zap.NewNop().Sugar(),
	//}

	impl := principalServices.NewPrincipalServiceImpl(mockPrincipalRepo, zap.NewNop().Sugar())
	// Set up expectations and behaviors for FetchAttendance
	teacherID := "teacher1"
	month := 1
	year := 2024
	mockPrincipalRepo.EXPECT().FetchAttendance(teacherID, month, year).Return([]int{1, 2, 3})

	// Set up expectations and behaviors for FetchDay
	mockPrincipalRepo.EXPECT().FetchDay(1).Return(1)
	mockPrincipalRepo.EXPECT().FetchDay(2).Return(2)
	mockPrincipalRepo.EXPECT().FetchDay(3).Return(3)

	// Set up expectations and behaviors for FetchPunch
	mockPrincipalRepo.EXPECT().FetchPunch(1).Return([]attendanceServiceBean.PunchRecord{{ID: 11, PunchIn: "2022-01-01T09:00:00Z", PunchOut: "2022-01-01T17:00:00Z"}})
	mockPrincipalRepo.EXPECT().FetchPunch(2).Return([]attendanceServiceBean.PunchRecord{{ID: 21, PunchIn: "2022-01-02T09:00:00Z", PunchOut: "2022-01-02T17:00:00Z"}})
	mockPrincipalRepo.EXPECT().FetchPunch(3).Return([]attendanceServiceBean.PunchRecord{{ID: 31, PunchIn: "2022-01-03T09:00:00Z", PunchOut: "2022-01-03T17:00:00Z"}})

	// Set up expectations and behaviors for ClassMapAttendancePunch
	mockPrincipalRepo.EXPECT().ClassMapAttendancePunch(11).Return("ClassA")
	mockPrincipalRepo.EXPECT().ClassMapAttendancePunch(21).Return("ClassB")
	mockPrincipalRepo.EXPECT().ClassMapAttendancePunch(31).Return("ClassC")

	response := impl.GetTeacherAttendance(teacherID, month, year)

	expectedResponse := principalServiceBean.TeacherAttendanceResponse{
		ID:    teacherID,
		Month: month,
		Year:  year,
		Attendance: map[int][]principalServiceBean.AttendanceEntry{
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

	assert.Equal(t, expectedResponse, response)
}
