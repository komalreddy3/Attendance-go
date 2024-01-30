package teacherRouter

import (
	"github.com/gorilla/mux"
	"github.com/komalreddy3/Attendance-go/api/teacherApi/teacherResthandler"
)

type TeacherRouter struct {
	TeacherHandler teacherResthandler.TeacherHandler
}

type TeacherRouterInterfaceImpl interface {
	SetupRoutes(TeacherSubRouter *mux.Router)
}

func NewTeacherRouterImpl(TeacherHandler teacherResthandler.TeacherHandler) *TeacherRouter {
	return &TeacherRouter{
		TeacherHandler: TeacherHandler,
	}
}
func (impl *TeacherRouter) SetupRoutes(TeacherSubRouter *mux.Router) {
	TeacherSubRouter.Path("/punchIn").HandlerFunc(impl.TeacherHandler.TeacherPunchInHandler).Methods("POST")
	TeacherSubRouter.Path("/punchOut").HandlerFunc(impl.TeacherHandler.TeacherPunchOutHandler).Methods("POST")
	TeacherSubRouter.Path("/attendance").HandlerFunc(impl.TeacherHandler.GetTeacherAttendanceByMonthHandler).Methods("GET")
	TeacherSubRouter.Path("/classAttendance").HandlerFunc(impl.TeacherHandler.GetClassAttendanceHandler).Methods("GET")

}
