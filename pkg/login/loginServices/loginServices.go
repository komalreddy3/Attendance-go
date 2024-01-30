package loginServices

import (
	"fmt"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginRepository"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginRepository/loginModels"
	"go.uber.org/zap"
	"log"
	"net/http"

	"os"
)

type LoginServiceImpl struct {
	loginRepository loginRepository.LoginRepo
	logger          *zap.SugaredLogger
}
type LoginService interface {
	NewJWT(PrivateKey []byte, PublicKey []byte) loginModels.JWT
	AuthenticateRole(cookie *http.Cookie, check string) bool
	Login(username, password, userRole string) string
}

func NewLoginServiceImpl(loginRepository loginRepository.LoginRepo, logger *zap.SugaredLogger) *LoginServiceImpl {
	return &LoginServiceImpl{
		loginRepository,
		logger,
	}
}
func (impl *LoginServiceImpl) NewJWT(PrivateKey []byte, PublicKey []byte) loginModels.JWT {
	return loginModels.JWT{
		PrivateKey: PrivateKey,
		PublicKey:  PublicKey,
	}
}

var tok string
var cont interface{}

func (impl *LoginServiceImpl) AuthenticateRole(cookie *http.Cookie, check string) bool {

	fmt.Println("CONTENT:", cont)
	content, err := loginModels.JwtToken.Validate(cookie.Value)
	if err != nil {
		impl.logger.Errorw("Something wrong with token validation", err)
	}
	var checkType loginModels.UserRoleType
	if check == "principal" {
		checkType = loginModels.Principal
	}
	if check == "teacher" {
		checkType = loginModels.Teacher
	}
	if check == "student" {
		checkType = loginModels.Student
	}
	if impl.loginRepository.CheckRole(content) != checkType {
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
	}
	fmt.Println("good to go")
	return true
}
func (impl *LoginServiceImpl) Login(username, password, userRole string) string {
	userID, err := impl.loginRepository.CheckCreds(username, userRole)
	if err != nil {
		impl.logger.Errorw("Invalid login credentials", err)
		return ""
	}
	if impl.loginRepository.AuthenticateUser(username, password, userRole) {
		prvKey, err := os.ReadFile("private.pem")
		if err != nil {
			log.Fatalln(err)
		}
		pubKey, err := os.ReadFile("public.pem")
		if err != nil {
			log.Fatalln(err)
		}
		loginModels.JwtToken = impl.NewJWT(prvKey, pubKey)
		// 1. Create a new JWT token.
		tok, err = loginModels.JwtToken.Create(userID)
		fmt.Println("user", userID)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("TOKEN:", tok)

		return tok
	}
	impl.logger.Errorw("Invalid login credentials", "error", err)
	return ""
}
