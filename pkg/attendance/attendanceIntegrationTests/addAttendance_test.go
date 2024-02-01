package attendanceIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestAddAttendance(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := attendanceRepository.NewAttendanceRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Test data
	userID := "testUserID"
	currentDate := time.Now()

	// Test the AddAttendance function
	newAttendanceID, err := repo.AddAttendance(userID, currentDate, 0)

	// Assert the result as per your expectations
	assert.NoError(t, err, "Unexpected error during AddAttendance")

	// Query the database to verify the creation
	var createdAttendanceRecord attendanceModels.Attendance
	err = testDB.Model(&createdAttendanceRecord).
		Where("user_id = ? AND day = ? AND month = ? AND year = ?", userID, currentDate.Day(), int(currentDate.Month()), currentDate.Year()).
		Select()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, createdAttendanceRecord.ID, "Attendance record not created")
	assert.Equal(t, newAttendanceID, createdAttendanceRecord.ID, "Returned attendance ID does not match created record ID")
}
