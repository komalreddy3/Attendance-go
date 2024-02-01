package attendanceIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestPunchOutCheck(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := attendanceRepository.NewAttendanceRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Insert test data into the database
	userID := "testUserID"
	classID := 1

	// Insert a punch-in record for testing
	var punchInOutRecord attendanceModels.PunchInOut
	_, err = testDB.Model(&attendanceModels.PunchInOut{
		UserID:  userID,
		PunchIn: time.Now().Add(-1 * time.Hour).Format("15:04:05"),
	}).Returning("*").Insert(&punchInOutRecord)
	if err != nil {
		t.Fatal(err)
	}

	// Insert a class mapping record for testing
	_, err = testDB.Model(&attendanceModels.ClassMappingAttendance{
		PunchID: punchInOutRecord.ID,
		ClassID: classID,
	}).Insert()
	if err != nil {
		t.Fatal(err)
	}

	// Test the PunchOutCheck function
	err = repo.PunchOutCheck(userID, classID)

	// Assert the result as per your expectations
	assert.NoError(t, err, "Unexpected error during PunchOutCheck")

	// Query the database to verify the punch-out check
	var updatedPunchInOutRecord attendanceModels.PunchInOut
	err = testDB.Model(&updatedPunchInOutRecord).
		Where("id = ?", punchInOutRecord.ID).
		Select()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, updatedPunchInOutRecord.PunchOut, "PunchOut time not updated")
}
