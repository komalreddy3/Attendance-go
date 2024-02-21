package attendanceServices

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceAdapter"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"go.uber.org/zap"
	"time"
)

type AttendanceServiceImpl struct {
	attendanceRepository attendanceRepository.AttendanceRepo
	logger               *zap.SugaredLogger
}
type AttendanceService interface {
	HasAttendance(userid string) (int, error)
	FetchClassMapPunch(id int) string
	FetchDay(id int) int
	FetchAttendance(day, month, year int) []int
	FetchPunch(id int) []attendanceServiceBean.PunchRecord
	FetchAttendanceofUser(userid string, month int, year int) []int
	UpdatePunchOut(attedid int, classid int, currentDate time.Time) error
	CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error
	AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error)
	PunchOut(userID string, id int) error
	PunchOutCheck(userid string, classid int) error
	PunchOutNull(userid string) string
	CreatePunchInTeacher(userID string, currentDate time.Time, attendanceID int, className string) error
}

func NewAttendanceServiceImpl(attendanceRepository attendanceRepository.AttendanceRepo, logger *zap.SugaredLogger) *AttendanceServiceImpl {
	return &AttendanceServiceImpl{
		attendanceRepository,
		logger,
	}
}
func (impl AttendanceServiceImpl) CreatePunchInTeacher(userID string, currentDate time.Time, attendanceID int, className string) error {
	return impl.attendanceRepository.CreatePunchInTeacher(userID, currentDate, attendanceID, className)
}
func (impl AttendanceServiceImpl) PunchOutNull(userid string) string {
	return impl.attendanceRepository.PunchOutNull(userid)
}
func (impl AttendanceServiceImpl) HasAttendance(userid string) (int, error) {
	return impl.attendanceRepository.HasAttendance(userid)
}
func (impl AttendanceServiceImpl) FetchAttendance(day, month, year int) []int {
	return attendanceAdapter.StudentRec(impl.attendanceRepository.FetchAttendance(day, month, year))
}
func (impl AttendanceServiceImpl) FetchClassMapPunch(id int) string {
	return impl.attendanceRepository.FetchClassMapPunch(id)
}
func (impl AttendanceServiceImpl) FetchPunch(id int) []attendanceServiceBean.PunchRecord {
	return attendanceAdapter.PunchRec(impl.attendanceRepository.FetchPunch(id))
}
func (impl AttendanceServiceImpl) FetchDay(id int) int {
	return impl.attendanceRepository.FetchDay(id)
}
func (impl AttendanceServiceImpl) FetchAttendanceofUser(userid string, month int, year int) []int {
	return attendanceAdapter.StudentRec(impl.attendanceRepository.FetchAttendanceofUser(userid, month, year))

}
func (impl AttendanceServiceImpl) UpdatePunchOut(attedid int, classid int, currentDate time.Time) error {
	return impl.attendanceRepository.UpdatePunchOut(attedid, classid, currentDate)
}
func (impl AttendanceServiceImpl) CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error {
	return impl.attendanceRepository.CreatePunchIn(userID, currentDate, attendanceID, className)
}
func (impl AttendanceServiceImpl) AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error) {
	return impl.attendanceRepository.AddAttendance(userid, currentDate, newAttendanceID)
}
func (impl AttendanceServiceImpl) PunchOut(userID string, id int) error {
	return impl.attendanceRepository.PunchOut(userID, id)
}
func (impl AttendanceServiceImpl) PunchOutCheck(userid string, classid int) error {
	return impl.attendanceRepository.PunchOutCheck(userid, classid)
}
