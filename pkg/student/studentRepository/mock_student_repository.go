// Code generated by MockGen. DO NOT EDIT.
// Source: studentRepository/studentRepository.go

// Package studentRepository is a generated GoMock package.
package studentRepository

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	attendanceServiceBean "github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
)

// MockStudentRepo is a mock of StudentRepo interface.
type MockStudentRepo struct {
	ctrl     *gomock.Controller
	recorder *MockStudentRepoMockRecorder
}

// MockStudentRepoMockRecorder is the mock recorder for MockStudentRepo.
type MockStudentRepoMockRecorder struct {
	mock *MockStudentRepo
}

// NewMockStudentRepo creates a new mock instance.
func NewMockStudentRepo(ctrl *gomock.Controller) *MockStudentRepo {
	mock := &MockStudentRepo{ctrl: ctrl}
	mock.recorder = &MockStudentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentRepo) EXPECT() *MockStudentRepoMockRecorder {
	return m.recorder
}

// AddAttendance mocks base method.
func (m *MockStudentRepo) AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAttendance", userid, currentDate, newAttendanceID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAttendance indicates an expected call of AddAttendance.
func (mr *MockStudentRepoMockRecorder) AddAttendance(userid, currentDate, newAttendanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAttendance", reflect.TypeOf((*MockStudentRepo)(nil).AddAttendance), userid, currentDate, newAttendanceID)
}

// ClassMapAttendancePunch mocks base method.
func (m *MockStudentRepo) ClassMapAttendancePunch(id int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassMapAttendancePunch", id)
	ret0, _ := ret[0].(string)
	return ret0
}

// ClassMapAttendancePunch indicates an expected call of ClassMapAttendancePunch.
func (mr *MockStudentRepoMockRecorder) ClassMapAttendancePunch(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassMapAttendancePunch", reflect.TypeOf((*MockStudentRepo)(nil).ClassMapAttendancePunch), id)
}

// CreatePunchIn mocks base method.
func (m *MockStudentRepo) CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePunchIn", userID, currentDate, attendanceID, className)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePunchIn indicates an expected call of CreatePunchIn.
func (mr *MockStudentRepoMockRecorder) CreatePunchIn(userID, currentDate, attendanceID, className interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePunchIn", reflect.TypeOf((*MockStudentRepo)(nil).CreatePunchIn), userID, currentDate, attendanceID, className)
}

// EnrollCheck mocks base method.
func (m *MockStudentRepo) EnrollCheck(userid, class string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnrollCheck", userid, class)
}

// EnrollCheck indicates an expected call of EnrollCheck.
func (mr *MockStudentRepoMockRecorder) EnrollCheck(userid, class interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrollCheck", reflect.TypeOf((*MockStudentRepo)(nil).EnrollCheck), userid, class)
}

// FetchAttendance mocks base method.
func (m *MockStudentRepo) FetchAttendance(userid string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttendance", userid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAttendance indicates an expected call of FetchAttendance.
func (mr *MockStudentRepoMockRecorder) FetchAttendance(userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttendance", reflect.TypeOf((*MockStudentRepo)(nil).FetchAttendance), userid)
}

// FetchAttendanceWithMonth mocks base method.
func (m *MockStudentRepo) FetchAttendanceWithMonth(id string, month, year int) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttendanceWithMonth", id, month, year)
	ret0, _ := ret[0].([]int)
	return ret0
}

// FetchAttendanceWithMonth indicates an expected call of FetchAttendanceWithMonth.
func (mr *MockStudentRepoMockRecorder) FetchAttendanceWithMonth(id, month, year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttendanceWithMonth", reflect.TypeOf((*MockStudentRepo)(nil).FetchAttendanceWithMonth), id, month, year)
}

// FetchClass mocks base method.
func (m *MockStudentRepo) FetchClass(enrolledClass string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchClass", enrolledClass)
	ret0, _ := ret[0].(int)
	return ret0
}

// FetchClass indicates an expected call of FetchClass.
func (mr *MockStudentRepoMockRecorder) FetchClass(enrolledClass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchClass", reflect.TypeOf((*MockStudentRepo)(nil).FetchClass), enrolledClass)
}

// FetchDay mocks base method.
func (m *MockStudentRepo) FetchDay(id int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDay", id)
	ret0, _ := ret[0].(int)
	return ret0
}

// FetchDay indicates an expected call of FetchDay.
func (mr *MockStudentRepoMockRecorder) FetchDay(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDay", reflect.TypeOf((*MockStudentRepo)(nil).FetchDay), id)
}

// FetchPunch mocks base method.
func (m *MockStudentRepo) FetchPunch(id int) []attendanceServiceBean.PunchRecord {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPunch", id)
	ret0, _ := ret[0].([]attendanceServiceBean.PunchRecord)
	return ret0
}

// FetchPunch indicates an expected call of FetchPunch.
func (mr *MockStudentRepoMockRecorder) FetchPunch(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPunch", reflect.TypeOf((*MockStudentRepo)(nil).FetchPunch), id)
}

// PunchCheck mocks base method.
func (m *MockStudentRepo) PunchCheck(userid string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchCheck", userid)
	ret0, _ := ret[0].(string)
	return ret0
}

// PunchCheck indicates an expected call of PunchCheck.
func (mr *MockStudentRepoMockRecorder) PunchCheck(userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchCheck", reflect.TypeOf((*MockStudentRepo)(nil).PunchCheck), userid)
}

// PunchOutCheck mocks base method.
func (m *MockStudentRepo) PunchOutCheck(userid string, classid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchOutCheck", userid, classid)
	ret0, _ := ret[0].(error)
	return ret0
}

// PunchOutCheck indicates an expected call of PunchOutCheck.
func (mr *MockStudentRepoMockRecorder) PunchOutCheck(userid, classid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchOutCheck", reflect.TypeOf((*MockStudentRepo)(nil).PunchOutCheck), userid, classid)
}

// UpdatePunchOut mocks base method.
func (m *MockStudentRepo) UpdatePunchOut(punchid, classid int, currentDate time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePunchOut", punchid, classid, currentDate)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePunchOut indicates an expected call of UpdatePunchOut.
func (mr *MockStudentRepoMockRecorder) UpdatePunchOut(punchid, classid, currentDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePunchOut", reflect.TypeOf((*MockStudentRepo)(nil).UpdatePunchOut), punchid, classid, currentDate)
}