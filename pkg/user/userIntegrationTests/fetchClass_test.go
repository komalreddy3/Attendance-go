package userIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"go.uber.org/zap"
	"testing"
)

func TestFetchClass(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Test the FetchClass function
	enrolledClassName := "testClass"
	classID := repo.FetchClass(enrolledClassName)

	// Assert the result as per your expectations
	if classID != 1 { // Adjust as per your expected data
		t.Errorf("Expected class ID %d, got %d", 1, classID)
	}
}
