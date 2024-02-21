package studentRepository

import (
	"github.com/go-pg/pg/v10"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices"
	"go.uber.org/zap"
	"time"
)

type StudentRepository struct {
	dbConnection       *pg.DB
	userServices       userServices.UserService
	attendanceServices attendanceServices.AttendanceService
	logger             *zap.SugaredLogger
}
type StudentRepo interface {
	EnrollCheck(userid, class string)
	PunchCheck(userid string) string
	FetchAttendance(userid string) (int, error)
	FetchAttendanceWithMonth(id string, month, year int) []int
	FetchDay(id int) int
	FetchPunch(id int) []attendanceServiceBean.PunchRecord
	ClassMapAttendancePunch(id int) string
	FetchClass(enrolledClass string) int
	PunchOutCheck(userid string, classid int) error
	CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error
	AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error)
	UpdatePunchOut(punchid int, classid int, currentDate time.Time) error
	FetchClassUser(userid string) int
}

func NewStudentRepositoryImpl(dbConnection *pg.DB, userServices userServices.UserService, attendanceServices attendanceServices.AttendanceService, logger *zap.SugaredLogger) *StudentRepository {
	return &StudentRepository{
		dbConnection:       dbConnection,
		userServices:       userServices,
		attendanceServices: attendanceServices,
		logger:             logger,
	}
}
func (impl *StudentRepository) FetchClassUser(userid string) int {
	return impl.userServices.FetchClassUser(userid)
}
func (impl *StudentRepository) EnrollCheck(userid, class string) {
	// Check if the user is enrolled in the class
	//var classMappingUser studentModels.ClassMappingUser
	//err := impl.dbConnection.Model(&classMappingUser).
	//	Where("user_id = ? AND class_id IN (SELECT class_id FROM classes WHERE class_name = ?)", userid, class).
	//	Select()
	err := impl.userServices.CheckEnrollment(userid, class)

	if err != nil {
		impl.logger.Errorw("You haven't enrolled in this class", "error", err)
		return
	}

}
func (impl *StudentRepository) PunchCheck(userid string) string {
	// Check if the user has already punched out from any class on the same day
	var enrolledClass []string
	enrolledClass, err := impl.userServices.CheckPunchOut(userid)
	if err != nil {
		impl.logger.Errorw("Error fetching enrolled classes", "error", err)
		return ""
	}
	return enrolledClass[0]
}
func (impl *StudentRepository) FetchAttendance(userid string) (int, error) {
	// Check if the user has an entry in the attendance table on the same day for the given class
	//var existingAttendance studentModels.Attendance
	//currentDate := time.Now()
	////fmt.Println("user_id day month year", userid)
	////fmt.Printf("%T %T %T", currentDate.Day(), int(currentDate.Month()), currentDate.Year())
	//err := impl.dbConnection.Model(&existingAttendance).
	//	Column("id").
	//	Where("user_id = ? AND day = ? AND month = ? AND year = ?", userid, currentDate.Day(), int(currentDate.Month()), currentDate.Year()).
	//	Select()
	//fmt.Println(1)
	//fmt.Println(existingAttendance)
	//return existingAttendance.ID, err
	id, err := impl.attendanceServices.HasAttendance(userid)
	return id, err

}
func (impl *StudentRepository) FetchAttendanceWithMonth(id string, month, year int) []int {
	var ids []int // Array to store the fetched IDs
	//var studentAttendances []studentModels.Attendance
	//err := impl.dbConnection.Model(&studentAttendances).
	//	Where("user_id = ? AND month = ? AND year = ?", id, month, year).
	//	Select() // Fetch records that match the conditions
	//if err != nil {
	//	impl.logger.Errorw("Error querying teacher attendance from the database", "error", err)
	//}
	ids = impl.attendanceServices.FetchAttendanceofUser(id, month, year)
	return ids
	//// Extract 'ID' values into the 'ids' array
	//for _, attendance := range studentAttendances {
	//	ids = append(ids, attendance.ID)
	//}
	//if err != nil {
	//	impl.logger.Errorw("Error querying teacher attendance from the database", "error", err)
	//}
	//return ids
}
func (impl *StudentRepository) FetchDay(id int) int {
	//var attendance studentModels.Attendance
	//err := impl.dbConnection.Model(&attendance).
	//	Where("id = ?", id).
	//	Select()
	//if err != nil {
	//	impl.logger.Errorw("Cant fetch day from attendance id", err)
	//}
	//return attendance.Day
	return impl.attendanceServices.FetchDay(id)
}
func (impl *StudentRepository) FetchPunch(id int) []attendanceServiceBean.PunchRecord {
	// Fetch punch-in/out records for each attendance
	//var punchInOutRecords []studentModels.PunchInOut
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
	return impl.attendanceServices.FetchPunch(id)
	//// Create a slice to store the result
	//var punchRecords []studentServiceBean.PunchRecord
	//// Populate punchRecords with data from punchInOutRecords
	//for _, record := range punchInOutRecords {
	//	punchRecord := studentServiceBean.PunchRecord{
	//		ID:       record.ID,
	//		PunchIn:  record.PunchIn,
	//		PunchOut: record.PunchOut,
	//	}
	//	punchRecords = append(punchRecords, punchRecord)
	//}
	//
	//return punchRecords
}
func (impl *StudentRepository) ClassMapAttendancePunch(id int) string {
	// Fetch class information and class name
	//var className string
	//err := impl.dbConnection.Model(&studentModels.ClassMappingAttendance{}).
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
func (impl *StudentRepository) FetchClass(enrolledClass string) int {
	// User has an entry in the attendance table on the same day for the given class
	// Fetch the ClassID based on the provided class name
	//var classInfo studentModels.Class
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
func (impl *StudentRepository) PunchOutCheck(userid string, classid int) error {
	// Check if the user has punched out
	//var existingPunch studentModels.PunchInOut
	//err := impl.dbConnection.Model(&existingPunch).
	//	Join("JOIN class_mapping_attendances ON class_mapping_attendances.punch_id = punch_in_out.id").
	//	Where("punch_in_out.user_id = ? AND class_mapping_attendances.class_id = ? AND punch_in_out.punch_out IS NULL ", userid, classid).
	//	Select()
	//return err
	return impl.attendanceServices.PunchOutCheck(userid, classid)

}

func (impl *StudentRepository) CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error {
	//punchInTime := currentDate.Format("15:04:05")
	//
	//// Create a new punch-in record
	//_, err := impl.dbConnection.Model(&studentModels.PunchInOut{
	//	AttendanceID: attendanceID,
	//	UserID:       userID,
	//	PunchIn:      punchInTime,
	//}).Insert()
	//if err != nil {
	//	return err
	//}
	//// Fetch the PunchID for the newly created punch-in record
	//var punch studentModels.PunchInOut
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
	//var classInfo studentModels.Class
	//err = impl.dbConnection.Model(&classInfo).
	//	Column("class_id").
	//	Where("class_name = ?", className).
	//	Select()
	//
	//if err != nil {
	//	return err
	//}
	//// Create a new ClassMappingAttendance record
	//_, err = impl.dbConnection.Model(&studentModels.ClassMappingAttendance{
	//	PunchID: punch.ID,
	//	ClassID: classInfo.ClassID,
	//}).Insert()
	//return err
	return impl.attendanceServices.CreatePunchIn(userID, currentDate, attendanceID, className)
}
func (impl *StudentRepository) AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error) {
	//var atID int
	//_, err := impl.dbConnection.Model(&studentModels.Attendance{
	//	UserID: userid,
	//	Day:    currentDate.Day(),
	//	Month:  int(currentDate.Month()),
	//	Year:   currentDate.Year(),
	//}).Returning("id").Insert(&atID)
	////_, err := impl.dbConnection.Model(&studentAttendance).Returning("id").Insert(&newAttendanceID)
	//return atID, err
	return impl.attendanceServices.AddAttendance(userid, currentDate, newAttendanceID)
}
func (impl *StudentRepository) UpdatePunchOut(attenid int, classid int, currentDate time.Time) error {
	// Update the existing punch record with punch-out time
	//_, err := impl.dbConnection.Model(&studentModels.PunchInOut{}).
	//	Set("punch_out = ?", currentDate.Format("15:04:05")).
	//	Where("punch_in_out.attendance_id= ? AND punch_in_out.punch_out IS NULL", attenid, classid).
	//	Update()
	//
	//if err != nil {
	//	impl.logger.Errorw("Error updating punch-out time", "error", err)
	//	return
	//}
	return impl.attendanceServices.UpdatePunchOut(attenid, classid, currentDate)
}
