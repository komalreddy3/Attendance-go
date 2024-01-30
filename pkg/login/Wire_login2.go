package Login

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginRepository"
	"github.com/komalreddy3/Attendance-go/pkg/login/loginServices"
)

var Login2Wire = wire.NewSet(
	loginRepository.NewLoginRepositoryImpl,
	wire.Bind(new(loginRepository.LoginRepo), new(*loginRepository.LoginRepository)),
	loginServices.NewLoginServiceImpl,
	wire.Bind(new(loginServices.LoginService), new(*loginServices.LoginServiceImpl)),
)
