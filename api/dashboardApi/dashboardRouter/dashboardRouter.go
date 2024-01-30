package dashboardRouter

import (
	"github.com/gorilla/mux"
	"github.com/komalreddy3/Attendance-go/api/dashboardApi/dashboardResthandler"
)

type DashboardRouter struct {
	DashboardHandler dashboardResthandler.DashboardHandler
}
type DashboardRouterInterfaceImpl interface {
	SetupRoutes(DashboardSubRouter *mux.Router)
}

func NewDashboardRouterImpl(DashboardHandler dashboardResthandler.DashboardHandler) *DashboardRouter {
	return &DashboardRouter{
		DashboardHandler: DashboardHandler,
	}
}

func (impl *DashboardRouter) SetupRoutes(DashboardSubRouter *mux.Router) {
	DashboardSubRouter.Path("/principal").HandlerFunc(impl.DashboardHandler.Principal).Methods("GET")
	DashboardSubRouter.Path("/teacher").HandlerFunc(impl.DashboardHandler.Teacher).Methods("GET")
	DashboardSubRouter.Path("/student").HandlerFunc(impl.DashboardHandler.Student).Methods("GET")

}
