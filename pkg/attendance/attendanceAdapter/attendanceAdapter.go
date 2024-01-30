package attendanceAdapter

import (
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceRepository/attendanceModels"
	"github.com/komalreddy3/Attendance-go/pkg/attendance/attendanceServices/attendanceServiceBean"
)

func StudentRec(studentAttendances []attendanceModels.Attendance) []int {
	var ids []int
	// Extract 'ID' values into the 'ids' array
	for _, attendance := range studentAttendances {
		ids = append(ids, attendance.ID)
	}
	return ids
}
func PunchRec(punchInOutRecords []attendanceModels.PunchInOut) []attendanceServiceBean.PunchRecord {
	// Create a slice to store the result
	var punchRecords []attendanceServiceBean.PunchRecord
	// Populate punchRecords with data from punchInOutRecords
	for _, record := range punchInOutRecords {
		punchRecord := attendanceServiceBean.PunchRecord{
			ID:       record.ID,
			UserID:   record.UserID,
			PunchIn:  record.PunchIn,
			PunchOut: record.PunchOut,
		}
		punchRecords = append(punchRecords, punchRecord)
	}

	return punchRecords
}
