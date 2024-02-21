package principalResthandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginServices"
	"github.com/komalreddy3/Attendance-go/pkg/principal/principalServices"
	"go.uber.org/zap"
	"net/http"
)

type PrincipalHandler interface {
	AddStudentHandler(w http.ResponseWriter, r *http.Request)
	AddTeacherHandler(w http.ResponseWriter, r *http.Request)
	GetTeacherAttendanceHandler(w http.ResponseWriter, r *http.Request)
	PrincipalStudentsHandler(w http.ResponseWriter, r *http.Request)
	PrincipalTeachersHandler(w http.ResponseWriter, r *http.Request)
	ClassInsertHandler(w http.ResponseWriter, r *http.Request)
	PrincipalClassesHandler(w http.ResponseWriter, r *http.Request)
}
type PrincipalRestHandler struct {
	principalServices principalServices.PrincipalService
	loginServices     loginServices.LoginService
	logger            *zap.SugaredLogger
}

func NewPrincipalRestHandler(principalServices principalServices.PrincipalService, loginServices loginServices.LoginService, logger *zap.SugaredLogger) *PrincipalRestHandler {
	return &PrincipalRestHandler{
		principalServices: principalServices,
		loginServices:     loginServices,
		logger:            logger,
	}
}

func (impl PrincipalRestHandler) ClassInsertHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "principal") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var request struct {
		ClassName string `json:"class_name"`
	}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	text := impl.principalServices.ClassInsert(request.ClassName)
	err = json.NewEncoder(w).Encode(text)
	fmt.Println(err)
	if text != 0 {
		//w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}
}
func (impl PrincipalRestHandler) AddStudentHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "principal") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var request struct {
		StudentID string `json:"studentID"`
		ClassName string `json:"class_name"`
	}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	text := impl.principalServices.AddStudent(request.StudentID, request.ClassName)

	if text == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}
}
func (impl PrincipalRestHandler) AddTeacherHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "principal") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var request struct {
		TeacherID  string   `json:"teacherID"`
		ClassNames []string `json:"class_names"`
	}
	//var teacherID string
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		impl.logger.Errorw("Invalid request body", "error", err)
		return
	}
	fmt.Println("In resthandler ", request.ClassNames)
	text := impl.principalServices.AddTeacher(request.TeacherID, request.ClassNames)

	if text == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}
}
func (impl PrincipalRestHandler) GetTeacherAttendanceHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "principal") == false {
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
	err = json.NewEncoder(w).Encode(impl.principalServices.GetTeacherAttendance(request.ID, request.Month, request.Year))
	if err != nil {
		impl.logger.Errorw("Cant produce output for get teacher attendance", err)
	}

}
func (impl PrincipalRestHandler) PrincipalStudentsHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "principal") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)

		w.Write([]byte(`{"success": false}`))
		return
	}
	responseJSON, err := json.Marshal(impl.principalServices.PrincipalStudents())
	if err != nil {
		impl.logger.Errorw("Error encoding response data", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		impl.logger.Errorw("cant produce output properly for PrincipalStudents function", err)
	}
}
func (impl PrincipalRestHandler) PrincipalTeachersHandler(w http.ResponseWriter, r *http.Request) {
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
	if impl.loginServices.AuthenticateRole(cookie, "principal") == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		//w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"success": false}`))
		return
	}
	responseJSON, err := json.Marshal(impl.principalServices.PrincipalTeachers())
	if err != nil {
		impl.logger.Errorw("Error encoding response data", "error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		impl.logger.Errorw("cant produce output properly for PrincipalStudents function", err)
	}
}
func (impl PrincipalRestHandler) PrincipalClassesHandler(w http.ResponseWriter, r *http.Request) {
	//var cookie *http.Cookie
	//cookie, err := r.Cookie("jwt")
	//if err != nil {
	//	switch {
	//	case errors.Is(err, http.ErrNoCookie):
	//		impl.logger.Errorw("cookie not found", err)
	//	default:
	//		impl.logger.Errorw("server error", err)
	//	}
	//	return
	//}
	//if impl.loginServices.AuthenticateRole(cookie, "principal") == false {
	//	http.Redirect(w, r, "/login", http.StatusSeeOther)
	//	//w.WriteHeader(http.StatusMethodNotAllowed)
	//	w.Write([]byte(`{"success": false}`))
	//	return
	//}
	responseJSON, err := json.Marshal(impl.principalServices.PrincipalClasses())
	if err != nil {
		impl.logger.Errorw("Error encoding response data", "error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		impl.logger.Errorw("cant produce output properly for Principal Classes function", err)
	}
}
