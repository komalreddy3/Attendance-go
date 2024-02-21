package teacherRouter

import (
	"github.com/gorilla/handlers"
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
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8000"}) // Replace with your React frontend URL
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Apply CORS middleware to the entire router
	TeacherSubRouter.Use(handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders))
	TeacherSubRouter.Path("/punchIn").HandlerFunc(impl.TeacherHandler.TeacherPunchInHandler).Methods("POST")
	TeacherSubRouter.Path("/punchOut").HandlerFunc(impl.TeacherHandler.TeacherPunchOutHandler).Methods("POST")
	TeacherSubRouter.Path("/attendance").HandlerFunc(impl.TeacherHandler.GetTeacherAttendanceByMonthHandler).Methods("GET", "POST")
	TeacherSubRouter.Path("/classAttendance").HandlerFunc(impl.TeacherHandler.GetClassAttendanceHandler).Methods("GET", "POST")

}
