package attendanceModels

// Attendance represents attendance records
type Attendance struct {
	tableName struct{}     `sql:"attendance"`
	ID        int          `json:"id" pg:",pk"`
	UserID    string       `json:"user_id" pg:",fk"`
	Day       int          `json:"day"`
	Month     int          `json:"month"`
	Year      int          `json:"year"`
	PunchMap  []PunchInOut `json:"-" pg:"rel:has-many"`
}

// ClassMappingAttendance represents the mapping between punchinout and class
type ClassMappingAttendance struct {
	tableName struct{} `sql:"classmapattendance"`
	ID        int      `json:"id" pg:",pk"`
	PunchID   int      `json:"punch_id" pg:",fk"`
	ClassID   int      `json:"class_id" pg:",fk"`
	Class     Class    `json:"-" pg:"rel:has-one"`
}

// PunchInOut represents a punch-in or punch-out record
type PunchInOut struct {
	tableName    struct{}                 `sql:"punchinout"`
	ID           int                      `json:"id" pg:",pk"`
	AttendanceID int                      `json:"attendance_id" pg:",fk"`
	UserID       string                   `json:"user_id"`
	PunchIn      string                   `json:"punch_in" `
	PunchOut     string                   `json:"punch_out" `
	ClassMap     []ClassMappingAttendance `json:"-"`
}
type Class struct {
	tableName struct{} `sql:"class"`
	ClassID   int      `json:"class_id" pg:",pk"`
	ClassName string   `json:"class_name"`
}
