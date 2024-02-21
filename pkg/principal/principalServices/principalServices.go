package principalServices

import (
	"fmt"
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
	AddStudent(StudentID, ClassName string) string
	AddTeacher(TeacherID string, ClassNames []string) string
	GetTeacherAttendance(ID string, Month int, Year int) principalServiceBean.TeacherAttendanceResponse
	PrincipalStudents() []userServiceBean.CustomUserInfo
	PrincipalTeachers() []userServiceBean.CustomUserInfo
	ClassInsert(className string) int
	PrincipalClasses() []string
}

func NewPrincipalServiceImpl(principalRepository principalRepository.PrincipalRepo, logger *zap.SugaredLogger) *PrincipalServiceImpl {
	return &PrincipalServiceImpl{
		principalRepository,
		logger,
	}
}
func (impl *PrincipalServiceImpl) PrincipalClasses() []string {
	return impl.principalRepository.AllClasses()
}
func (impl *PrincipalServiceImpl) ClassInsert(className string) int {
	id, _ := impl.principalRepository.InsertClass(className)
	return id
}
func (impl *PrincipalServiceImpl) AddStudent(StudentID, ClassName string) string {
	//server.AuthenticateRole(w, r, "principal")
	if impl.principalRepository.CheckAlreadyInsert(StudentID, "student") != nil {
		return "Already inserted student with the same id"
	}
	// Set the student's password to be the same as the username
	password := StudentID

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("this is passwrd by principal services", hashedPassword)
	fmt.Println("this is passwrd by principal services", string(hashedPassword))
	fmt.Println("this is passwrd by principal services", []byte(string(hashedPassword)))

	if err != nil {
		impl.logger.Errorw("Error hashing the password", "error", err)
		return "Problem with password"
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
	return ""
}
func (impl *PrincipalServiceImpl) AddTeacher(TeacherID string, ClassNames []string) string {
	//server.AuthenticateRole(w, r, "principal")
	if impl.principalRepository.CheckAlreadyInsert(TeacherID, "teacher") != nil {
		return "Already inserted teacher with the same id"
	}
	// Set the teacher's password to be the same as the username
	password := TeacherID
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		impl.logger.Errorw("Error hashing the password", "error", err)
		return "Problem with hashing password"
	}
	// Insert teacher data into the postgres database
	impl.principalRepository.InsertingTeacher(TeacherID, TeacherID, string(hashedPassword))
	// Map existing class names to their IDs
	classIDMap, err := impl.principalRepository.ClassMappingTeacher(ClassNames)
	fmt.Println("classNmaes", ClassNames)
	fmt.Println(len(ClassNames))
	fmt.Println("classIDMap", classIDMap)
	// Iterate over the map and print each key-value pair
	fmt.Println("Class ID Mapping:")
	for className, classID := range classIDMap {
		fmt.Printf("Class Name: %s, Class ID: %d\n", className, classID)
	}
	// Insert new classes and map them to the teacher
	for _, className := range ClassNames {
		// Check if the class name already exists
		classID, exists := classIDMap[className]
		fmt.Println("classIDMap[className] className exists", classIDMap[className], className, exists)
		done := false
		if !exists {
			classID, err := impl.principalRepository.InsertClass(className)
			if err != nil {
				impl.logger.Errorw("Inserting class error", err)
			}
			// Update the classIDMap with the new class ID
			classIDMap[className] = classID
			impl.principalRepository.InsertClassMap(TeacherID, classID)
			done = true
		}
		// Map the teacher to the class in ClassMappingUser
		if done == false {
			impl.principalRepository.InsertClassMap(TeacherID, classID)
		}

	}
	return ""
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
			if entry.LastPunchOut != "" {
				classPunches[className] = entry
			}
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
	for i, _ := range students {
		students[i].Classnames = impl.principalRepository.FetchClasses(students[i].Username)
	}
	return students

}
func (impl *PrincipalServiceImpl) PrincipalTeachers() []userServiceBean.CustomUserInfo {
	//server.AuthenticateRole(w, r, "principal")
	// Fetch teachers and their associated classes
	var teachers []userServiceBean.CustomUserInfo
	teachers = impl.principalRepository.FetchUser("teacher")
	for i, _ := range teachers {
		teachers[i].Classnames = impl.principalRepository.FetchClasses(teachers[i].Username)
	}
	return teachers

}
