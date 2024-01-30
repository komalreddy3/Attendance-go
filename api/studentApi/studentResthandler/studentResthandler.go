package studentResthandler

import (
	"encoding/json"
	"errors"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginServices"
	"github.com/komalreddy3/Attendance-go/pkg/student/studentServices"
	"go.uber.org/zap"
	"net/http"
)

type StudentRestHandler struct {
	studentServices studentServices.StudentService
	loginServices   loginServices.LoginService
	logger          *zap.SugaredLogger
}
type StudentHandler interface {
	StudentPunchInHandler(w http.ResponseWriter, r *http.Request)
	StudentPunchOutHandler(w http.ResponseWriter, r *http.Request)
	GetStudentAttendanceByMonthHandler(w http.ResponseWriter, r *http.Request)
}

func NewStudentRestHandler(studentServices studentServices.StudentService, loginServices loginServices.LoginService, logger *zap.SugaredLogger) *StudentRestHandler {
	return &StudentRestHandler{
		studentServices: studentServices,
		loginServices:   loginServices,
		logger:          logger,
	}
}
func (impl StudentRestHandler) StudentPunchInHandler(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			impl.logger.Errorw("cookie not found", err)
		default:
			impl.logger.Errorw("server error", err)
		}
		return
	}
	if impl.loginServices.AuthenticateRole(cookie, "student") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var request struct {
		UserID string `json:"user_id"`
		Class  string `json:"class"`
	}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		impl.logger.Errorw("Invalid request body", "error", err)
		return
	}
	userID := request.UserID
	class := request.Class
	impl.studentServices.StudentPunchIn(userID, class)
}
func (impl StudentRestHandler) StudentPunchOutHandler(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			impl.logger.Errorw("cookie not found", err)
		default:
			impl.logger.Errorw("server error", err)
		}
		return
	}
	if impl.loginServices.AuthenticateRole(cookie, "student") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var request struct {
		UserID string `json:"user_id"`
		Class  string `json:"class"`
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		impl.logger.Errorw("Invalid request body", "error", err)
		return
	}

	userID := request.UserID
	class := request.Class
	impl.studentServices.StudentPunchOut(userID, class)
}
func (impl StudentRestHandler) GetStudentAttendanceByMonthHandler(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			impl.logger.Errorw("cookie not found", err)
		default:
			impl.logger.Errorw("server error", err)
		}
		return
	}
	if impl.loginServices.AuthenticateRole(cookie, "student") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var request struct {
		ID    string `json:"id"`
		Month int    `json:"month"`
		Year  int    `json:"year"`
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		impl.logger.Errorw("Invalid request body", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(impl.studentServices.GetStudentAttendanceByMonth(request.ID, request.Month, request.Year))
	if err != nil {
		impl.logger.Errorw("cant produce output properly for GetStudentAttendanceByMonth", err)
	}
}
