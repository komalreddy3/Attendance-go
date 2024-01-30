package dashboardResthandler

import (
	"errors"
	"github.com/komalreddy3/Attendance-go/pkg/dashboard/dashboardServices"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginServices"
	"go.uber.org/zap"
	"net/http"
)

type DashboardHandler interface {
	Principal(w http.ResponseWriter, r *http.Request)
	Teacher(w http.ResponseWriter, r *http.Request)
	Student(w http.ResponseWriter, r *http.Request)
}
type DashboardRestHandler struct {
	dashboardServices dashboardServices.DashboardService
	loginServices     loginServices.LoginService
	logger            *zap.SugaredLogger
}

func NewDashboardRestHandler(dashboardServices dashboardServices.DashboardService, loginServices loginServices.LoginService, logger *zap.SugaredLogger) *DashboardRestHandler {
	return &DashboardRestHandler{
		dashboardServices: dashboardServices,
		loginServices:     loginServices,
		logger:            logger,
	}
}
func (impl DashboardRestHandler) Principal(w http.ResponseWriter, r *http.Request) {
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
	impl.dashboardServices.Principal()
}
func (impl DashboardRestHandler) Teacher(w http.ResponseWriter, r *http.Request) {
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
	impl.dashboardServices.Teacher()
}
func (impl DashboardRestHandler) Student(w http.ResponseWriter, r *http.Request) {
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
	impl.dashboardServices.Student()
}
