package principalRepository

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices/userServiceBean"
	"go.uber.org/zap"
)

type PrincipalRepository struct {
	dbConnection       *pg.DB
	userServices       userServices.UserService
	attendanceServices attendanceServices.AttendanceService
	logger             *zap.SugaredLogger
}
type PrincipalRepo interface {
	InsertingStudent(id, username, password string)
	InsertingTeacher(id, username, password string)
	ClassMapping(classname string) int
	InsertClass(classname string) (int, error)
	InsertClassMap(id string, classId int)
	ClassMappingTeacher(classname []string) (map[string]int, error)
	FetchAttendance(id string, month, year int) []int
	FetchDay(id int) int
	FetchPunch(id int) []attendanceServiceBean.PunchRecord
	ClassMapAttendancePunch(id int) string
	FetchUser(role string) []userServiceBean.CustomUserInfo
	CheckAlreadyInsert(id string, role string) error
	FetchClasses(username string) []string
	AllClasses() []string
}

func NewPrincipalRepositoryImpl(dbConnection *pg.DB, userServices userServices.UserService, attendanceServices attendanceServices.AttendanceService, logger *zap.SugaredLogger) *PrincipalRepository {
	return &PrincipalRepository{
		dbConnection:       dbConnection,
		userServices:       userServices,
		attendanceServices: attendanceServices,
		logger:             logger,
	}
}
func (impl PrincipalRepository) FetchClasses(username string) []string {
	return impl.userServices.FetchClasses(username)
}
func (impl PrincipalRepository) AllClasses() []string {
	return impl.userServices.AllClasses()
}
func (impl PrincipalRepository) InsertingStudent(id, username, password string) {
	// Insert student data into the postgres database
	err := impl.userServices.InsertingStudent(id, username, password)
	if err != nil {
		impl.logger.Errorw("Error inserting student data into the database", "error", err)
		return
	}
}
func (impl PrincipalRepository) CheckAlreadyInsert(id string, role string) error {
	//// Insert student data into the postgres database
	//err := impl.userServices.InsertingStudent(id, username, password)
	//if err != nil {
	//	impl.logger.Errorw("Error inserting student data into the database", "error", err)
	//	return
	//}
	var userRole userModels.UserRoleType
	if role == "student" {
		userRole = userModels.Student
	} else {
		userRole = userModels.Teacher
	}
	var userID string
	err := impl.dbConnection.Model(&userModels.User{}).
		Column("id").
		Where("role = ? AND id=?", userRole, id).
		Select(&userID)
	if err == nil {
		return errors.New("Already inserted")
	}
	return nil
}
func (impl PrincipalRepository) InsertingTeacher(id, username, password string) {
	// Insert student data into the postgres database
	err := impl.userServices.InsertingTeacher(id, username, password)
	if err != nil {
		impl.logger.Errorw("Error inserting student data into the database", "error", err)
		return
	}
}
func (impl PrincipalRepository) ClassMapping(classname string) int {
	// Map the class name to its ID
	return impl.userServices.FetchClass(classname)
}
func (impl PrincipalRepository) InsertClass(classname string) (int, error) {
	return impl.userServices.InsertClass(classname)
}
func (impl PrincipalRepository) InsertClassMap(id string, classId int) {
	err := impl.userServices.InsertClassMap(id, classId)
	if err != nil {
		impl.logger.Errorw("Error inserting class mapping data into the database", "error", err)
		return
	}
	var classMappingUser []userModels.ClassMappingUser
	err = impl.dbConnection.Model(&classMappingUser).Select()
	if err != nil {
		panic(err)
	}

	for _, user := range classMappingUser {
		fmt.Printf("ID: %d, UserID: %s, ClassID: %d\n", user.ID, user.UserID, user.ClassID)
	}
	//fmt.Println("in check enrollemnt", class, userid)
	//err = impl.dbConnection.Model(&classMappingUser).
	//	Where("user_id = ? AND class_id IN (SELECT class_id FROM classes WHERE class_name = ?)", id, class).
	//	Select()
	//return err
}
func (impl PrincipalRepository) ClassMappingTeacher(classname []string) (map[string]int, error) {
	// Map existing class names to their IDs
	return impl.userServices.ClassMappingTeacher(classname)

}
func (impl PrincipalRepository) FetchAttendance(id string, month, year int) []int {
	var ids []int // Array to store the fetched IDs
	ids = impl.attendanceServices.FetchAttendanceofUser(id, month, year)
	return ids
}
func (impl PrincipalRepository) FetchDay(id int) int {
	return impl.attendanceServices.FetchDay(id)
}
func (impl PrincipalRepository) FetchPunch(id int) []attendanceServiceBean.PunchRecord {
	// Fetch punch-in/out records for each attendance
	return impl.attendanceServices.FetchPunch(id)

}
func (impl PrincipalRepository) ClassMapAttendancePunch(id int) string {
	// Fetch class information and class name
	return impl.attendanceServices.FetchClassMapPunch(id)
}
func (impl PrincipalRepository) FetchUser(role string) []userServiceBean.CustomUserInfo {

	return impl.userServices.FetchUser(role)
}
