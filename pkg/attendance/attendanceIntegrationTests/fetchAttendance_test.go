package attendanceIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestFetchAttendance(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := attendanceRepository.NewAttendanceRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Insert test data into the database
	day := 1
	month := 1
	year := 2022
	userID := "testUserID"
	attendance := attendanceModels.Attendance{
		UserID: userID,
		Day:    day,
		Month:  month,
		Year:   year,
	}
	_, err = testDB.Model(&attendance).Insert()
	if err != nil {
		t.Fatal(err)
	}

	// Test the FetchAttendance function
	studentAttendances := repo.FetchAttendance(day, month, year)

	// Assert the result as per your expectations
	assert.Len(t, studentAttendances, 1, "Expected 1 student attendance, got %d", len(studentAttendances))

	// Additional assertions
	assert.Equal(t, userID, studentAttendances[0].UserID, "Unexpected UserID")
	assert.Equal(t, day, studentAttendances[0].Day, "Unexpected Day")
	assert.Equal(t, month, studentAttendances[0].Month, "Unexpected Month")
	assert.Equal(t, year, studentAttendances[0].Year, "Unexpected Year")
}
