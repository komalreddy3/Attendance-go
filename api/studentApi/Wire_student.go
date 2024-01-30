package studentApi

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/api/studentApi/studentResthandler"
	"github.com/komalreddy3/Attendance-go/api/studentApi/studentRouter"
)

var StudentWire = wire.NewSet(
	studentRouter.NewStudentRouterImpl,
	wire.Bind(new(studentRouter.StudentRouterInterfaceImpl), new(*studentRouter.StudentRouter)),
	studentResthandler.NewStudentRestHandler,
	wire.Bind(new(studentResthandler.StudentHandler), new(*studentResthandler.StudentRestHandler)),
)
