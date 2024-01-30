package userModels

// User represents a generic user
type UserRoleType string

const (
	Principal UserRoleType = "principal"
	Teacher   UserRoleType = "teacher"
	Student   UserRoleType = "student"
)

type User struct {
	ID       string             `json:"id" pg:",pk"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Role     UserRoleType       `json:"role"` // "principal", "teacher", or "student"
	ClassMap []ClassMappingUser `json:"-" pg:"rel:has-many"`
}

// ClassMappingUser represents the mapping between user and class
type ClassMappingUser struct {
	ID      int    `json:"id" pg:",pk"`
	UserID  string `json:"user_id" pg:",fk"`
	ClassID int    `json:"class_id" pg:",fk"`
	Class   Class  `json:"-" pg:"rel:has-one"`
}
type Class struct {
	ClassID   int    `json:"class_id" pg:",pk"`
	ClassName string `json:"class_name"`
}
