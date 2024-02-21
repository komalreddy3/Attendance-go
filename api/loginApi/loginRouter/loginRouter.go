package loginRouter

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/komalreddy3/Attendance-go/api/loginApi/loginResthandler"
)

type LoginRouter struct {
	LoginHandler loginResthandler.LoginHandler
}
type LoginRouterInterfaceImpl interface {
	SetupRoutes(LoginRouter *mux.Router)
}

func NewLoginRouterImpl(LoginHandler loginResthandler.LoginHandler) *LoginRouter {
	return &LoginRouter{
		LoginHandler: LoginHandler,
	}
}
func (impl *LoginRouter) SetupRoutes(LoginRouter *mux.Router) {
	// Define allowed origins, methods, and headers for CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8000"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	LoginRouter.Use(handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders))
	LoginRouter.HandleFunc("/login", impl.LoginHandler.Login).Methods("GET", "POST")
	LoginRouter.HandleFunc("/logout", impl.LoginHandler.Logout).Methods("GET", "POST")
}
