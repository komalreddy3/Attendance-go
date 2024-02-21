package loginRepository

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginRepository/loginModels"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type LoginRepository struct {
	dbConnection *pg.DB
	logger       *zap.SugaredLogger
}
type LoginRepo interface {
	CheckCreds(username, userRole string) (string, error)
	AuthenticateUser(username, password, userRole string) bool
	CheckRole(content interface{}) loginModels.UserRoleType
}

func NewLoginRepositoryImpl(dbConnection *pg.DB, logger *zap.SugaredLogger) *LoginRepository {
	return &LoginRepository{
		dbConnection: dbConnection,
		logger:       logger,
	}
}
func (impl LoginRepository) CheckCreds(username, userRole string) (string, error) {
	var user loginModels.User
	var role loginModels.UserRoleType
	if userRole == "principal" {
		role = loginModels.Principal
	}
	if userRole == "teacher" {
		role = loginModels.Teacher
	}
	if userRole == "student" {
		role = loginModels.Student
	}
	err := impl.dbConnection.Model((*loginModels.User)(nil)).
		Column("id", "role").
		Where("username = ? AND role = ?", username, role).
		Select(&user)

	return user.ID, err
}
func (impl LoginRepository) AuthenticateUser(username, password, userRole string) bool {
	var user loginModels.User
	var role loginModels.UserRoleType
	if userRole == "principal" {
		role = loginModels.Principal
		return true
	}
	if userRole == "teacher" {
		role = loginModels.Teacher
	}
	if userRole == "student" {
		role = loginModels.Student
	}
	err := impl.dbConnection.Model(&user).
		Where("username = ? AND role = ?", username, role).
		Select()
	fmt.Println(err)
	if err != nil {
		return false
	}

	// Compare the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(err)
	return err == nil
}
func (impl LoginRepository) CheckRole(content interface{}) loginModels.UserRoleType {
	var user loginModels.User
	err := impl.dbConnection.Model(&user).
		Where("id = ? ", content).
		Select()
	if err != nil {
		log.Fatalln(err)
	}
	return user.Role
}
