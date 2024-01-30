package loginRouter

import (
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
	LoginRouter.HandleFunc("/login", impl.LoginHandler.Login).Methods("GET", "POST")
	LoginRouter.HandleFunc("/logout", impl.LoginHandler.Logout).Methods("GET")
}
