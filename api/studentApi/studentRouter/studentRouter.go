package studentRouter

import (
	"github.com/gorilla/mux"
	"github.com/komalreddy3/Attendance-go/api/studentApi/studentResthandler"
)

type StudentRouter struct {
	StudentHandler studentResthandler.StudentHandler
}
type StudentRouterInterfaceImpl interface {
	SetupRoutes(StudentSubRouter *mux.Router)
}

func NewStudentRouterImpl(StudentHandler studentResthandler.StudentHandler) *StudentRouter {
	return &StudentRouter{
		StudentHandler: StudentHandler,
	}
}
func (impl *StudentRouter) SetupRoutes(StudentSubRouter *mux.Router) {
	StudentSubRouter.Path("/punchIn").HandlerFunc(impl.StudentHandler.StudentPunchInHandler).Methods("POST")
	StudentSubRouter.Path("/punchOut").HandlerFunc(impl.StudentHandler.StudentPunchOutHandler).Methods("POST")
	StudentSubRouter.Path("/attendance").HandlerFunc(impl.StudentHandler.GetStudentAttendanceByMonthHandler).Methods("GET")

}
