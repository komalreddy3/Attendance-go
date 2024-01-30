package principalServices

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/principal/principalRepository"
	"github.com/komalreddy3/Attendance-go/pkg/principal/principalServices/principalServiceBean"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices/userServiceBean"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type PrincipalServiceImpl struct {
	principalRepository principalRepository.PrincipalRepo
	logger              *zap.SugaredLogger
}
type PrincipalService interface {
	AddStudent(StudentID, ClassName string)
	AddTeacher(TeacherID string, ClassNames []string)
	GetTeacherAttendance(ID string, Month int, Year int) principalServiceBean.TeacherAttendanceResponse
	PrincipalStudents() []userServiceBean.CustomUserInfo
	PrincipalTeachers() []userServiceBean.CustomUserInfo
}

func NewPrincipalServiceImpl(principalRepository principalRepository.PrincipalRepo, logger *zap.SugaredLogger) *PrincipalServiceImpl {
	return &PrincipalServiceImpl{
		principalRepository,
		logger,
	}
}
func (impl *PrincipalServiceImpl) AddStudent(StudentID, ClassName string) {
	//server.AuthenticateRole(w, r, "principal")
	// Set the student's password to be the same as the username
	password := StudentID

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		impl.logger.Errorw("Error hashing the password", "error", err)
		return
	}
	// Insert student data into the postgres database
	impl.principalRepository.InsertingStudent(StudentID, StudentID, string(hashedPassword))
	// Map the class name to its ID
	classID := impl.principalRepository.ClassMapping(ClassName)
	// If the class doesn't exist, insert it
	if classID == 0 {
		classID, err = impl.principalRepository.InsertClass(ClassName)
		if err != nil {
			impl.logger.Errorw("error inserting class", err)
		}
	}
	// Map the student to the class in ClassMappingUser
	impl.principalRepository.InsertClassMap(StudentID, classID)
	//w.WriteHeader(http.StatusOK)
}
func (impl *PrincipalServiceImpl) AddTeacher(TeacherID string, ClassNames []string) {
	//server.AuthenticateRole(w, r, "principal")

	// Set the teacher's password to be the same as the username
	password := TeacherID
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		impl.logger.Errorw("Error hashing the password", "error", err)
		return
	}
	// Insert teacher data into the postgres database
	impl.principalRepository.InsertingTeacher(TeacherID, TeacherID, string(hashedPassword))
	// Map existing class names to their IDs
	classIDMap, err := impl.principalRepository.ClassMappingTeacher(ClassNames)
	// Insert new classes and map them to the teacher
	for _, className := range ClassNames {
		// Check if the class name already exists
		classID, exists := classIDMap[className]
		if !exists {
			classID, err := impl.principalRepository.InsertClass(className)
			if err != nil {
				impl.logger.Errorw("Inserting class error", err)
			}
			// Update the classIDMap with the new class ID
			classIDMap[className] = classID
		}
		// Map the teacher to the class in ClassMappingUser
		impl.principalRepository.InsertClassMap(TeacherID, classID)
	}
}
func (impl *PrincipalServiceImpl) GetTeacherAttendance(ID string, Month int, Year int) principalServiceBean.TeacherAttendanceResponse {
	//server.AuthenticateRole(w, r, "principal")

	var ids []int
	//ids = impl.principalRepository.FetchAttendance(ID, Month, Year)
	ids = impl.principalRepository.FetchAttendance(ID, Month, Year)
	// Define a struct to hold attendance information with type
	result := make(map[int][]principalServiceBean.AttendanceEntry)
	for _, attendance := range ids {
		day := impl.principalRepository.FetchDay(attendance)
		var punchInOutRecords []attendanceServiceBean.PunchRecord
		//punchInOutRecords = impl.principalRepository.FetchPunch(attendance)
		punchInOutRecords = impl.principalRepository.FetchPunch(attendance)
		if _, ok := result[day]; !ok {
			result[day] = make([]principalServiceBean.AttendanceEntry, 0)
		}
		// Create a map to store entries for each class
		classPunches := make(map[string]principalServiceBean.AttendanceEntry)
		for _, record := range punchInOutRecords {
			className := impl.principalRepository.ClassMapAttendancePunch(record.ID)
			// Check if class entry exists, and update first punch-in and last punch-out
			entry, exists := classPunches[className]
			if !exists {
				entry = principalServiceBean.AttendanceEntry{Class: className, FirstPunchIn: record.PunchIn}
			}
			entry.LastPunchOut = record.PunchOut
			classPunches[className] = entry
		}
		// Add entries for each class to the result
		for _, v := range classPunches {
			result[day] = append(result[day], v)
		}
	}
	response := principalServiceBean.TeacherAttendanceResponse{
		ID:         ID,
		Month:      Month,
		Year:       Year,
		Attendance: result,
	}

	return response
}
func (impl *PrincipalServiceImpl) PrincipalStudents() []userServiceBean.CustomUserInfo {
	//server.AuthenticateRole(w, r, "principal")
	// Fetch students and their associated classes
	var students []userServiceBean.CustomUserInfo
	students = impl.principalRepository.FetchUser("student")
	return students

}
func (impl *PrincipalServiceImpl) PrincipalTeachers() []userServiceBean.CustomUserInfo {
	//server.AuthenticateRole(w, r, "principal")
	// Fetch teachers and their associated classes
	var teachers []userServiceBean.CustomUserInfo
	teachers = impl.principalRepository.FetchUser("teacher")
	return teachers

}
