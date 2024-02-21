package userRepository

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"go.uber.org/zap"
)

type UserRepository struct {
	dbConnection *pg.DB
	logger       *zap.SugaredLogger
}
type UserRepo interface {
	CheckEnrollment(userid, class string) error
	CheckPunchOut(userid string) ([]string, error)
	FetchClass(enrolledClass string) int
	FetchStudent(userid string) string
	InsertingStudent(id, username, password string) error
	InsertingTeacher(id, username, password string) error
	InsertClass(className string) (int, error)
	InsertClassMap(id string, classId int) error
	ClassMappingTeacher(classname []string) (map[string]int, error)
	FetchUser(role string) []userModels.User
	FetchClassUser(userid string) int
	FetchClasses(username string) []string
	AllClasses() []string
}

func NewUserRepositoryImpl(dbConnection *pg.DB, logger *zap.SugaredLogger) *UserRepository {
	return &UserRepository{
		dbConnection: dbConnection,
		logger:       logger,
	}
}
func (impl UserRepository) AllClasses() []string {
	// Fetch all class names
	var classes []userModels.Class
	err := impl.dbConnection.Model(&classes).Select()
	if err != nil {
		//log.Fatalf("Error fetching class names: %v\n", err)
		impl.logger.Errorw("error fetching class name")
		return nil
	}

	var classNames []string
	for _, class := range classes {
		classNames = append(classNames, class.ClassName)
	}

	fmt.Println("Class Names:")
	for _, name := range classNames {
		fmt.Println(name)
	}
	return classNames
}
func (impl UserRepository) FetchClasses(username string) []string {
	// Loop through each student and fetch their associated class names
	//for i, student := range students {
	var classMappingUsers []userModels.ClassMappingUser
	var userID string
	err := impl.dbConnection.Model(&userModels.User{}).
		Column("id").
		Where("username=?", username).
		Select(&userID)
	if err != nil {
		impl.logger.Errorw("Error fetching student", err)
	}
	err = impl.dbConnection.Model(&classMappingUsers).
		Where("user_id = ?", userID).
		Select()
	//fmt.Println(userID)
	//fmt.Println(classMappingUsers)
	//fmt.Println(len(classMappingUsers))
	// Create a map to store class names by class ID
	classIDToName := make(map[int]string)

	// Extract class IDs from classMappingUsers and collect unique IDs
	var classIDs []int
	for _, mapping := range classMappingUsers {
		classIDs = append(classIDs, mapping.ClassID)
	}

	// Query the Class table to fetch class names for each class ID
	var classes []userModels.Class
	err = impl.dbConnection.Model(&classes).
		Where("class_id IN (?)", pg.In(classIDs)).
		Select()
	if err != nil {
		// Handle the error
	}

	// Populate classIDToName map with class names
	for _, classInfo := range classes {
		classIDToName[classInfo.ClassID] = classInfo.ClassName
	}
	//// Extract class names from classMappingUsers
	//classnames := make([]string, len(classMappingUsers))
	//for j, mapping := range classMappingUsers {
	//	classnames[j] = mapping.Class.ClassName
	//}
	// Extract class names from classMappingUsers
	classnames := make([]string, len(classMappingUsers))
	for j, mapping := range classMappingUsers {
		// Get class name from classIDToName map
		classnames[j] = classIDToName[mapping.ClassID]
	}

	// Assign classnames to the student
	//students[i].Classnames = classnames
	//}
	return classnames
}
func (impl UserRepository) FetchClassUser(userid string) int {
	var classMapping userModels.ClassMappingUser
	_ = impl.dbConnection.Model(&classMapping).
		Where("user_id = ?", userid).
		Select()
	return classMapping.ClassID
}
func (impl UserRepository) CheckEnrollment(userid, class string) error {
	// Check if the user is enrolled in the class
	var classMappingUser userModels.ClassMappingUser
	fmt.Println("in check enrollemnt", class, userid)
	err := impl.dbConnection.Model(&classMappingUser).
		Where("user_id = ? AND class_id IN (SELECT class_id FROM classes WHERE class_name = ?)", userid, class).
		Select()
	return err
}
func (impl UserRepository) CheckPunchOut(userid string) ([]string, error) {
	var enrolledClass []string
	err := impl.dbConnection.Model(&userModels.ClassMappingUser{}).
		ColumnExpr("DISTINCT class_name").
		Join("JOIN classes ON class_mapping_user.class_id = classes.class_id").
		Where("user_id = ? ", userid).
		Select(&enrolledClass)
	return enrolledClass, err
}
func (impl UserRepository) FetchClass(enrolledClass string) int {
	var classInfo userModels.Class
	err := impl.dbConnection.Model(&classInfo).
		Column("class_id").
		Where("class_name = ?", enrolledClass).
		Select()

	if err != nil {
		impl.logger.Errorw("Error fetching class ID", "error", err)
		return 0
	}
	return classInfo.ClassID
}
func (impl UserRepository) FetchStudent(userid string) string {
	var userID string
	err := impl.dbConnection.Model(&userModels.User{}).
		Column("id").
		Where("role = ? AND id=?", userModels.Student, userid).
		Select(&userID)
	if err != nil {
		impl.logger.Errorw("Error fetching student", err)
	}
	return userID
}
func (impl UserRepository) FetchUser(role string) []userModels.User {
	var students []userModels.User

	err := impl.dbConnection.Model(&students).
		Column("id", "username").
		Where("role = ?", role).
		Select()
	if err != nil {
		impl.logger.Errorw("Error fetching user data", "error", err)

		return nil
	}
	return students
}
func (impl UserRepository) InsertingStudent(id, username, password string) error {
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {
	//	impl.logger.Errorw("Error hashing the password", "error", err)
	//	// http.Error(w, "Error hashing the password", http.StatusInternalServerError)
	//	return err
	//}
	_, err := impl.dbConnection.Model(&userModels.User{ID: id, Username: username, Password: password, Role: userModels.Student}).Insert()
	if err != nil {
		impl.logger.Errorw("Error inserting student data into the database", "error", err)
		return err
	}
	return err
}
func (impl UserRepository) InsertingTeacher(id, username, password string) error {
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {
	//	impl.logger.Errorw("Error hashing the password", "error", err)
	//	// http.Error(w, "Error hashing the password", http.StatusInternalServerError)
	//	return err
	//}
	_, err := impl.dbConnection.Model(&userModels.User{ID: id, Username: username, Password: password, Role: userModels.Teacher}).Insert()
	if err != nil {
		impl.logger.Errorw("Error inserting student data into the database", "error", err)
		return err
	}
	return err
}
func (impl UserRepository) InsertClass(classname string) (int, error) {
	newClass := userModels.Class{
		ClassName: classname,
	}
	var id int
	err := impl.dbConnection.Model(&userModels.Class{}).
		Column("class_id").
		Where("class_name = ?", classname).
		Select(&id)
	if err != nil {
		_, err := impl.dbConnection.Model(&newClass).Returning("class_id").Insert()
		if err != nil {
			impl.logger.Errorw("Error inserting new class data", "error", err)
		}
		return newClass.ClassID, err
	}
	return 0, err

}
func (impl UserRepository) InsertClassMap(id string, classId int) error {
	classMapping := userModels.ClassMappingUser{
		UserID:  id,
		ClassID: classId,
	}
	_, err := impl.dbConnection.Model(&classMapping).Insert()
	if err != nil {
		impl.logger.Errorw("Error inserting class mapping data into the database", "error", err)
		return err
	}
	return err
}
func (impl UserRepository) ClassMappingTeacher(classname []string) (map[string]int, error) {
	classIDMap := make(map[string]int)
	// Select class_id and class_name from the database
	var rows []struct {
		ClassID   int
		ClassName string
	}
	err := impl.dbConnection.Model(&userModels.Class{}).Column("class_id", "class_name").Where("class_name IN (?)", pg.In(classname)).Select(&rows)
	for _, row := range rows {
		classIDMap[row.ClassName] = row.ClassID
	}
	fmt.Println(err)
	return classIDMap, err
}
