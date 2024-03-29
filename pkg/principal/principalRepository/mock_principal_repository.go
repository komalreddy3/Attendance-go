// Code generated by MockGen. DO NOT EDIT.
// Source: principalRepository/principalRepository.go

// Package principalRepository is a generated GoMock package.
package principalRepository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	attendanceServiceBean "github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	userServiceBean "github.com/komalreddy3/Attendance-go/pkg/user/userServices/userServiceBean"
)

// MockPrincipalRepo is a mock of PrincipalRepo interface.
type MockPrincipalRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPrincipalRepoMockRecorder
}

// MockPrincipalRepoMockRecorder is the mock recorder for MockPrincipalRepo.
type MockPrincipalRepoMockRecorder struct {
	mock *MockPrincipalRepo
}

// NewMockPrincipalRepo creates a new mock instance.
func NewMockPrincipalRepo(ctrl *gomock.Controller) *MockPrincipalRepo {
	mock := &MockPrincipalRepo{ctrl: ctrl}
	mock.recorder = &MockPrincipalRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrincipalRepo) EXPECT() *MockPrincipalRepoMockRecorder {
	return m.recorder
}

// ClassMapAttendancePunch mocks base method.
func (m *MockPrincipalRepo) ClassMapAttendancePunch(id int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassMapAttendancePunch", id)
	ret0, _ := ret[0].(string)
	return ret0
}

// ClassMapAttendancePunch indicates an expected call of ClassMapAttendancePunch.
func (mr *MockPrincipalRepoMockRecorder) ClassMapAttendancePunch(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassMapAttendancePunch", reflect.TypeOf((*MockPrincipalRepo)(nil).ClassMapAttendancePunch), id)
}

// ClassMapping mocks base method.
func (m *MockPrincipalRepo) ClassMapping(classname string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassMapping", classname)
	ret0, _ := ret[0].(int)
	return ret0
}

// ClassMapping indicates an expected call of ClassMapping.
func (mr *MockPrincipalRepoMockRecorder) ClassMapping(classname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassMapping", reflect.TypeOf((*MockPrincipalRepo)(nil).ClassMapping), classname)
}

// ClassMappingTeacher mocks base method.
func (m *MockPrincipalRepo) ClassMappingTeacher(classname []string) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassMappingTeacher", classname)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClassMappingTeacher indicates an expected call of ClassMappingTeacher.
func (mr *MockPrincipalRepoMockRecorder) ClassMappingTeacher(classname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassMappingTeacher", reflect.TypeOf((*MockPrincipalRepo)(nil).ClassMappingTeacher), classname)
}

// FetchAttendance mocks base method.
func (m *MockPrincipalRepo) FetchAttendance(id string, month, year int) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAttendance", id, month, year)
	ret0, _ := ret[0].([]int)
	return ret0
}

// FetchAttendance indicates an expected call of FetchAttendance.
func (mr *MockPrincipalRepoMockRecorder) FetchAttendance(id, month, year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttendance", reflect.TypeOf((*MockPrincipalRepo)(nil).FetchAttendance), id, month, year)
}

// FetchDay mocks base method.
func (m *MockPrincipalRepo) FetchDay(id int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDay", id)
	ret0, _ := ret[0].(int)
	return ret0
}

// FetchDay indicates an expected call of FetchDay.
func (mr *MockPrincipalRepoMockRecorder) FetchDay(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDay", reflect.TypeOf((*MockPrincipalRepo)(nil).FetchDay), id)
}

// FetchPunch mocks base method.
func (m *MockPrincipalRepo) FetchPunch(id int) []attendanceServiceBean.PunchRecord {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPunch", id)
	ret0, _ := ret[0].([]attendanceServiceBean.PunchRecord)
	return ret0
}

// FetchPunch indicates an expected call of FetchPunch.
func (mr *MockPrincipalRepoMockRecorder) FetchPunch(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPunch", reflect.TypeOf((*MockPrincipalRepo)(nil).FetchPunch), id)
}

// FetchUser mocks base method.
func (m *MockPrincipalRepo) FetchUser(role string) []userServiceBean.CustomUserInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUser", role)
	ret0, _ := ret[0].([]userServiceBean.CustomUserInfo)
	return ret0
}

// FetchUser indicates an expected call of FetchUser.
func (mr *MockPrincipalRepoMockRecorder) FetchUser(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUser", reflect.TypeOf((*MockPrincipalRepo)(nil).FetchUser), role)
}

// InsertClass mocks base method.
func (m *MockPrincipalRepo) InsertClass(classname string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertClass", classname)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertClass indicates an expected call of InsertClass.
func (mr *MockPrincipalRepoMockRecorder) InsertClass(classname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertClass", reflect.TypeOf((*MockPrincipalRepo)(nil).InsertClass), classname)
}

// InsertClassMap mocks base method.
func (m *MockPrincipalRepo) InsertClassMap(id string, classId int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertClassMap", id, classId)
}

// InsertClassMap indicates an expected call of InsertClassMap.
func (mr *MockPrincipalRepoMockRecorder) InsertClassMap(id, classId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertClassMap", reflect.TypeOf((*MockPrincipalRepo)(nil).InsertClassMap), id, classId)
}

// InsertingStudent mocks base method.
func (m *MockPrincipalRepo) InsertingStudent(id, username, password string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertingStudent", id, username, password)
}

// InsertingStudent indicates an expected call of InsertingStudent.
func (mr *MockPrincipalRepoMockRecorder) InsertingStudent(id, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertingStudent", reflect.TypeOf((*MockPrincipalRepo)(nil).InsertingStudent), id, username, password)
}

// InsertingTeacher mocks base method.
func (m *MockPrincipalRepo) InsertingTeacher(id, username, password string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertingTeacher", id, username, password)
}

// InsertingTeacher indicates an expected call of InsertingTeacher.
func (mr *MockPrincipalRepoMockRecorder) InsertingTeacher(id, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertingTeacher", reflect.TypeOf((*MockPrincipalRepo)(nil).InsertingTeacher), id, username, password)
}
