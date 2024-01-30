package teacherResthandler

import (
	"encoding/json"
	"errors"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginServices"
	"github.com/komalreddy3/Attendance-go/pkg/teacher/teacherServices"
	"go.uber.org/zap"
	"net/http"
)

type TeacherRestHandler struct {
	teacherServices teacherServices.TeacherService
	loginServices   loginServices.LoginService
	logger          *zap.SugaredLogger
}
type TeacherHandler interface {
	TeacherPunchInHandler(w http.ResponseWriter, r *http.Request)
	TeacherPunchOutHandler(w http.ResponseWriter, r *http.Request)
	GetTeacherAttendanceByMonthHandler(w http.ResponseWriter, r *http.Request)
	GetClassAttendanceHandler(w http.ResponseWriter, r *http.Request)
}

func NewTeacherRestHandler(teacherServices teacherServices.TeacherService, loginServices loginServices.LoginService, logger *zap.SugaredLogger) *TeacherRestHandler {
	return &TeacherRestHandler{
		loginServices:   loginServices,
		teacherServices: teacherServices,
		logger:          logger,
	}
}

func (impl TeacherRestHandler) TeacherPunchInHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "teacher") == false {
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
	impl.teacherServices.TeacherPunchIn(userID, class)
}
func (impl TeacherRestHandler) TeacherPunchOutHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "teacher") == false {
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
	impl.teacherServices.TeacherPunchOut(userID, class)
}
func (impl TeacherRestHandler) GetTeacherAttendanceByMonthHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "teacher") == false {
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
	err = json.NewEncoder(w).Encode(impl.teacherServices.GetTeacherAttendanceByMonth(request.ID, request.Month, request.Year))
	if err != nil {
		impl.logger.Errorw("cant produce output properly for GetStudentAttendanceByMonth", err)
	}

}
func (impl TeacherRestHandler) GetClassAttendanceHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "teacher") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var request struct {
		Class string `json:"class"`
		Day   int    `json:"day"`
		Month int    `json:"month"`
		Year  int    `json:"year"`
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		impl.logger.Errorw("Invalid request body", "error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(impl.teacherServices.GetClassAttendance(request.Class, request.Day, request.Month, request.Year))
	if err != nil {
		impl.logger.Errorw("cant produce output properly for GetClassAttendance", err)
	}

}
