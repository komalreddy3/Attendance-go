package userModels

// User represents a generic user
type UserRoleType string

const (
	Principal UserRoleType = "principal"
	Teacher   UserRoleType = "teacher"
	Student   UserRoleType = "student"
)

type User struct {
	tableName struct{}           `sql:"user"`
	ID        string             `json:"id" pg:",pk"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Role      UserRoleType       `json:"role"` // "principal", "teacher", or "student"
	ClassMap  []ClassMappingUser `json:"-" pg:"rel:has-many"`
}

// ClassMappingUser represents the mapping between user and class
type ClassMappingUser struct {
	tableName struct{} `sql:"classmapuser"`
	ID        int      `json:"id" pg:",pk"`
	UserID    string   `json:"user_id" pg:",fk"`
	ClassID   int      `json:"class_id" pg:",fk"`
	Class     Class    `json:"-" pg:"rel:has-one"`
}
type Class struct {
	tableName struct{} `sql:"class"`
	ClassID   int      `json:"class_id" pg:",pk"`
	ClassName string   `json:"class_name"`
}
