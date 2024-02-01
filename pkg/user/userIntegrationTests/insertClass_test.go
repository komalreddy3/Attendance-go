package userIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"go.uber.org/zap"
	"testing"
)

func TestInsertClass(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Test the InsertClass function
	className := "testClass"
	classID, err := repo.InsertClass(className)

	// Assert the result as per your expectations
	if err != nil {
		t.Errorf("InsertClass failed: %v", err)
	}

	// Fetch the inserted class from the database and assert its existence
	fetchedClass := &userModels.Class{ClassID: classID}
	err = testDB.Model(fetchedClass).WherePK().Select()
	if err != nil {
		t.Errorf("Error fetching class from the database: %v", err)
	}

	if fetchedClass.ClassName != className {
		t.Errorf("Expected class name %s, got %s", className, fetchedClass.ClassName)
	}
}
