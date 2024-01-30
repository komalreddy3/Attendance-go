package principalApi

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/api/principalApi/principalResthandler"
	"github.com/komalreddy3/Attendance-go/api/principalApi/principalRouter"
)

var PrincipalWire = wire.NewSet(

	principalRouter.NewPrincipalRouterImpl,
	wire.Bind(new(principalRouter.PrincipalRouterInterfaceImpl), new(*principalRouter.PrincipalRouter)),
	principalResthandler.NewPrincipalRestHandler,
	wire.Bind(new(principalResthandler.PrincipalHandler), new(*principalResthandler.PrincipalRestHandler)),
)
