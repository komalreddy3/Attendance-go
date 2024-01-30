package api

import (
	"github.com/gorilla/mux"
	"github.com/komalreddy3/Attendance-go/api/dashboardApi/dashboardRouter"
	"github.com/komalreddy3/Attendance-go/api/loginApi/loginRouter"
	"github.com/komalreddy3/Attendance-go/api/principalApi/principalRouter"
	"github.com/komalreddy3/Attendance-go/api/studentApi/studentRouter"
	"github.com/komalreddy3/Attendance-go/api/teacherApi/teacherRouter"
)

type MuxRouter struct {
	Router             *mux.Router
	PrincipalRouter    principalRouter.PrincipalRouterInterfaceImpl
	TeacherSubRouter   teacherRouter.TeacherRouterInterfaceImpl
	StudentSubRouter   studentRouter.StudentRouterInterfaceImpl
	DashboardSubRouter dashboardRouter.DashboardRouterInterfaceImpl
	LoginSubRouter     loginRouter.LoginRouterInterfaceImpl
}
type MuxRouterInterfaceImpl interface {
	Init()
}

func NewMuxRouterImpl(
	PrincipalRouter principalRouter.PrincipalRouterInterfaceImpl,
	TeacherSubRouter teacherRouter.TeacherRouterInterfaceImpl,
	StudentSubRouter studentRouter.StudentRouterInterfaceImpl,
	DashboardSubRouter dashboardRouter.DashboardRouterInterfaceImpl,
	LoginSubRouter loginRouter.LoginRouterInterfaceImpl,
) *MuxRouter {
	//PrincipalSubRouter *mux.Router, TeacherSubRouter *mux.Router, StudentSubRouter *mux.Router
	return &MuxRouter{
		Router:             mux.NewRouter(),
		PrincipalRouter:    PrincipalRouter,
		TeacherSubRouter:   TeacherSubRouter,
		StudentSubRouter:   StudentSubRouter,
		DashboardSubRouter: DashboardSubRouter,
		LoginSubRouter:     LoginSubRouter,
	}
}

func (r *MuxRouter) Init() {

	PrincipalSubRouter := r.Router.PathPrefix("/principal").Subrouter()
	r.PrincipalRouter.InitPrincipalRouter(PrincipalSubRouter)
	TeacherSubRouter := r.Router.PathPrefix("/teacher").Subrouter()
	r.TeacherSubRouter.SetupRoutes(TeacherSubRouter)

	StudentSubRouter := r.Router.PathPrefix("/student").Subrouter()
	r.StudentSubRouter.SetupRoutes(StudentSubRouter)

	DashboardSubRouter := r.Router.PathPrefix("/dashboard").Subrouter()
	r.DashboardSubRouter.SetupRoutes(DashboardSubRouter)

	r.LoginSubRouter.SetupRoutes(r.Router)
}
