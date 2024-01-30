package attendanceRepository

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"go.uber.org/zap"
	"time"
)

type AttendanceRepository struct {
	dbConnection *pg.DB
	logger       *zap.SugaredLogger
}
type AttendanceRepo interface {
	HasAttendance(userid string) (int, error)
	FetchAttendance(day, month, year int) []attendanceModels.Attendance
	FetchClassMapPunch(id int) string
	FetchPunch(id int) []attendanceModels.PunchInOut
	FetchDay(id int) int
	FetchAttendanceofUser(id string, month int, year int) []attendanceModels.Attendance
	UpdatePunchOut(attedid int, classid int, currentDate time.Time) error
	CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error
	AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error)
	PunchOut(userID string, id int) error
	PunchOutCheck(userid string, classid int) error
}

func NewAttendanceRepositoryImpl(dbConnection *pg.DB, logger *zap.SugaredLogger) *AttendanceRepository {
	return &AttendanceRepository{
		dbConnection: dbConnection,
		logger:       logger,
	}
}
func (impl AttendanceRepository) HasAttendance(userid string) (int, error) {
	var existingAttendance attendanceModels.Attendance
	currentDate := time.Now()
	err := impl.dbConnection.Model(&existingAttendance).
		Column("id").
		Where("user_id = ? AND day = ? AND month = ? AND year = ?", userid, currentDate.Day(), int(currentDate.Month()), currentDate.Year()).
		Select()
	return existingAttendance.ID, err
}
func (impl AttendanceRepository) FetchAttendance(day, month, year int) []attendanceModels.Attendance {
	var studentAttendances []attendanceModels.Attendance
	err := impl.dbConnection.Model(&studentAttendances).
		Where("day = ? AND month = ? AND year = ?", day, month, year).
		Select()

	if err != nil {
		impl.logger.Errorw("Error querying student attendance from the database", "error", err)

	}
	return studentAttendances
}
func (impl AttendanceRepository) FetchAttendanceofUser(id string, month int, year int) []attendanceModels.Attendance {
	var studentAttendances []attendanceModels.Attendance
	err := impl.dbConnection.Model(&studentAttendances).
		Where("user_id = ? AND month = ? AND year = ?", id, month, year).
		Select() // Fetch records that match the conditions
	if err != nil {
		impl.logger.Errorw("Error querying teacher attendance from the database", "error", err)
	}
	return studentAttendances
}
func (impl AttendanceRepository) FetchClassMapPunch(id int) string {
	var className string
	err := impl.dbConnection.Model(&attendanceModels.ClassMappingAttendance{}).
		Column("classes.class_name").
		Where("punch_id = ?", id).
		Join("JOIN classes ON class_mapping_attendance.class_id = classes.class_id").
		Select(&className)

	if err != nil {
		impl.logger.Errorw("Error querying class name from the database", "error", err)
	}
	return className
}
func (impl AttendanceRepository) FetchPunch(id int) []attendanceModels.PunchInOut {
	var punchInOutRecords []attendanceModels.PunchInOut
	err := impl.dbConnection.Model(&punchInOutRecords).
		Column("id", "attendance_id", "user_id", "punch_in", "punch_out").
		Where("attendance_id = ?", id).
		Order("punch_in", "punch_out").
		Select()

	if err != nil {
		impl.logger.Errorw("Error querying punch-in/out records from the database", "error", err)
		return nil
	}
	return punchInOutRecords
}
func (impl AttendanceRepository) FetchDay(id int) int {
	var attendance attendanceModels.Attendance
	err := impl.dbConnection.Model(&attendance).
		Where("id = ?", id).
		Select()
	if err != nil {
		impl.logger.Errorw("Cant fetch day from attendance id", err)
	}
	return attendance.Day
}
func (impl AttendanceRepository) UpdatePunchOut(attedid int, classid int, currentDate time.Time) error {
	_, err := impl.dbConnection.Model(&attendanceModels.PunchInOut{}).
		Set("punch_out = ?", currentDate.Format("15:04:05")).
		Where("punch_in_out.attendance_id= ? AND punch_in_out.punch_out IS NULL", attedid, classid).
		Update()
	return err
}
func (impl AttendanceRepository) CreatePunchIn(userID string, currentDate time.Time, attendanceID int, className string) error {
	punchInTime := currentDate.Format("15:04:05")

	// Create a new punch-in record
	_, err := impl.dbConnection.Model(&attendanceModels.PunchInOut{
		AttendanceID: attendanceID,
		UserID:       userID,
		PunchIn:      punchInTime,
	}).Insert()
	if err != nil {
		return err
	}
	// Fetch the PunchID for the newly created punch-in record
	var punch attendanceModels.PunchInOut
	err = impl.dbConnection.Model(&punch).
		Column("id").
		Where("attendance_id = ? AND punch_in = ?", attendanceID, punchInTime).
		Select()
	fmt.Println(err)
	if err != nil {
		return err
	}

	// Fetch the ClassID based on the provided class name
	var classInfo attendanceModels.Class
	err = impl.dbConnection.Model(&classInfo).
		Column("class_id").
		Where("class_name = ?", className).
		Select()

	if err != nil {
		return err
	}
	// Create a new ClassMappingAttendance record
	_, err = impl.dbConnection.Model(&attendanceModels.ClassMappingAttendance{
		PunchID: punch.ID,
		ClassID: classInfo.ClassID,
	}).Insert()
	return err
}
func (impl AttendanceRepository) AddAttendance(userid string, currentDate time.Time, newAttendanceID int) (int, error) {

	var atID int
	_, err := impl.dbConnection.Model(&attendanceModels.Attendance{
		UserID: userid,
		Day:    currentDate.Day(),
		Month:  int(currentDate.Month()),
		Year:   currentDate.Year(),
	}).Returning("id").Insert(&atID)
	//_, err := impl.dbConnection.Model(&studentAttendance).Returning("id").Insert(&newAttendanceID)
	return atID, err
}
func (impl AttendanceRepository) PunchOut(userID string, id int) error {

	var existingPunch attendanceModels.PunchInOut
	err := impl.dbConnection.Model(&existingPunch).
		Column("attendance_id").
		Where("user_id = ? AND punch_id = ?", userID, id).
		Select()
	return err
}
func (impl AttendanceRepository) PunchOutCheck(userid string, classid int) error {
	var existingPunch attendanceModels.PunchInOut
	err := impl.dbConnection.Model(&existingPunch).
		Join("JOIN class_mapping_attendances ON class_mapping_attendances.punch_id = punch_in_out.id").
		Where("punch_in_out.user_id = ? AND class_mapping_attendances.class_id = ? AND punch_in_out.punch_out IS NULL ", userid, classid).
		Select()
	return err
}
