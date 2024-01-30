package student

import (
	"github.com/google/wire"
	"github.com/komalreddy3/Attendance-go/pkg/student/studentRepository"
	"github.com/komalreddy3/Attendance-go/pkg/student/studentServices"
)

var Student2Wire = wire.NewSet(
	studentRepository.NewStudentRepositoryImpl,
	wire.Bind(new(studentRepository.StudentRepo), new(*studentRepository.StudentRepository)),
	studentServices.NewStudentServiceImpl,
	wire.Bind(new(studentServices.StudentService), new(*studentServices.StudentServiceImpl)),
)
