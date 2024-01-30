package attendance

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices"
)

var AttendanceWire = wire.NewSet(

	attendanceRepository.NewAttendanceRepositoryImpl,
	wire.Bind(new(attendanceRepository.AttendanceRepo), new(*attendanceRepository.AttendanceRepository)),
	attendanceServices.NewAttendanceServiceImpl,
	wire.Bind(new(attendanceServices.AttendanceService), new(*attendanceServices.AttendanceServiceImpl)),
)
