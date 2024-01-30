package dashboardApi

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/api/dashboardApi/dashboardResthandler"
	"github.com/komalreddy3/Attendance-go/api/dashboardApi/dashboardRouter"
)

var DashboardWire = wire.NewSet(
	dashboardRouter.NewDashboardRouterImpl,
	wire.Bind(new(dashboardRouter.DashboardRouterInterfaceImpl), new(*dashboardRouter.DashboardRouter)),
	dashboardResthandler.NewDashboardRestHandler,
	wire.Bind(new(dashboardResthandler.DashboardHandler), new(*dashboardResthandler.DashboardRestHandler)),
)
