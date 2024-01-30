package dashboard

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/pkg/dashboard/dashboardRepository"
	"github.com/komalreddy3/Attendance-go/pkg/dashboard/dashboardServices"
)

var Dashboard2Wire = wire.NewSet(
	dashboardRepository.NewDashboardRepositoryImpl,
	wire.Bind(new(dashboardRepository.DashboardRepo), new(*dashboardRepository.DashboardRepository)),
	dashboardServices.NewDashboardServiceImpl,
	wire.Bind(new(dashboardServices.DashboardService), new(*dashboardServices.DashboardServiceImpl)),
)
