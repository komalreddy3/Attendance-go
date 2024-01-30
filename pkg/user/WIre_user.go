package user

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userServices"
)

// var AttendanceWire = wire.NewSet(
//
//	attendanceRepository.NewAttendanceRepositoryImpl,
//	wire.Bind(new(attendanceRepository.AttendanceRepo), new(attendanceRepository.AttendanceRepository)),
//	attendanceServices.NewAttendanceServiceImpl,
//	wire.Bind(new(attendanceServices.AttendanceService), new(attendanceServices.AttendanceServiceImpl)),
//
// )
var UserWire = wire.NewSet(
	userRepository.NewUserRepositoryImpl,
	wire.Bind(new(userRepository.UserRepo), new(*userRepository.UserRepository)),
	userServices.NewUserServiceImpl,
	wire.Bind(new(userServices.UserService), new(*userServices.UserServiceImpl)),
)
