package teacherRepository

import (
	"github.com/go-pg/pg/v10"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices"
	"go.uber.org/zap"
	"time"
)

type TeacherRepository struct {
	dbConnection       *pg.DB
	userServices       userServices.UserService
	attendanceServices attendanceServices.AttendanceService
	logger             *zap.SugaredLogger
}
type TeacherRepo interface {
	EnrollCheckTeacher(userid, class string)
	PunchCheckTeacher(userid string) []string
	FetchAttendance(userid string) (int, error)
	FetchClass(enrolledClass string) int
	PunchOutCheck(userid string, classid int) error
	PunchOut(userID string, id int) error
	AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error)
	CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error
	UpdatePunchOut(punchid int, classid int, currentDate time.Time)
	FetchAttendanceWithMonth(id string, month, year int) []int
	FetchDay(id int) int
	FetchPunch(id int) []attendanceServiceBean.PunchRecord
	ClassMapAttendancePunch(id int) string
	FetchStudentAttendance(day, month, year int) []int
	FetchStudent(userid string) string
}

func NewTeacherRepositoryImpl(dbConnection *pg.DB, userServices userServices.UserService, attendanceServices attendanceServices.AttendanceService, logger *zap.SugaredLogger) *TeacherRepository {
	return &TeacherRepository{
		dbConnection:       dbConnection,
		userServices:       userServices, //FetchUser(role string) []common2.CustomUserInfo
		attendanceServices: attendanceServices,
		logger:             logger,
	}
}
func (impl *TeacherRepository) EnrollCheckTeacher(userid, class string) {
	// Check if the user is enrolled in the class
	//var classMappingUser teacherModels.ClassMappingUser
	//err := impl.dbConnection.Model(&classMappingUser).
	//	Where("user_id = ? AND class_id IN (SELECT class_id FROM classes WHERE class_name = ?)", userid, class).
	//	Select()
	err := impl.userServices.CheckEnrollment(userid, class)
	if err != nil {
		impl.logger.Errorw("You haven't enrolled in this class", "error", err)
		return
	}

}
func (impl *TeacherRepository) PunchCheckTeacher(userid string) []string {
	// Check if the user has already punched out from any class on the same day
	//var enrolledClass []string
	//err := impl.dbConnection.Model(&teacherModels.ClassMappingUser{}).
	//	ColumnExpr("DISTINCT class_name").
	//	Join("JOIN classes ON class_mapping_user.class_id = classes.class_id").
	//	Where("user_id = ? ", userid).
	//	Select(&enrolledClass)
	enrolledClass, err := impl.userServices.CheckPunchOut(userid)
	if err != nil {
		impl.logger.Errorw("Error fetching enrolled classes", "error", err)
	}
	return enrolledClass
}
func (impl *TeacherRepository) FetchAttendance(userid string) (int, error) {
	// Check if the user has an entry in the attendance table on the same day for the given class
	//var existingAttendance teacherModels.Attendance
	//currentDate := time.Now()
	//err := impl.dbConnection.Model(&existingAttendance).
	//	Column("id").
	//	Where("user_id = ? AND day = ? AND month = ? AND year = ?", userid, currentDate.Day(), int(currentDate.Month()), currentDate.Year()).
	//	Select()
	id, err := impl.attendanceServices.HasAttendance(userid)
	return id, err
}
func (impl *TeacherRepository) FetchClass(enrolledClass string) int {
	// User has an entry in the attendance table on the same day for the given class
	// Fetch the ClassID based on the provided class name
	//var classInfo teacherModels.Class
	//err := impl.dbConnection.Model(&classInfo).
	//	Column("class_id").
	//	Where("class_name = ?", enrolledClass).
	//	Select()
	//
	//if err != nil {
	//	impl.logger.Errorw("Error fetching class ID", "error", err)
	//	return 0
	//}
	//return classInfo.ClassID
	return impl.userServices.FetchClass(enrolledClass)
}
func (impl *TeacherRepository) PunchOutCheck(userid string, classid int) error {
	// Check if the user has punched out
	//var existingPunch teacherModels.PunchInOut
	//err := impl.dbConnection.Model(&existingPunch).
	//	Join("JOIN class_mapping_attendances ON class_mapping_attendances.punch_id = punch_in_out.id").
	//	Where("punch_in_out.user_id = ? AND class_mapping_attendances.class_id = ? AND punch_in_out.punch_out IS NULL ", userid, classid).
	//	Select()
	//return err
	return impl.attendanceServices.PunchOutCheck(userid, classid)

}
func (impl *TeacherRepository) PunchOut(userID string, id int) error {
	// User has an entry in the attendance table on the same day for the requested class
	// Check if the user has punched out
	//var existingPunch teacherModels.PunchInOut
	//err := impl.dbConnection.Model(&existingPunch).
	//	Column("attendance_id").
	//	Where("user_id = ? AND punch_id = ?", userID, id).
	//	Select()
	//return err
	return impl.attendanceServices.PunchOut(userID, id)
}
func (impl *TeacherRepository) AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error) {
	//_, err := impl.dbConnection.Model(&model.Attendance{
	//	UserID: userid,
	//	Day:    currentDate.Day(),
	//	Month:  int(currentDate.Month()),
	//	Year:   currentDate.Year(),
	//}).Returning("id").Insert(&newAttendanceID)
	////return err
	//var atID int
	//_, err := impl.dbConnection.Model(&teacherModels.Attendance{
	//	UserID: userid,
	//	Day:    currentDate.Day(),
	//	Month:  int(currentDate.Month()),
	//	Year:   currentDate.Year(),
	//}).Returning("id").Insert(&atID)
	////_, err := impl.dbConnection.Model(&studentAttendance).Returning("id").Insert(&newAttendanceID)
	//return atID, err
	return impl.attendanceServices.AddAttendance(userid, currentDate, newAttendanceID)
}
func (impl *TeacherRepository) CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error {
	//punchInTime := currentDate.Format("15:04:05")
	//
	//// Create a new punch-in record
	//_, err := impl.dbConnection.Model(&teacherModels.PunchInOut{
	//	AttendanceID: attendanceID,
	//	UserID:       userID,
	//	PunchIn:      punchInTime,
	//}).Insert()
	//if err != nil {
	//	return err
	//}
	//// Fetch the PunchID for the newly created punch-in record
	//var punch teacherModels.PunchInOut
	//err = impl.dbConnection.Model(&punch).
	//	Column("id").
	//	Where("attendance_id = ? AND punch_in = ?", attendanceID, punchInTime).
	//	Select()
	//fmt.Println(err)
	//if err != nil {
	//	return err
	//}
	//
	//// Fetch the ClassID based on the provided class name
	//var classInfo teacherModels.Class
	//err = impl.dbConnection.Model(&classInfo).
	//	Column("class_id").
	//	Where("class_name = ?", className).
	//	Select()
	//
	//if err != nil {
	//	return err
	//}
	//// Create a new ClassMappingAttendance record
	//_, err = impl.dbConnection.Model(&teacherModels.ClassMappingAttendance{
	//	PunchID: punch.ID,
	//	ClassID: classInfo.ClassID,
	//}).Insert()
	//return err
	return impl.attendanceServices.CreatePunchIn(userID, currentDate, attendanceID, className)
}
func (impl *TeacherRepository) UpdatePunchOut(attedid int, classid int, currentDate time.Time) {
	// Update the existing punch record with punch-out time
	//_, err := impl.dbConnection.Model(&teacherModels.PunchInOut{}).
	//	Set("punch_out = ?", currentDate.Format("15:04:05")).
	//	Where("punch_in_out.attendance_id= ? AND punch_in_out.punch_out IS NULL", attedid, classid).
	//	Update()
	err := impl.attendanceServices.UpdatePunchOut(attedid, classid, currentDate)
	if err != nil {
		impl.logger.Errorw("Error updating punch-out time", "error", err)
		return
	}
}
func (impl *TeacherRepository) FetchAttendanceWithMonth(id string, month, year int) []int {
	//var ids []int // Array to store the fetched IDs
	//var studentAttendances []teacherModels.Attendance
	//err := impl.dbConnection.Model(&studentAttendances).
	//	Where("user_id = ? AND month = ? AND year = ?", id, month, year).
	//	Select() // Fetch records that match the conditions
	//if err != nil {
	//	impl.logger.Errorw("Error querying teacher attendance from the databazse", "error", err)
	//}
	//return studentAttendances
	//// Extract 'ID' values into the 'ids' array
	//for _, attendance := range studentAttendances {
	//	ids = append(ids, attendance.ID)
	//}
	//if err != nil {
	//	impl.logger.Errorw("Error querying teacher attendance from the database", "error", err)
	//}
	//return ids
	return impl.attendanceServices.FetchAttendanceofUser(id, month, year)
}
func (impl *TeacherRepository) FetchDay(id int) int {
	//var attendance teacherModels.Attendance
	//err := impl.dbConnection.Model(&attendance).
	//	Where("id = ?", id).
	//	Select()
	//if err != nil {
	//	impl.logger.Errorw("Cant fetch day from attendance id", err)
	//}
	//return attendance.Day
	return impl.attendanceServices.FetchDay(id)
}
func (impl *TeacherRepository) FetchPunch(id int) []attendanceServiceBean.PunchRecord {
	// Fetch punch-in/out records for each attendance
	//var punchInOutRecords []teacherModels.PunchInOut
	//err := impl.dbConnection.Model(&punchInOutRecords).
	//	Column("id", "attendance_id", "user_id", "punch_in", "punch_out").
	//	Where("attendance_id = ?", id).
	//	Order("punch_in", "punch_out").
	//	Select()
	//
	//if err != nil {
	//	impl.logger.Errorw("Error querying punch-in/out records from the database", "error", err)
	//	return nil
	//}
	//return punchInOutRecords
	//// Create a slice to store the result
	//var punchRecords []adapter.PunchRecord
	//// Populate punchRecords with data from punchInOutRecords
	//for _, record := range punchInOutRecords {
	//	punchRecord := adapter.PunchRecord{
	//		ID:       record.ID,
	//		UserID:   record.UserID,
	//		PunchIn:  record.PunchIn,
	//		PunchOut: record.PunchOut,
	//	}
	//	punchRecords = append(punchRecords, punchRecord)
	//}
	//
	//return punchRecords
	return impl.attendanceServices.FetchPunch(id)
}
func (impl *TeacherRepository) ClassMapAttendancePunch(id int) string {
	// Fetch class information and class name
	//var className string
	//err := impl.dbConnection.Model(&teacherModels.ClassMappingAttendance{}).
	//	Column("classes.class_name").
	//	Where("punch_id = ?", id).
	//	Join("JOIN classes ON class_mapping_attendance.class_id = classes.class_id").
	//	Select(&className)
	//
	//if err != nil {
	//	impl.logger.Errorw("Error querying class name from the database", "error", err)
	//}
	//return className
	return impl.attendanceServices.FetchClassMapPunch(id)
}
func (impl *TeacherRepository) FetchStudentAttendance(day, month, year int) []int {
	// Fetch attendance records for the student for the given month and year
	//var studentAttendances []teacherModels.Attendance
	//err := impl.dbConnection.Model(&studentAttendances).
	//	Where("day = ? AND month = ? AND year = ?", day, month, year).
	//	Select()
	//
	//if err != nil {
	//	impl.logger.Errorw("Error querying student attendance from the database", "error", err)
	//
	//}
	//var ids []int
	//// Extract 'ID' values into the 'ids' array
	//for _, attendance := range studentAttendances {
	//	ids = append(ids, attendance.ID)
	//}
	//return ids, err
	return impl.attendanceServices.FetchAttendance(day, month, year)
}
func (impl *TeacherRepository) FetchStudent(userid string) string {
	//var userID string
	//err := impl.dbConnection.Model(&teacherModels.User{}).
	//	Column("id").
	//	Where("role = ? AND id=?", "student", userid).
	//	Select(&userID)
	//if err != nil {
	//	impl.logger.Errorw("Error fetching student", err)
	//}
	//return userID
	return impl.userServices.FetchStudent(userid)

}
