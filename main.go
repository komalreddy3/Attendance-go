package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/joho/godotenv"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"log"
	"os"
)

type db struct {
	db *pg.DB
}

var DB *pg.DB

func NewDbConnection() *pg.DB {
	//DB = NewDbConnection()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB = pg.Connect(&pg.Options{
		//Addr: "localhost:5432",
		Addr: os.Getenv("DB_ADDR") + ":5432",
		//User:     os.Getenv("DB_USER"),
		//Password: os.Getenv("DB_PASSWORD"),
		//Database: os.Getenv("DB_NAME"),
		User:     "postgres",
		Password: "abc@123",
		Database: "management",
	})
	fmt.Println(os.Getenv("DB_ADDR") + ":5432")
	err = DB.Model(&attendanceModels.Attendance{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Model(&attendanceModels.PunchInOut{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Model(&attendanceModels.ClassMappingAttendance{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Model(&userModels.ClassMappingUser{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Model(&userModels.User{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Model(&userModels.Class{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	// Check if the principal user exists

	err = DB.Model(&userModels.User{}).Where("role = ?", userModels.Principal).Select()
	var set bool = true
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Not set principal")
		set = false
	}

	// If the principal user does not exist, insert it into the database
	if !set {
		_, err := DB.Model(&userModels.User{ID: "user1", Username: "user1", Password: "user1", Role: userModels.Principal}).Insert()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Principal user inserted successfully.")
	} else {
		fmt.Println("Principal user already exists.")
	}
	// Perform a SELECT query to fetch rows from your table
	var rows []userModels.User
	if err := DB.Model(&userModels.User{}).Where("username = ?", "user1").Select(&rows); err != nil {
		fmt.Println("Error querying database")
	}
	fmt.Println("checked")
	// Print the fetched rows
	for _, row := range rows {
		fmt.Printf("ID: %s, Name: %s\n Password:%s\n", row.ID, row.Username, row.Password)
	}
	return DB
}
func main() {
	// Create a new MuxRouter instance

	app, err := InitializeEvent()
	if err != nil {
		log.Fatal(err)
	}
	app.Start()
	//router := api.NewMuxRouterImpl(mux.NewRouter())
	//router.Init()
	//logger, err := zap.NewProduction()
	//if err != nil {
	//	panic(err)
	//}
	//defer logger.Sync()
	//DB := NewDbConnection()
	//userRepo := userRepository.NewUserRepositoryImpl(DB, logger.Sugar())
	//loginRepo := loginRepository.NewLoginRepositoryImpl(DB, logger.Sugar())
	//attendanceRepo := attendanceRepository.NewAttendanceRepositoryImpl(DB, logger.Sugar())
	//
	//servicesUser := userServices.NewUserServiceImpl(*userRepo, logger.Sugar())
	//servicesLogin := loginServices.NewLoginServiceImpl(*loginRepo, logger.Sugar())
	//servicesAttendance := attendanceServices.NewAttendanceServiceImpl(*attendanceRepo, logger.Sugar())
	//teacherRepo := teacherRepository.NewTeacherRepositoryImpl(DB, *servicesUser, *servicesAttendance, logger.Sugar())
	//studentRepo := studentRepository.NewStudentRepositoryImpl(DB, *servicesUser, *servicesAttendance, logger.Sugar())
	//principalRepo := principalRepository.NewPrincipalRepositoryImpl(DB, *servicesUser, *servicesAttendance, logger.Sugar())
	//
	//servicesStudent := studentServices.NewStudentServiceImpl(*studentRepo, logger.Sugar())
	//servicesPrincipal := principalServices.NewPrincipalServiceImpl(*principalRepo, logger.Sugar())
	//servicesTeacher := teacherServices.NewTeacherServiceImpl(*teacherRepo, logger.Sugar())
	//servicesDashboard := dashboardServices.NewDashboardServiceImpl(logger.Sugar())
	//
	//sResthandler := studentResthandler.NewStudentRestHandler(*servicesStudent, *servicesLogin, logger.Sugar())
	//tResthandler := teacherResthandler.NewTeacherRestHandler(*servicesTeacher, *servicesLogin, logger.Sugar())
	//pResthandler := principalResthandler.NewPrincipalRestHandler(*servicesPrincipal, *servicesLogin, logger.Sugar())
	//dResthandler := dashboardResthandler.NewDashboardRestHandler(*servicesDashboard, *servicesLogin, logger.Sugar())
	//lResthandler := loginResthandler.NewLoginRestHandler(*servicesLogin, logger.Sugar())
	//
	//sRouter := studentRouter.NewStudentRouterImpl(router, *sResthandler)
	//tRouter := teacherRouter.NewTeacherRouterImpl(router, *tResthandler)
	//pRouter := principalRouter.NewPrincipalRouterImpl(router, *pResthandler)
	//dRouter := dashboardRouter.NewDashboardRouterImpl(router, *dResthandler)
	//lRouter := loginRouter.NewLoginRouterImpl(router, *lResthandler)
	//
	//sRouter.SetupRoutes()
	//tRouter.SetupRoutes()
	//pRouter.SetupRoutes()
	//dRouter.SetupRoutes()
	//lRouter.SetupRoutes()
	//fmt.Println("running")
	//log.Fatal(http.ListenAndServe("localhost:8080", router.Router))

}
