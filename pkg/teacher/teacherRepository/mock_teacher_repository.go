// Code generated by MockGen. DO NOT EDIT.
// Source: teacherRepository/teacherRepo.go

// Package teacherRepository is a generated GoMock package.
package teacherRepository

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	attendanceServiceBean "github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
)

// MockTeacherRepo is a mock of TeacherRepo interface.
type MockTeacherRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTeacherRepoMockRecorder
}

// MockTeacherRepoMockRecorder is the mock recorder for MockTeacherRepo.
type MockTeacherRepoMockRecorder struct {
	mock *MockTeacherRepo
}

// NewMockTeacherRepo creates a new mock instance.
func NewMockTeacherRepo(ctrl *gomock.Controller) *MockTeacherRepo {
	mock := &MockTeacherRepo{ctrl: ctrl}
	mock.recorder = &MockTeacherRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeacherRepo) EXPECT() *MockTeacherRepoMockRecorder {
	return m.recorder
}

// AddAttendance mocks base method.
func (m *MockTeacherRepo) AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAttendance", userid, currentDate, newAttendanceID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAttendance indicates an expected call of AddAttendance.
func (mr *MockTeacherRepoMockRecorder) AddAttendance(userid, currentDate, newAttendanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAttendance", reflect.TypeOf((*MockTeacherRepo)(nil).AddAttendance), userid, currentDate, newAttendanceID)
}

// ClassMapAttendancePunch mocks base method.
func (m *MockTeacherRepo) ClassMapAttendancePunch(id int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassMapAttendancePunch", id)
	ret0, _ := ret[0].(string)
	return ret0
}

// ClassMapAttendancePunch indicates an expected call of ClassMapAttendancePunch.
func (mr *MockTeacherRepoMockRecorder) ClassMapAttendancePunch(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassMapAttendancePunch", reflect.TypeOf((*MockTeacherRepo)(nil).ClassMapAttendancePunch), id)
}

// CreatePunchIn mocks base method.
func (m *MockTeacherRepo) CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePunchIn", userID, currentDate, attendanceID, className)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePunchIn indicates an expected call of CreatePunchIn.
func (mr *MockTeacherRepoMockRecorder) CreatePunchIn(userID, currentDate, attendanceID, className interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePunchIn", reflect.TypeOf((*MockTeacherRepo)(nil).CreatePunchIn), userID, currentDate, attendanceID, className)
}

// EnrollCheckTeacher mocks base method.
func (m *MockTeacherRepo) EnrollCheckTeacher(userid, class string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnrollCheckTeacher", userid, class)
}

// EnrollCheckTeacher indicates an expected call of EnrollCheckTeacher.
func (mr *MockTeacherRepoMockRecorder) EnrollCheckTeacher(userid, class interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrollCheckTeacher", reflect.TypeOf((*MockTeacherRepo)(nil).EnrollCheckTeacher), userid, class)
}

// FetchAttendance mocks base method.
func (m *MockTeacherRepo) FetchAttendance(userid string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttendance", userid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAttendance indicates an expected call of FetchAttendance.
func (mr *MockTeacherRepoMockRecorder) FetchAttendance(userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttendance", reflect.TypeOf((*MockTeacherRepo)(nil).FetchAttendance), userid)
}

// FetchAttendanceWithMonth mocks base method.
func (m *MockTeacherRepo) FetchAttendanceWithMonth(id string, month, year int) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttendanceWithMonth", id, month, year)
	ret0, _ := ret[0].([]int)
	return ret0
}

// FetchAttendanceWithMonth indicates an expected call of FetchAttendanceWithMonth.
func (mr *MockTeacherRepoMockRecorder) FetchAttendanceWithMonth(id, month, year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttendanceWithMonth", reflect.TypeOf((*MockTeacherRepo)(nil).FetchAttendanceWithMonth), id, month, year)
}

// FetchClass mocks base method.
func (m *MockTeacherRepo) FetchClass(enrolledClass string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchClass", enrolledClass)
	ret0, _ := ret[0].(int)
	return ret0
}

// FetchClass indicates an expected call of FetchClass.
func (mr *MockTeacherRepoMockRecorder) FetchClass(enrolledClass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchClass", reflect.TypeOf((*MockTeacherRepo)(nil).FetchClass), enrolledClass)
}

// FetchDay mocks base method.
func (m *MockTeacherRepo) FetchDay(id int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDay", id)
	ret0, _ := ret[0].(int)
	return ret0
}

// FetchDay indicates an expected call of FetchDay.
func (mr *MockTeacherRepoMockRecorder) FetchDay(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDay", reflect.TypeOf((*MockTeacherRepo)(nil).FetchDay), id)
}

// FetchPunch mocks base method.
func (m *MockTeacherRepo) FetchPunch(id int) []attendanceServiceBean.PunchRecord {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPunch", id)
	ret0, _ := ret[0].([]attendanceServiceBean.PunchRecord)
	return ret0
}

// FetchPunch indicates an expected call of FetchPunch.
func (mr *MockTeacherRepoMockRecorder) FetchPunch(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPunch", reflect.TypeOf((*MockTeacherRepo)(nil).FetchPunch), id)
}

// FetchStudent mocks base method.
func (m *MockTeacherRepo) FetchStudent(userid string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchStudent", userid)
	ret0, _ := ret[0].(string)
	return ret0
}

// FetchStudent indicates an expected call of FetchStudent.
func (mr *MockTeacherRepoMockRecorder) FetchStudent(userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchStudent", reflect.TypeOf((*MockTeacherRepo)(nil).FetchStudent), userid)
}

// FetchStudentAttendance mocks base method.
func (m *MockTeacherRepo) FetchStudentAttendance(day, month, year int) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchStudentAttendance", day, month, year)
	ret0, _ := ret[0].([]int)
	return ret0
}

// FetchStudentAttendance indicates an expected call of FetchStudentAttendance.
func (mr *MockTeacherRepoMockRecorder) FetchStudentAttendance(day, month, year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchStudentAttendance", reflect.TypeOf((*MockTeacherRepo)(nil).FetchStudentAttendance), day, month, year)
}

// PunchCheckTeacher mocks base method.
func (m *MockTeacherRepo) PunchCheckTeacher(userid string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchCheckTeacher", userid)
	ret0, _ := ret[0].([]string)
	return ret0
}

// PunchCheckTeacher indicates an expected call of PunchCheckTeacher.
func (mr *MockTeacherRepoMockRecorder) PunchCheckTeacher(userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchCheckTeacher", reflect.TypeOf((*MockTeacherRepo)(nil).PunchCheckTeacher), userid)
}

// PunchOut mocks base method.
func (m *MockTeacherRepo) PunchOut(userID string, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchOut", userID, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// PunchOut indicates an expected call of PunchOut.
func (mr *MockTeacherRepoMockRecorder) PunchOut(userID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchOut", reflect.TypeOf((*MockTeacherRepo)(nil).PunchOut), userID, id)
}

// PunchOutCheck mocks base method.
func (m *MockTeacherRepo) PunchOutCheck(userid string, classid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchOutCheck", userid, classid)
	ret0, _ := ret[0].(error)
	return ret0
}

// PunchOutCheck indicates an expected call of PunchOutCheck.
func (mr *MockTeacherRepoMockRecorder) PunchOutCheck(userid, classid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchOutCheck", reflect.TypeOf((*MockTeacherRepo)(nil).PunchOutCheck), userid, classid)
}

// UpdatePunchOut mocks base method.
func (m *MockTeacherRepo) UpdatePunchOut(punchid, classid int, currentDate time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePunchOut", punchid, classid, currentDate)
}

// UpdatePunchOut indicates an expected call of UpdatePunchOut.
func (mr *MockTeacherRepoMockRecorder) UpdatePunchOut(punchid, classid, currentDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePunchOut", reflect.TypeOf((*MockTeacherRepo)(nil).UpdatePunchOut), punchid, classid, currentDate)
}
