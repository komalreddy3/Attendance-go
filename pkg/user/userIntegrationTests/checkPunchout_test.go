package userIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"go.uber.org/zap"
	"testing"
)

func TestCheckPunchOut(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Test the CheckPunchOut function
	userID := "testUserID"
	enrolledClasses, err := repo.CheckPunchOut(userID)

	// Assert the result as per your expectations
	if err != nil {
		t.Errorf("CheckPunchOut failed: %v", err)
	}

	expectedEnrolledClasses := []string{"testClass"} // Adjust as per your expected data
	if !stringSlicesEqual(enrolledClasses, expectedEnrolledClasses) {
		t.Errorf("Expected enrolled classes %v, got %v", expectedEnrolledClasses, enrolledClasses)
	}
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
