package attendanceIntegrationTests

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"go.uber.org/zap"
	"log"
	"testing"
	"time"
)

var testDB *pg.DB

func setupTestDB() (*pg.DB, error) {
	// Set up the test database connection
	db := pg.Connect(&pg.Options{
		Addr: "localhost:5432",
		//User:     os.Getenv("DB_USER"),
		//Password: os.Getenv("DB_PASSWORD"),
		//Database: os.Getenv("DB_NAME"),
		User:     "postgres",
		Password: "abc@123",
		Database: "test",
	})

	// Create necessary tables for testing
	err := db.Model(&attendanceModels.Attendance{}).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatal(err)
	}
	// Create necessary tables for testing
	err = db.Model(&attendanceModels.ClassMappingAttendance{}, &attendanceModels.Class{}).
		CreateTable(&orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Model(&attendanceModels.Class{}, &attendanceModels.Class{}).
		CreateTable(&orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Model(&attendanceModels.PunchInOut{}).
		CreateTable(&orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatal(err)
	}
	return db, err

}

func teardownTestDB(db *pg.DB) {
	// Drop tables after testing
	err := db.Model(&attendanceModels.Attendance{}).DropTable(&orm.DropTableOptions{IfExists: true, Cascade: true})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Model(&attendanceModels.Class{}).DropTable(&orm.DropTableOptions{IfExists: true, Cascade: true})
	if err != nil {
		log.Fatal(err)
	}
	// Drop tables after testing
	err = db.Model(&attendanceModels.ClassMappingAttendance{}, &attendanceModels.Class{}).
		DropTable(&orm.DropTableOptions{IfExists: true, Cascade: true})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Model(&attendanceModels.PunchInOut{}).
		DropTable(&orm.DropTableOptions{IfExists: true, Cascade: true})
	if err != nil {
		log.Fatal(err)
	}
}

func TestHasAttendance(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := attendanceRepository.NewAttendanceRepositoryImpl(testDB, zap.NewNop().Sugar())
	// Insert test data into the database
	userID := "testUserID"
	currentDate := time.Now()
	attendance := attendanceModels.Attendance{
		UserID: userID,
		Day:    currentDate.Day(),
		Month:  int(currentDate.Month()),
		Year:   currentDate.Year(),
	}
	_, err = testDB.Model(&attendance).Insert()
	if err != nil {
		t.Fatal(err)
	}

	// Test the HasAttendance function
	existingID, err := repo.HasAttendance(userID)

	// Assert the result as per your expectations
	if err != nil {
		t.Errorf("HasAttendance failed: %v", err)
	}

	// Optionally, you can assert the existingID or other expectations
	if existingID == 0 {
		t.Errorf("Expected existing attendance ID, got 0")
	}
}
