package attendanceIntegrationTests

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestFetchClassMapPunch(t *testing.T) {
	// Set up the test database
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer teardownTestDB(testDB)

	// Initialize the repository with the test database
	repo := attendanceRepository.NewAttendanceRepositoryImpl(testDB, zap.NewNop().Sugar())

	// Insert test data into the database
	classID := 1
	punchID := 1
	className := "TestClassName"
	classMappingAttendance := attendanceModels.ClassMappingAttendance{
		ClassID: classID,
		PunchID: punchID,
	}
	_, err = testDB.Model(&classMappingAttendance).Insert()
	if err != nil {
		t.Fatal(err)
	}

	class := attendanceModels.Class{
		ClassID:   classID,
		ClassName: className,
	}
	_, err = testDB.Model(&class).Insert()
	if err != nil {
		t.Fatal(err)
	}

	// Test the FetchClassMapPunch function
	fetchedClassName := repo.FetchClassMapPunch(punchID)

	// Assert the result as per your expectations
	assert.Equal(t, className, fetchedClassName, "Unexpected class name")
}
