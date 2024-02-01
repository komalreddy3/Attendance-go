package loginResthandler

import (
	"github.com/komalreddy3/Attendance-go/pkg/login/loginServices"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"time"
)

type LoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}
type LoginRestHandler struct {
	loginServices loginServices.LoginService
	logger        *zap.SugaredLogger
}

func NewLoginRestHandler(loginServices loginServices.LoginService, logger *zap.SugaredLogger) *LoginRestHandler {
	return &LoginRestHandler{
		loginServices: loginServices,
		logger:        logger,
	}
}
func (impl LoginRestHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var userRole string

		switch r.FormValue("role") {
		case "principal":
			userRole = "principal"
		case "teacher":
			userRole = "teacher"
		case "student":
			userRole = "student"
		default:
			http.Error(w, "Invalid role", http.StatusBadRequest)
			return
		}
		if impl.loginServices.Login(username, password, userRole) != "" {
			cookie := http.Cookie{Name: "jwt", Value: impl.loginServices.Login(username, password, userRole), Expires: time.Now().Add(time.Minute * 10), HttpOnly: true}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/dashboard/"+userRole, http.StatusSeeOther)
		}
		impl.logger.Errorw("Invalid login credentials", "error")
		return

	}
	// Render the login form
	tmpl, _ := template.New("login").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Login</title>
		</head>
		<body>
			<h1>Login Page</h1>
			<form method="post" action="/login">
				<label for="username">Username:</label>
				<input type="text" id="username" name="username" required><br>
				<label for="password">Password:</label>
				<input type="password" id="password" name="password" required><br>
				<label for="role">Role:</label>
				<input type="text" id="role" name="role" placeholder="principal..teacher..student" required><br>
				<button type="submit">Login</button>
			</form>
		</body>
		</html>
	`)
	err := tmpl.Execute(w, nil)
	if err != nil {
		impl.logger.Errorw("Execution of login went wrong", err)
	}

}
func (impl LoginRestHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "jwt", Value: "", Expires: time.Now().Add(-time.Minute * 10), HttpOnly: true}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
