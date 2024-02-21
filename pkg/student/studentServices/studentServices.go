package studentServices

import (
	"errors"
	"fmt"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/student/studentRepository"
	"github.com/komalreddy3/Attendance-go/pkg/student/studentServices/studentServiceBean"
	"go.uber.org/zap"

	"time"
)

type StudentServiceImpl struct {
	studentRepository studentRepository.StudentRepo
	logger            *zap.SugaredLogger
}
type StudentService interface {
	StudentPunchIn(userID, class string) error
	StudentPunchOut(userID, class string) error
	GetStudentAttendanceByMonth(ID string, Month int, Year int) studentServiceBean.TeacherAttendanceResponse
	RemovePunchIn(userID, class string) error
}

func NewStudentServiceImpl(studentRepository studentRepository.StudentRepo, logger *zap.SugaredLogger) *StudentServiceImpl {
	return &StudentServiceImpl{
		studentRepository: studentRepository,
		logger:            logger,
	}
}
func (impl *StudentServiceImpl) RemovePunchIn(userID, class string) error {
	// Check if the user has already punched out from any class on the same day
	enrolledClass := impl.studentRepository.PunchCheck(userID)
	classId := impl.studentRepository.FetchClass(enrolledClass)
	// Check if the user has punched out
	err := impl.studentRepository.PunchOutCheck(userID, classId)
	return err
}
func (impl *StudentServiceImpl) StudentPunchIn(userID, class string) error {
	//server.AuthenticateRole(w, r, "student")

	// Check if the user is enrolled in the class
	//impl.studentRepository.EnrollCheck(userID, class)
	// Check if the user has already punched out from any class on the same day
	enrolledClass := impl.studentRepository.PunchCheck(userID)
	// Check if the user has an entry in the attendance table on the same day for the given class
	existid, err := impl.studentRepository.FetchAttendance(userID)
	if err == nil {
		// User has an entry in the attendance table on the same day for the given class
		classId := impl.studentRepository.FetchClass(enrolledClass)
		// Check if the user has punched out
		err = impl.studentRepository.PunchOutCheck(userID, classId)
		if err == nil {
			impl.logger.Errorw("You haven't punched out from "+enrolledClass+" yet", "error", err)
			//return err
			return errors.New("not punched out")
		}
	}
	if err == nil {
		// User has an entry in the attendance table on the same day for the requested class

		// Check if the user has punched out
		// Fetch the ClassID based on the provided class name
		//classId := impl.studentRepository.FetchClass(enrolledClass)
		// Check if the user has punched out
		//err := impl.studentRepository.PunchOutCheck(userID, classId)
		if err == nil {
			err = impl.studentRepository.CreatePunchIn(userID, time.Now(), existid, class)
			if err != nil {
				impl.logger.Errorw("Error recording punch-in", "error", err)
				return err
			}
			//w.WriteHeader(http.StatusOK)
			return nil
		}
	}
	// No existing attendance record for the user and class on the same day, create a new one
	var newAttendanceID int
	newAttendanceID, err = impl.studentRepository.FetchAttendance(userID)
	fmt.Println("NEW ATTENDNACE ID userid", newAttendanceID, userID)
	if err != nil {
		newAttendanceID, err = impl.studentRepository.AddAttendance(userID, time.Now(), newAttendanceID)
		if err != nil {
			impl.logger.Errorw("Error creating or updating attendance record", "error", err)
			return err
		}
	}
	fmt.Println(newAttendanceID)
	err = impl.studentRepository.CreatePunchIn(userID, time.Now(), newAttendanceID, class)
	if err != nil {
		impl.logger.Errorw("Error recording punch-in", "error", err)
		//return err
	}
	//w.WriteHeader(http.StatusOK)
	return nil
}

func (impl *StudentServiceImpl) StudentPunchOut(userID, class string) error {
	//server.AuthenticateRole(w, r, "student")

	currentDate := time.Now()
	// Check if the user is enrolled in the class
	//impl.studentRepository.EnrollCheck(userID, class)
	existid, err := impl.studentRepository.FetchAttendance(userID)
	if err != nil {
		impl.logger.Errorw("No attendance record found for the user on the specified day", "error", err)
		return err
	}
	fmt.Println("existid", existid)
	classid := impl.studentRepository.FetchClassUser(userID)
	// Check if the user has punched in for the given class and attendance record
	err = impl.studentRepository.PunchOutCheck(userID, classid)
	if err != nil {
		impl.logger.Errorw("You haven't punched in for "+class+" yet", "error", err)
		return err
	}
	// Update the existing punch record with punch-out time
	impl.studentRepository.UpdatePunchOut(existid, classid, currentDate)
	//w.WriteHeader(http.StatusOK)
	return nil

}
func (impl *StudentServiceImpl) GetStudentAttendanceByMonth(ID string, Month int, Year int) studentServiceBean.TeacherAttendanceResponse {
	//server.AuthenticateRole(w, r, "student")

	existsid := impl.studentRepository.FetchAttendanceWithMonth(ID, Month, Year)
	// Define a struct to hold attendance information with type
	result := make(map[int][]studentServiceBean.AttendanceEntry)
	for _, attendance := range existsid {
		day := impl.studentRepository.FetchDay(attendance)
		var punchInOutRecords []attendanceServiceBean.PunchRecord
		punchInOutRecords = impl.studentRepository.FetchPunch(attendance)
		if _, ok := result[day]; !ok {
			result[day] = make([]studentServiceBean.AttendanceEntry, 0)
		}
		// Create a map to store entries for each class
		classPunches := make(map[string]studentServiceBean.AttendanceEntry)
		for _, record := range punchInOutRecords {
			className := impl.studentRepository.ClassMapAttendancePunch(record.ID)
			// Check if class entry exists, and update first punch-in and last punch-out
			entry, exists := classPunches[className]
			if !exists {
				entry = studentServiceBean.AttendanceEntry{Class: className, FirstPunchIn: record.PunchIn}
			}
			entry.LastPunchOut = record.PunchOut
			if entry.LastPunchOut != "" {
				classPunches[className] = entry
			}
		}
		// Add entries for each class to the result
		for _, v := range classPunches {
			result[day] = append(result[day], v)
		}
	}
	response := studentServiceBean.TeacherAttendanceResponse{
		ID:         ID,
		Month:      Month,
		Year:       Year,
		Attendance: result,
	}

	return response

}
