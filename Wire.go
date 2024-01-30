//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/api"
	"github.com/komalreddy3/Attendance-go/api/dashboardApi"
	"github.com/komalreddy3/Attendance-go/api/loginApi"
	"github.com/komalreddy3/Attendance-go/api/principalApi"
	"github.com/komalreddy3/Attendance-go/api/studentApi"
	"github.com/komalreddy3/Attendance-go/api/teacherApi"
	"github.com/komalreddy3/Attendance-go/pkg/attendance"
	"github.com/komalreddy3/Attendance-go/pkg/dashboard"
	"github.com/komalreddy3/Attendance-go/pkg/login"
	"github.com/komalreddy3/Attendance-go/pkg/principal"
	"github.com/komalreddy3/Attendance-go/pkg/student"
	"github.com/komalreddy3/Attendance-go/pkg/teacher"
	"github.com/komalreddy3/Attendance-go/pkg/user"
)

func InitializeEvent() (*App, error) {
	wire.Build(
		NewLogger,
		NewDbConnection,
		NewApp,
		api.NewMuxRouterImpl, // func
		dashboardApi.DashboardWire,
		loginApi.LoginWire,
		principalApi.PrincipalWire,
		studentApi.StudentWire,
		teacherApi.TeacherWire,
		attendance.AttendanceWire,
		dashboard.Dashboard2Wire,
		Login.Login2Wire,
		principal.Principal2Wire,
		student.Student2Wire,
		teacher.Teacher2Wire,
		user.UserWire,
	)
	return &App{}, nil
}
