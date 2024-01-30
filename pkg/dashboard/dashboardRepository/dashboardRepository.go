package dashboardRepository

import (
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type DashboardRepository struct {
	dbConnection *pg.DB
	logger       *zap.SugaredLogger
}
type DashboardRepo interface {
}

func NewDashboardRepositoryImpl(dbConnection *pg.DB, logger *zap.SugaredLogger) *DashboardRepository {
	return &DashboardRepository{
		dbConnection: dbConnection,
		logger:       logger,
	}
}
