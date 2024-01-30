package principal

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/pkg/principal/principalRepository"
	"github.com/komalreddy3/Attendance-go/pkg/principal/principalServices"
)

var Principal2Wire = wire.NewSet(
	principalRepository.NewPrincipalRepositoryImpl,
	wire.Bind(new(principalRepository.PrincipalRepo), new(*principalRepository.PrincipalRepository)),
	principalServices.NewPrincipalServiceImpl,
	wire.Bind(new(principalServices.PrincipalService), new(*principalServices.PrincipalServiceImpl)),
)
