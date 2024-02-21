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
	RemovePunchinHandler(w http.ResponseWriter, r *http.Request)
}

func NewStudentRestHandler(studentServices studentServices.StudentService, loginServices loginServices.LoginService, logger *zap.SugaredLogger) *StudentRestHandler {
	return &StudentRestHandler{
		studentServices: studentServices,
		loginServices:   loginServices,
		logger:          logger,
	}
}
func (impl StudentRestHandler) RemovePunchinHandler(w http.ResponseWriter, r *http.Request) {
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
	//userID := request.UserID
	userID := impl.loginServices.GetId(cookie)
	class := request.Class
	err = impl.studentServices.RemovePunchIn(userID, class)
	if err == nil {
		w.Write([]byte(`{"allow": false}`))
	} else {
		w.Write([]byte(`{"allow": true}`))
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
	//userID := request.UserID
	userID := impl.loginServices.GetId(cookie)
	class := request.Class
	//err = impl.studentServices.RemovePunchIn(userID, class)
	//if err == nil {
	//	w.Write([]byte(`{"allow": false}`))
	//}
	err = impl.studentServices.StudentPunchIn(userID, class)
	if err != nil {
		// Respond with failure
		//w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false}`))
	}
	// Respond with success
	//w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true}`))
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

	//userID := request.UserID
	userID := impl.loginServices.GetId(cookie)
	class := request.Class
	err = impl.studentServices.StudentPunchOut(userID, class)
	if err != nil {
		// Respond with failure
		//w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false}`))
	}
	// Respond with success
	//w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true}`))
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
		w.Write([]byte(`{"success": false}`))
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
	userID := impl.loginServices.GetId(cookie)
	//err = json.NewEncoder(w).Encode(impl.studentServices.GetStudentAttendanceByMonth(request.ID, request.Month, request.Year))
	err = json.NewEncoder(w).Encode(impl.studentServices.GetStudentAttendanceByMonth(userID, request.Month, request.Year))
	//err = json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "attendanceData": impl.studentServices.GetStudentAttendanceByMonth(request.ID, request.Month, request.Year)})
	if err != nil {
		impl.logger.Errorw("cant produce output properly for GetStudentAttendanceByMonth", err)
	}
	// Respond with success
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte(`{"success": true}`))
}
