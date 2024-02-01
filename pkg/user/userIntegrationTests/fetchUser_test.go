package userIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"go.uber.org/zap"
	"testing"
)

func TestFetchUser(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Test the FetchUser function for students
	students := repo.FetchUser("student")

	// Assert the result as per your expectations
	if len(students) != 2 {
		t.Errorf("Expected 2 students, got %d", len(students))
	}

	// Test the FetchUser function for teachers
	teachers := repo.FetchUser("teacher")

	// Assert the result as per your expectations
	if len(teachers) != 1 {
		t.Errorf("Expected 1 teacher, got %d", len(teachers))
	}
}
