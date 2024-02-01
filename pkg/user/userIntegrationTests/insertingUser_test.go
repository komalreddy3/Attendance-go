package userIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository"
	"github.com/komalreddy3/Attendance-go/pkg/user/userRepository/userModels"
	"go.uber.org/zap"
	"testing"
)

func TestInsertingStudent(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := userRepository.NewUserRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Test the InsertingStudent function
	studentID := "testStudentID"
	username := "testStudentUsername"
	password := "testStudentPassword"

	err = repo.InsertingStudent(studentID, username, password)

	// Assert the result as per your expectations
	if err != nil {
		t.Errorf("InsertingStudent failed: %v", err)
	}

	// Fetch the student from the database and assert its existence
	fetchedStudent := &userModels.User{ID: studentID}
	err = testDB.Model(fetchedStudent).WherePK().Select()
	if err != nil {
		t.Errorf("Error fetching student from the database: %v", err)
	}

	if fetchedStudent.Username != username {
		t.Errorf("Expected student username %s, got %s", username, fetchedStudent.Username)
	}

}
