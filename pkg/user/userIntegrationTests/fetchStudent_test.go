package userIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"go.uber.org/zap"
	"testing"
)

func TestFetchStudent(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Test the FetchStudent function
	userID := "testUserID"
	fetchedUserID := repo.FetchStudent(userID)

	// Assert the result as per your expectations
	if fetchedUserID != userID {
		t.Errorf("Expected user ID %s, got %s", userID, fetchedUserID)
	}
}
