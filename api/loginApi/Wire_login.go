package loginApi

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/api/loginApi/loginResthandler"
	"github.com/komalreddy3/Attendance-go/api/loginApi/loginRouter"
)

// var DashboardWire = wire.NewSet(
//
//	dashboardRouter.NewDashboardRouterImpl,
//	wire.Bind(new(dashboardRouter.DashboardRouterInterfaceImpl), new(*dashboardRouter.DashboardRouter)),
//	dashboardResthandler.NewDashboardRestHandler,
//	wire.Bind(new(dashboardResthandler.DashboardHandler), new(*dashboardResthandler.DashboardRestHandler)),
//
// )
var LoginWire = wire.NewSet(
	loginRouter.NewLoginRouterImpl,
	wire.Bind(new(loginRouter.LoginRouterInterfaceImpl), new(*loginRouter.LoginRouter)),
	loginResthandler.NewLoginRestHandler,
	wire.Bind(new(loginResthandler.LoginHandler), new(*loginResthandler.LoginRestHandler)),
)
