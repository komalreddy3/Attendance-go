package userIntegrationTests

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"go.uber.org/zap"
	"log"
	"os"
	"testing"
)

var testDB *pg.DB

func setupTestDB() (*pg.DB, error) {
	// Set up the test database connection
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		//User:     "postgres",
		//Password: "abc@123",
		//Database: "test",
	})

	// Create necessary tables for testing
	err := db.Model(&userModels.User{}).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Model(&userModels.ClassMappingUser{}).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Model(&userModels.Class{}).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatal(err)
	}

	// Insert test data into the database
	insertTestData(db)

	return db, err
}

func insertTestData(db *pg.DB) {
	// Insert sample data for testing
	//class := userModels.Class{
	//	ClassID:   1,
	//	ClassName: "testClass",
	//}
	//_, err := db.Model(&class).Insert()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//user := userModels.User{
	//	ID:       "testUserID",
	//	Username: "testUsername",
	//	Password: "testPassword",
	//	Role:     "student",
	//}
	//_, err = db.Model(&user).Insert()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//classMapping := userModels.ClassMappingUser{
	//	ID:      1,
	//	UserID:  "testUserID",
	//	ClassID: 1,
	//}
	//_, err = db.Model(&classMapping).Insert()
	//if err != nil {
	//	log.Fatal(err)
	//}
	// Insert sample data for testing
	users := []userModels.User{
		{
			ID:       "student1",
			Username: "student1_username",
			Password: "student1_password",
			Role:     "student",
		},
		{
			ID:       "student2",
			Username: "student2_username",
			Password: "student2_password",
			Role:     "student",
		},
		{
			ID:       "teacher1",
			Username: "teacher1_username",
			Password: "teacher1_password",
			Role:     "teacher",
		},
	}

	_, err := db.Model(&users).Insert()
	if err != nil {
		log.Fatal(err)
	}
}

func teardownTestDB(db *pg.DB) {
	// Drop tables after testing
	err := db.Model(&userModels.User{}).DropTable(&orm.DropTableOptions{IfExists: true, Cascade: true})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Model(&userModels.ClassMappingUser{}).DropTable(&orm.DropTableOptions{IfExists: true, Cascade: true})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Model(&userModels.Class{}).DropTable(&orm.DropTableOptions{IfExists: true, Cascade: true})
	if err != nil {
		log.Fatal(err)
	}
}

func TestCheckEnrollment(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Test the CheckEnrollment function
	userID := "testUserID"
	className := "testClass"
	err = repo.CheckEnrollment(userID, className)

	// Assert the result as per your expectations
	if err != nil {
		t.Errorf("CheckEnrollment failed: %v", err)
	}
}
