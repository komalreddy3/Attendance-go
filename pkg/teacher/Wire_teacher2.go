package teacher

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/pkg/teacher/teacherRepository"
	"github.com/komalreddy3/Attendance-go/pkg/teacher/teacherServices"
)

var Teacher2Wire = wire.NewSet(
	teacherRepository.NewTeacherRepositoryImpl,
	wire.Bind(new(teacherRepository.TeacherRepo), new(*teacherRepository.TeacherRepository)),
	teacherServices.NewTeacherServiceImpl,
	wire.Bind(new(teacherServices.TeacherService), new(*teacherServices.TeacherServiceImpl)),
)
