package attendanceIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestFetchDay(t *testing.T) {
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
	attendanceID := 1
	attendance := attendanceModels.Attendance{
		ID:    attendanceID,
		Day:   day,
		Month: 1,
		Year:  2022,
	}
	_, err = testDB.Model(&attendance).Insert()
	if err != nil {
		t.Fatal(err)
	}

	// Test the FetchDay function
	fetchedDay := repo.FetchDay(attendanceID)

	// Assert the result as per your expectations
	assert.Equal(t, day, fetchedDay, "Unexpected day")
}
