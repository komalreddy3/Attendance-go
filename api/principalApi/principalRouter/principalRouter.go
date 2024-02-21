package principalRouter

import (
	"github.com/gorilla/mux"
	"github.com/komalreddy3/Attendance-go/api/principalApi/principalResthandler"
)

type PrincipalRouter struct {
	PrincipalHandler principalResthandler.PrincipalHandler
}
type PrincipalRouterInterfaceImpl interface {
	InitPrincipalRouter(PrincipalSubRouter *mux.Router)
}

func NewPrincipalRouterImpl(PrincipalHandler principalResthandler.PrincipalHandler) *PrincipalRouter {
	return &PrincipalRouter{
		PrincipalHandler: PrincipalHandler,
	}
}

func (impl *PrincipalRouter) InitPrincipalRouter(PrincipalSubRouter *mux.Router) {
	PrincipalSubRouter.Path("/addStudent").HandlerFunc(impl.PrincipalHandler.AddStudentHandler).Methods("POST")
	PrincipalSubRouter.Path("/addTeacher").HandlerFunc(impl.PrincipalHandler.AddTeacherHandler).Methods("POST")
	PrincipalSubRouter.Path("/teacherAttendance").HandlerFunc(impl.PrincipalHandler.GetTeacherAttendanceHandler).Methods("GET", "POST")
	PrincipalSubRouter.Path("/Students").HandlerFunc(impl.PrincipalHandler.PrincipalStudentsHandler).Methods("GET")
	PrincipalSubRouter.Path("/Teachers").HandlerFunc(impl.PrincipalHandler.PrincipalTeachersHandler).Methods("GET")
	PrincipalSubRouter.Path("/Class").HandlerFunc(impl.PrincipalHandler.ClassInsertHandler).Methods("POST")
	PrincipalSubRouter.Path("/Classes").HandlerFunc(impl.PrincipalHandler.PrincipalClassesHandler).Methods("GET")
}
