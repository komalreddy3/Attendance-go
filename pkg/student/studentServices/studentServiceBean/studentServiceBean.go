package studentServiceBean

type PunchRecord struct {
	ID       int
	UserID   string
	PunchIn  string
	PunchOut string
}
type AttendanceEntry struct {
	Class        string `json:"class"`
	FirstPunchIn string `json:"firstPunchIn"`
	LastPunchOut string `json:"lastPunchOut"`
}
type TeacherAttendanceResponse struct {
	ID         string                    `json:"id"`
	Month      int                       `json:"month"`
	Year       int                       `json:"year"`
	Attendance map[int][]AttendanceEntry `json:"attendance"`
}
