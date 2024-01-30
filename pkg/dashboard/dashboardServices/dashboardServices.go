package dashboardServices

import (
	"go.uber.org/zap"
)

type DashboardServiceImpl struct {
	logger *zap.SugaredLogger
}
type DashboardService interface {
	Student()
	Principal()
	Teacher()
}

func NewDashboardServiceImpl(logger *zap.SugaredLogger) *DashboardServiceImpl {
	return &DashboardServiceImpl{
		logger,
	}
}
func (impl *DashboardServiceImpl) Student() {
	//server.AuthenticateRole(w, r, "student")

}
func (impl *DashboardServiceImpl) Principal() {
	//server.AuthenticateRole(w, r, "principal")

}
func (impl *DashboardServiceImpl) Teacher() {
	//server.AuthenticateRole(w, r, "teacher")

}
