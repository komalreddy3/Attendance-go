package attendanceIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestUpdatePunchOut(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := attendanceRepository.NewAttendanceRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Insert test data into the database
	attendanceID := 1
	currentDate := time.Now()

	punchInOutRecord := attendanceModels.PunchInOut{
		AttendanceID: attendanceID,
		UserID:       "testUserID",
		PunchIn:      currentDate.Add(-1 * time.Hour).Format("15:04:05"),
	}
	_, err = testDB.Model(&punchInOutRecord).Insert()
	if err != nil {
		t.Fatal(err)
	}

	// Test the UpdatePunchOut function
	err = repo.UpdatePunchOut(attendanceID, 0, currentDate)

	// Assert the result as per your expectations
	assert.NoError(t, err, "Unexpected error during UpdatePunchOut")

	// Query the database to verify the update
	var updatedPunchInOutRecord attendanceModels.PunchInOut
	err = testDB.Model(&updatedPunchInOutRecord).
		Where("attendance_id = ? AND punch_out IS NOT NULL", attendanceID).
		Select()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, currentDate.Format("15:04:05"), updatedPunchInOutRecord.PunchOut, "Unexpected PunchOut time")
}
