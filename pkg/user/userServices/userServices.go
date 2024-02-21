package userServices

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userAdapter"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices/userServiceBean"
	"go.uber.org/zap"
)

type UserServiceImpl struct {
	userRepository userRepository.UserRepo
	logger         *zap.SugaredLogger
}
type UserService interface {
	CheckEnrollment(userid, class string) error
	CheckPunchOut(userid string) ([]string, error)
	FetchClass(enrolledClass string) int
	FetchStudent(userid string) string
	InsertingStudent(id, username, password string) error
	InsertingTeacher(id, username, password string) error
	InsertClass(classname string) (int, error)
	InsertClassMap(id string, classId int) error
	ClassMappingTeacher(classname []string) (map[string]int, error)
	FetchUser(role string) []userServiceBean.CustomUserInfo
	FetchClassUser(userid string) int
	FetchClasses(username string) []string
	AllClasses() []string
}

func NewUserServiceImpl(userRepository userRepository.UserRepo, logger *zap.SugaredLogger) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository,
		logger,
	}
}
func (impl UserServiceImpl) AllClasses() []string {
	return impl.userRepository.AllClasses()
}
func (impl UserServiceImpl) FetchClasses(username string) []string {
	return impl.userRepository.FetchClasses(username)
}
func (impl UserServiceImpl) CheckEnrollment(userid, class string) error {
	return impl.userRepository.CheckEnrollment(userid, class)
}
func (impl UserServiceImpl) CheckPunchOut(userid string) ([]string, error) {
	return impl.userRepository.CheckPunchOut(userid)
}
func (impl UserServiceImpl) FetchClass(enrolledClass string) int {
	return impl.userRepository.FetchClass(enrolledClass)
}
func (impl UserServiceImpl) FetchClassUser(userid string) int {
	return impl.userRepository.FetchClassUser(userid)
}
func (impl UserServiceImpl) FetchStudent(userid string) string {
	return impl.userRepository.FetchStudent(userid)
}
func (impl UserServiceImpl) InsertingStudent(id, username, password string) error {
	return impl.userRepository.InsertingStudent(id, username, password)
}
func (impl UserServiceImpl) InsertingTeacher(id, username, password string) error {
	return impl.userRepository.InsertingTeacher(id, username, password)
}
func (impl UserServiceImpl) InsertClass(classname string) (int, error) {
	return impl.userRepository.InsertClass(classname)
}
func (impl UserServiceImpl) InsertClassMap(id string, classId int) error {
	return impl.userRepository.InsertClassMap(id, classId)
}
func (impl UserServiceImpl) ClassMappingTeacher(classname []string) (map[string]int, error) {
	return impl.userRepository.ClassMappingTeacher(classname)
}
func (impl UserServiceImpl) FetchUser(role string) []userServiceBean.CustomUserInfo {
	return userAdapter.Userrec(impl.userRepository.FetchUser(role))
}
