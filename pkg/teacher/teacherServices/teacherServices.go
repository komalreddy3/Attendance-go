package teacherServices

import (
	"errors"
	"fmt"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/teacher/teacherRepository"
	"github.com/komalreddy3/Attendance-go/pkg/teacher/teacherServices/teacherServiceBean"
	"go.uber.org/zap"
	"time"
)

type TeacherServiceImpl struct {
	teacherRepository teacherRepository.TeacherRepo
	logger            *zap.SugaredLogger
}
type TeacherService interface {
	TeacherPunchIn(userID, class string) error
	TeacherPunchOut(userID, class string) error
	GetTeacherAttendanceByMonth(ID string, Month int, Year int) teacherServiceBean.TeacherAttendanceResponse
	GetClassAttendance(userID string, Class string, Day int, Month int, Year int) teacherServiceBean.ClassAttendanceResponse
}

func NewTeacherServiceImpl(teacherRepository teacherRepository.TeacherRepo, logger *zap.SugaredLogger) *TeacherServiceImpl {
	return &TeacherServiceImpl{
		teacherRepository: teacherRepository,
		logger:            logger,
	}
}
func (impl *TeacherServiceImpl) TeacherPunchIn(userID, class string) error {
	//server.AuthenticateRole(w, r, "teacher")

	// Check if the user is enrolled in the class
	err := impl.teacherRepository.EnrollCheckTeacher(userID, class)
	if err != nil {
		return err
	}
	// Check if the user has already punched out from any class on the same day
	enrolledClasses := impl.teacherRepository.PunchCheckTeacher(userID)
	for _, enrolledClass := range enrolledClasses {
		// Check if the user has an entry in the attendance table on the same day for the given class
		_, err := impl.teacherRepository.FetchAttendance(userID)
		if err == nil {
			// User has an entry in the attendance table on the same day for the given class
			classId := impl.teacherRepository.FetchClass(enrolledClass)
			// Check if the user has punched out
			err = impl.teacherRepository.PunchOutCheck(userID, classId)
			if err == nil {
				impl.logger.Errorw("You haven't punched out from "+enrolledClass+" yet", "error", err)
				return errors.New("not punched out")
			}
		}
	}
	// Check if the user has an entry in the attendance table on the same day for the given class
	existid, err := impl.teacherRepository.FetchAttendance(userID)
	if err == nil {
		// User has an entry in the attendance table on the same day for the requested class

		// Check if the user has punched out
		// Fetch the ClassID based on the provided class name
		classId := impl.teacherRepository.FetchClass(class)
		fmt.Println(classId)
		// Check if the user has punched out
		err := impl.teacherRepository.PunchOutCheck(userID, classId)
		if err == nil {
			err = impl.teacherRepository.CreatePunchInTeacher(userID, time.Now(), existid, class)
			if err != nil {
				impl.logger.Errorw("Error recording punch-in", "error", err)
				return err
			}
			//w.WriteHeader(http.StatusOK)
			return nil
		}

	}
	// Check if the user has an entry in the attendance table on the same day for the given class
	_, err = impl.teacherRepository.FetchAttendance(userID)
	if err == nil {
		// User has an entry in the attendance table on the same day for the requested class
		// Check if the user has punched out
		err := impl.teacherRepository.PunchOut(userID, existid)
		if err == nil {
			fmt.Println("called")
			err = impl.teacherRepository.CreatePunchInTeacher(userID, time.Now(), existid, class)
			if err != nil {
				impl.logger.Errorw("Error recording punch-in", "error", err)
				return err
			}
			//w.WriteHeader(http.StatusOK)
			return nil
		}
	}
	//if err == nil {
	//	newAttendanceID, err := impl.teacherRepository.FetchAttendance(userID)
	//	if err == nil {
	//		newAttendanceID, err = impl.teacherRepository.AddAttendance(userID, time.Now(), newAttendanceID)
	//		if err != nil {
	//			impl.logger.Errorw("Error creating or updating attendance record", "error", err)
	//			return
	//		}
	//	}
	//}
	////newAttendanceID, err := impl.teacherRepository.FetchAttendance(userID)
	//err = impl.teacherRepository.CreatePunchIn(userID, time.Now(), newAttendanceID, class)
	//if err != nil {
	//	impl.logger.Errorw("Error recording punch-in", "error", err)
	//	return
	//}
	var newAttendanceID int
	newAttendanceID, err = impl.teacherRepository.FetchAttendance(userID)
	fmt.Println("NEW ATTENDNACE ID userid", newAttendanceID, userID)
	if err != nil {
		newAttendanceID, err = impl.teacherRepository.AddAttendance(userID, time.Now(), newAttendanceID)
		if err != nil {
			impl.logger.Errorw("Error creating or updating attendance record", "error", err)
			return err
		}
	}
	fmt.Println(newAttendanceID)
	err = impl.teacherRepository.CreatePunchInTeacher(userID, time.Now(), newAttendanceID, class)
	if err != nil {
		impl.logger.Errorw("Error recording punch-in", "error", err)
		return err
	}
	//w.WriteHeader(http.StatusOK)
	return nil

}
func (impl *TeacherServiceImpl) TeacherPunchOut(userID, class string) error {
	//server.AuthenticateRole(w, r, "student")
	currentDate := time.Now()
	// Check if the user is enrolled in the class
	//impl.teacherRepository.EnrollCheckTeacher(userID, class)
	existid, err := impl.teacherRepository.FetchAttendance(userID)
	if err != nil {
		impl.logger.Errorw("No attendance record found for the user on the specified day", "error", err)
		return err
	}
	classname := impl.teacherRepository.PunchOutNull(userID)
	classid := impl.teacherRepository.FetchClass(classname)
	// Check if the user has punched in for the given class and attendance record
	err = impl.teacherRepository.PunchOutCheck(userID, classid)
	if err != nil {
		impl.logger.Errorw("You haven't punched in for "+class+" yet", "error", err)
		return err
	}
	// Update the existing punch record with punch-out time
	impl.teacherRepository.UpdatePunchOut(existid, classid, currentDate)
	//w.WriteHeader(http.StatusOK)
	return nil
}

func (impl *TeacherServiceImpl) GetTeacherAttendanceByMonth(ID string, Month int, Year int) teacherServiceBean.TeacherAttendanceResponse {
	//server.AuthenticateRole(w, r, "student")

	existsid := impl.teacherRepository.FetchAttendanceWithMonth(ID, Month, Year)
	// Define a struct to hold attendance information with type
	result := make(map[int][]teacherServiceBean.AttendanceEntry)
	for _, attendance := range existsid {
		day := impl.teacherRepository.FetchDay(attendance)
		var punchInOutRecords []attendanceServiceBean.PunchRecord
		punchInOutRecords = impl.teacherRepository.FetchPunch(attendance)
		if _, ok := result[day]; !ok {
			result[day] = make([]teacherServiceBean.AttendanceEntry, 0)
		}
		// Create a map to store entries for each class
		classPunches := make(map[string]teacherServiceBean.AttendanceEntry)
		for _, record := range punchInOutRecords {
			className := impl.teacherRepository.ClassMapAttendancePunch(record.ID)
			// Check if class entry exists, and update first punch-in and last punch-out
			entry, exists := classPunches[className]
			if !exists {
				entry = teacherServiceBean.AttendanceEntry{Class: className, FirstPunchIn: record.PunchIn}
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
	response := teacherServiceBean.TeacherAttendanceResponse{
		ID:         ID,
		Month:      Month,
		Year:       Year,
		Attendance: result,
	}

	return response

}

func (impl *TeacherServiceImpl) GetClassAttendance(userID string, Class string, Day int, Month int, Year int) teacherServiceBean.ClassAttendanceResponse {
	//server.AuthenticateRole(w, r, "teacher")
	// Check if the user is enrolled in the class
	err := impl.teacherRepository.EnrollCheckTeacher(userID, Class)
	if err != nil {
		return teacherServiceBean.ClassAttendanceResponse{}
	}
	// Fetch attendance records for the student for the given month and year
	var ids []int
	ids = impl.teacherRepository.FetchStudentAttendance(Day, Month, Year)
	// Define a struct to hold attendance information with type

	result := make(map[int][]teacherServiceBean.ClassAttendanceEntry)
	for _, attendance := range ids {
		day := impl.teacherRepository.FetchDay(attendance)
		var punchInOutRecords []attendanceServiceBean.PunchRecord
		punchInOutRecords = impl.teacherRepository.FetchPunch(attendance)
		if _, ok := result[day]; !ok {
			result[day] = make([]teacherServiceBean.ClassAttendanceEntry, 0)
		}
		// Create a map to store entries for each class
		classPunches := make(map[string]teacherServiceBean.ClassAttendanceEntry)
		for _, record := range punchInOutRecords {
			var userID string
			userID = impl.teacherRepository.FetchStudent(record.UserID)
			var className string
			className = impl.teacherRepository.ClassMapAttendancePunch(record.ID)
			if className == Class {
				//// Check if class entry exists, and update first punch-in and last punch-out
				//entry, exists := classPunches[className]
				//if !exists {
				//	entry = teacherServiceBean.ClassAttendanceEntry{Id: userID, FirstPunchIn: record.PunchIn}
				//}
				//entry.LastPunchOut = record.PunchOut
				//classPunches[className] = entry
				if userID != "" {
					// Check if class entry exists, and update first punch-in and last punch-out
					entry, exists := classPunches[className]
					if !exists {
						entry = teacherServiceBean.ClassAttendanceEntry{Id: userID, FirstPunchIn: record.PunchIn}
					}
					entry.LastPunchOut = record.PunchOut
					if entry.LastPunchOut != "" {
						classPunches[className] = entry
					}
				}
			}
		}
		// Add entries for each class to the result
		for _, v := range classPunches {
			result[day] = append(result[day], v)
		}
	}

	response := teacherServiceBean.ClassAttendanceResponse{
		Day:        Day,
		Month:      Month,
		Year:       Year,
		Attendance: result,
	}
	return response

}
