package userIntegrationTests

import (
	"fmt"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"go.uber.org/zap"
	"testing"
)

func TestInsertClassMap(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Insert a class for testing
	className := "testClass"
	classID, err := repo.InsertClass(className)
	if err != nil {
		t.Fatalf("Error inserting class for testing: %v", err)
	}

	// Test the InsertClassMap function
	userID := "testUserID"
	err = repo.InsertClassMap(userID, classID)

	// Assert the result as per your expectations
	if err != nil {
		t.Errorf("InsertClassMap failed: %v", err)
	}
	fmt.Println(userID, classID)
	// Fetch the inserted class mapping from the database and assert its existence
	fetchedClassMapping := &userModels.ClassMappingUser{UserID: userID, ClassID: classID}
	err = testDB.Model(fetchedClassMapping).Select()
	if err != nil {
		t.Errorf("Error fetching class mapping from the database: %v", err)
	}
}
