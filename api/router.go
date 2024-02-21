package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/komalreddy3/Attendance-go/api/dashboardApi/dashboardRouter"
	"github.com/komalreddy3/Attendance-go/api/loginApi/loginRouter"
	"github.com/komalreddy3/Attendance-go/api/principalApi/principalRouter"
	"github.com/komalreddy3/Attendance-go/api/studentApi/studentRouter"
	"github.com/komalreddy3/Attendance-go/api/teacherApi/teacherRouter"
	"github.com/rs/cors"
	"log"
	"net/http"
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
	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"}, // Update with your React frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})

	//handler := c.Handler(r)
	// Wrap the muxRouter.Router with CORS middleware
	handler := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, re *http.Request) {
		r.Router.ServeHTTP(w, re)
	}))
	PrincipalSubRouter := r.Router.PathPrefix("/principal").Subrouter()
	r.PrincipalRouter.InitPrincipalRouter(PrincipalSubRouter)
	TeacherSubRouter := r.Router.PathPrefix("/teacher").Subrouter()
	r.TeacherSubRouter.SetupRoutes(TeacherSubRouter)

	StudentSubRouter := r.Router.PathPrefix("/student").Subrouter()
	r.StudentSubRouter.SetupRoutes(StudentSubRouter)

	DashboardSubRouter := r.Router.PathPrefix("/dashboard").Subrouter()
	r.DashboardSubRouter.SetupRoutes(DashboardSubRouter)

	r.LoginSubRouter.SetupRoutes(r.Router)
	fmt.Println("running")
	log.Fatal(http.ListenAndServe(":9000", handler))
}
