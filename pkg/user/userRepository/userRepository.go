package userRepository

import (
	"github.com/go-pg/pg/v10"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
}

func NewUserRepositoryImpl(dbConnection *pg.DB, logger *zap.SugaredLogger) *UserRepository {
	return &UserRepository{
		dbConnection: dbConnection,
		logger:       logger,
	}
}
func (impl UserRepository) CheckEnrollment(userid, class string) error {
	// Check if the user is enrolled in the class
	var classMappingUser userModels.ClassMappingUser
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
		Where("role = ? AND id=?", "student", userid).
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		impl.logger.Errorw("Error hashing the password", "error", err)
		// http.Error(w, "Error hashing the password", http.StatusInternalServerError)
		return err
	}
	_, err = impl.dbConnection.Model(&userModels.User{ID: id, Username: username, Password: string(hashedPassword), Role: userModels.Student}).Insert()
	if err != nil {
		impl.logger.Errorw("Error inserting student data into the database", "error", err)
		return err
	}
	return err
}
func (impl UserRepository) InsertingTeacher(id, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		impl.logger.Errorw("Error hashing the password", "error", err)
		// http.Error(w, "Error hashing the password", http.StatusInternalServerError)
		return err
	}
	_, err = impl.dbConnection.Model(&userModels.User{ID: id, Username: username, Password: string(hashedPassword), Role: userModels.Teacher}).Insert()
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
	_, err := impl.dbConnection.Model(&newClass).Returning("class_id").Insert()
	if err != nil {
		impl.logger.Errorw("Error inserting new class data", "error", err)
	}
	return newClass.ClassID, err
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
		ClassName string
		ClassID   int
	}
	err := impl.dbConnection.Model(&userModels.Class{}).Column("class_id", "class_name").Where("class_name IN (?)", pg.In(classname)).Select(&rows)
	for _, row := range rows {
		classIDMap[row.ClassName] = row.ClassID
	}
	return classIDMap, err
}
