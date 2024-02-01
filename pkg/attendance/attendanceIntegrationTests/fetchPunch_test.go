package attendanceIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestFetchPunch(t *testing.T) {
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
	userID := "testUserID"
	punchIn := time.Now().Add(-1 * time.Hour).String()
	punchOut := time.Now().String()
	punchInOutRecord := attendanceModels.PunchInOut{
		AttendanceID: attendanceID,
		UserID:       userID,
		PunchIn:      punchIn,
		PunchOut:     punchOut,
	}
	_, err = testDB.Model(&punchInOutRecord).Insert()
	if err != nil {
		t.Fatal(err)
	}

	// Test the FetchPunch function
	fetchedPunchInOutRecords := repo.FetchPunch(attendanceID)

	// Assert the result as per your expectations
	assert.Len(t, fetchedPunchInOutRecords, 1, "Expected 1 punch-in/out record, got %d", len(fetchedPunchInOutRecords))

	// Additional assertions
	assert.Equal(t, userID, fetchedPunchInOutRecords[0].UserID, "Unexpected UserID")
	assert.Equal(t, punchIn, fetchedPunchInOutRecords[0].PunchIn, "Unexpected PunchIn time")
	assert.Equal(t, punchOut, fetchedPunchInOutRecords[0].PunchOut, "Unexpected PunchOut time")
}
