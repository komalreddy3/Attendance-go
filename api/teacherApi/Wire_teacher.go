package teacherApi

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/api/teacherApi/teacherResthandler"
	"github.com/komalreddy3/Attendance-go/api/teacherApi/teacherRouter"
)

var TeacherWire = wire.NewSet(
	teacherRouter.NewTeacherRouterImpl,
	wire.Bind(new(teacherRouter.TeacherRouterInterfaceImpl), new(*teacherRouter.TeacherRouter)),
	teacherResthandler.NewTeacherRestHandler,
	wire.Bind(new(teacherResthandler.TeacherHandler), new(*teacherResthandler.TeacherRestHandler)),
)
