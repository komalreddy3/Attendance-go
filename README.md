
# Attendance Management System
**_Using golang (Backend)[Wire for dependency injection] ReactJS (Frontend) Postgresql (Database) Zap Sugared Logger (Logging) JWT token (Authentication and Authorisation)_**

## Project Structure 

### Backend 

```
.
|-- App.go
|-- Dockerfile
|-- README.md
|-- Wire.go
|-- api
|   |-- dashboardApi
|   |   |-- Wire_dashboard.go
|   |   |-- dashboardResthandler
|   |   |   `-- dashboardResthandler.go
|   |   `-- dashboardRouter
|   |       `-- dashboardRouter.go
|   |-- loginApi
|   |   |-- Wire_login.go
|   |   |-- loginResthandler
|   |   |   `-- loginResthandler.go
|   |   `-- loginRouter
|   |       `-- loginRouter.go
|   |-- principalApi
|   |   |-- Wire_principal.go
|   |   |-- principalResthandler
|   |   |   `-- principalResthandler.go
|   |   `-- principalRouter
|   |       `-- principalRouter.go
|   |-- router.go
|   |-- studentApi
|   |   |-- Wire_student.go
|   |   |-- studentResthandler
|   |   |   `-- studentResthandler.go
|   |   `-- studentRouter
|   |       `-- studentRouter.go
|   `-- teacherApi
|       |-- Wire_teacher.go
|       |-- teacherResthandler
|       |   `-- teacherResthandler.go
|       `-- teacherRouter
|           `-- teacherRouter.go
|-- backendDeployment.yaml
|-- configMap.yaml
|-- databaseDeployment.yaml
|-- docker-compose.yaml
|-- go.mod
|-- go.sum
|-- ingress.yaml
|-- main
|-- main.go
|-- package-lock.json
|-- package.json
|-- pkg
|   |-- attendance
|   |   |-- Wire_attendance.go
|   |   |-- attendanceAdapter
|   |   |   `-- attendanceAdapter.go
|   |   |-- attendanceIntegrationTests
|   |   |   |-- addAttendance_test.go
|   |   |   |-- fetchAttendanceOfUser_test.go
|   |   |   |-- fetchAttendance_test.go
|   |   |   |-- fetchClassMapPunch_test.go
|   |   |   |-- fetchDay_test.go
|   |   |   |-- fetchPunch_test.go
|   |   |   |-- hasAttendance_test.go
|   |   |   |-- punchOutCheck_test.go
|   |   |   |-- punchOut_test.go
|   |   |   `-- updatePunchOut_test.go
|   |   |-- attendanceRepository
|   |   |   |-- attendanceModels
|   |   |   |   `-- attendanceModel.go
|   |   |   |-- attendanceRepository.go
|   |   |   `-- mock_attendance_repository.go
|   |   `-- attendanceServices
|   |       |-- attendanceService.go
|   |       `-- attendanceServiceBean
|   |           `-- attendanceServiceBean.go
|   |-- dashboard
|   |   |-- Wire_dashboard2.go
|   |   |-- dashboardRepository
|   |   |   `-- dashboardRepository.go
|   |   `-- dashboardServices
|   |       `-- dashboardServices.go
|   |-- login
|   |   |-- Wire_login2.go
|   |   |-- loginRepository
|   |   |   |-- loginModels
|   |   |   |   `-- loginModels.go
|   |   |   |-- loginRepository.go
|   |   |   `-- mock_login_repository.go
|   |   `-- loginServices
|   |       |-- loginServices.go
|   |       |-- loginServices_test.go
|   |       |-- mock_login_services.go
|   |       |-- private.pem
|   |       |-- private_pkcs8.pem
|   |       `-- public.pem
|   |-- principal
|   |   |-- Wire_principal2.go
|   |   |-- principalRepository
|   |   |   |-- mock_principal_repository.go
|   |   |   `-- principalRepository.go
|   |   |-- principalServices
|   |   |   |-- principalServiceBean
|   |   |   |   `-- principalServiceBean.go
|   |   |   `-- principalServices.go
|   |   `-- principalUnitTest
|   |       `-- principal_test.go
|   |-- student
|   |   |-- Wire_student2.go
|   |   |-- studentRepository
|   |   |   |-- mock_student_repository.go
|   |   |   `-- studentRepository.go
|   |   |-- studentServices
|   |   |   |-- studentServiceBean
|   |   |   |   `-- studentServiceBean.go
|   |   |   `-- studentServices.go
|   |   `-- studentUnitTests
|   |       `-- student_test.go
|   |-- teacher
|   |   |-- Wire_teacher2.go
|   |   |-- teacherRepository
|   |   |   |-- mock_teacher_repository.go
|   |   |   `-- teacherRepo.go
|   |   |-- teacherServices
|   |   |   |-- teacherServiceBean
|   |   |   |   `-- teacherServiceBean.go
|   |   |   `-- teacherServices.go
|   |   `-- teacherUnitTests
|   |       `-- teacher_test.go
|   `-- user
|       |-- WIre_user.go
|       |-- userAdapter
|       |   `-- userAdapter.go
|       |-- userIntegrationTests
|       |   |-- checkEnrollment_test.go
|       |   |-- checkPunchout_test.go
|       |   |-- classMappingTeacher_test.go
|       |   |-- fetchClass_test.go
|       |   |-- fetchStudent_test.go
|       |   |-- fetchUser_test.go
|       |   |-- insertClassMap_test.go
|       |   |-- insertClass_test.go
|       |   `-- insertingUser_test.go
|       |-- userRepository
|       |   |-- mock_user_repository.go
|       |   |-- userModels
|       |   |   `-- userModels.go
|       |   `-- userRepository.go
|       |-- userServices
|       |   |-- userServiceBean
|       |   |   `-- userServiceBean.go
|       |   `-- userServices.go
|       `-- userUnitTests
|           `-- fetchUser_test.go
|-- private.pem
|-- private_pkcs8.pem
|-- public.pem
|-- secret.yaml
`-- wire_gen.go

```

### Frontend

```
.
|-- Dockerfile
|-- cors-crd.yaml
|-- cors.yaml
|-- fconfigmap.yaml
|-- frontendDeployment.yaml
|-- nginx.conf
|-- package-lock.json
|-- package.json
|-- public
|   |-- index.html
|   |-- manifest.json
|   `-- robots.txt
`-- src
    |-- App.js
    |-- components
    |   |-- Dashboard.js
    |   |-- Login.js
    |   |-- Logout.js
    |   |-- login.css
    |   |-- principal
    |   |   |-- AddStudent.js
    |   |   |-- AddTeacher.js
    |   |   |-- ListStudents.js
    |   |   |-- ListTeachers.js
    |   |   |-- PrincipalAttendance.js
    |   |   |-- PrincipalDashboard.js
    |   |   |-- PrincipalSideBar.js
    |   |   `-- PrincipalTeacherAttendancePage.js
    |   |-- student
    |   |   |-- StudentAttendance.js
    |   |   |-- StudentDashboard.js
    |   |   |-- StudentPunchIn.js
    |   |   |-- StudentPunchOut.js
    |   |   `-- StudentSideBar.js
    |   `-- teacher
    |       |-- TeacherAttendance.js
    |       |-- TeacherAttendancePage.js
    |       |-- TeacherClassAttendance.js
    |       |-- TeacherClassAttendancePage.js
    |       |-- TeacherDashboard.js
    |       |-- TeacherPunchIn.js
    |       |-- TeacherPunchOut.js
    |       `-- TeacherSideBar.js
    `-- index.js
```

## Models

### UserRoleType Enumeration:
```go
type UserRoleType string

const (
Principal UserRoleType = "principal"
Teacher   UserRoleType = "teacher"
Student   UserRoleType = "student"
)
```
#### Purpose
Enumerates different user roles in the system.
#### Constants
**Principal:** Represents the role of a principal.

**Teacher:** Represents the role of a teacher.

**Student:** Represents the role of a student.

### User Model:
```go
type User struct {

tableName struct{}           `sql:"user"`
ID        string             `json:"id" pg:",pk"`
Username  string             `json:"username"`
Password  string             `json:"password"`
Role      UserRoleType       `json:"role"` // "principal", "teacher", or "student"
ClassMap  []ClassMappingUser `json:"-" pg:"rel:has-many"`

}
```
#### Purpose
Represents generic user information.
#### Fields
**ID:** Unique identifier for each user.

**Username:** User's username for authentication.

**Password:** User's password for authentication.

**Role:** Enumerated type (UserRoleType) representing the role of the user (principal, teacher, or student).

**ClassMap:** List of ClassMappingUser records indicating the mapping between the user and multiple classes. 

The "has-many" relationship denotes that one user can be associated with several class mappings, 

allowing for the representation of the user's connection to various classes.

### Class Model:
```go
type Class struct {
tableName struct{} `sql:"class"`
ClassID   int      `json:"class_id" pg:",pk"`
ClassName string   `json:"class_name"`
}
```
#### Purpose
Represents information about classes.

#### Fields:
**ClassID:** Unique identifier for each class.

**ClassName:** Name for the class.

### ClassMappingUser Model:
```go
type ClassMappingUser struct {
tableName struct{} `sql:"classmapuser"`
ID        int      `json:"id" pg:",pk"`
UserID    string   `json:"user_id" pg:",fk"`
ClassID   int      `json:"class_id" pg:",fk"`
Class     Class    `json:"-" pg:"rel:has-one"`
}
```
#### Purpose
Establishes the relationship between users and classes.

#### Fields
**ID:** Unique identifier for each mapping record.

**UserID:** Foreign key referencing the User model, indicating the user associated with the mapping.

**ClassID:** Foreign key referencing the Class model, indicating the class associated with the mapping.

**Class:** Represents the related Class information.
### Attendance Model:
```go
type Attendance struct {
tableName struct{}     `sql:"attendance"`
ID        int          `json:"id" pg:",pk"`
UserID    string       `json:"user_id" pg:",fk"`
Day       int          `json:"day"`
Month     int          `json:"month"`
Year      int          `json:"year"`
PunchMap  []PunchInOut `json:"-" pg:"rel:has-many"`
}
```
#### Purpose
Represents individual attendance records for users.
#### Fields
**ID:** Unique identifier for each attendance record.

**UserID:** Foreign key referencing the User model, indicating the user associated with the attendance record.

Day, Month, Year: Date components representing the day of attendance.

**PunchMap:** List of PunchInOut records representing punch-in and punch-out instances associated with this attendance.

### PunchInOut Model:
```go
type PunchInOut struct {
tableName    struct{}                 `sql:"punchinout"`
ID           int                      `json:"id" pg:",pk"`
AttendanceID int                      `json:"attendance_id" pg:",fk"`
UserID       string                   `json:"user_id"`
PunchIn      string                   `json:"punch_in" `
PunchOut     string                   `json:"punch_out" `
ClassMap     []ClassMappingAttendance `json:"-"`
}
```
#### Purpose
Represents individual punch-in or punch-out records for attendance.

#### Fields:
**ID:** Unique identifier for each punch-in/out record.

**AttendanceID:** Foreign key referencing the Attendance model, 
indicating the attendance record associated with the punch-in/out.

**UserID:** Represents the user associated with the punch-in/out.

**PunchIn, PunchOut:** Timestamps indicating the times of punch-in and punch-out.

**ClassMap:** List of ClassMappingAttendance records representing the mapping between this punch-in/out event and classes.

### ClassMappingAttendance Model:
```go
type ClassMappingAttendance struct {
tableName struct{} `sql:"classmapattendance"`
ID        int      `json:"id" pg:",pk"`
PunchID   int      `json:"punch_id" pg:",fk"`
ClassID   int      `json:"class_id" pg:",fk"`
Class     Class    `json:"-" pg:"rel:has-one"`
}
```
#### Purpose
Establishes the relationship between punch-in/out records and classes.
#### Fields
**ID:** Unique identifier for each mapping record.

**PunchID:** Foreign key referencing the PunchInOut model, indicating the punch-in/out event associated with the mapping.

**ClassID:** Foreign key referencing the Class model, indicating the class associated with the mapping.

**Class:** Represents the related Class information.


### JWT Model:
```go
type JWT struct {
PrivateKey []byte
PublicKey  []byte
}
```
#### Purpose
Represents the structure for handling JSON Web Tokens (JWT) for authentication and authorization.
#### Fields

**PrivateKey:** Private key used for JWT generation and decoding.

**PublicKey:** Public key used for JWT verification.

## API Routes

### Principal Routes:
**/principal/addStudent:** POST endpoint to add a new student to the system.

**/principal/addTeacher:** POST endpoint to add a new teacher to the system.

**/principal/teacherAttendance:** GET endpoint to retrieve attendance records for teachers.

**/principal/Students:** GET endpoint to retrieve a list of students managed by the principal.

**/principal/Teachers:** GET endpoint to retrieve a list of teachers managed by the principal.

**/principal/Class:** POST endpoint to add classes managed by principal

**/principal/Classes:** GET endpoint to retrieve classes managed by principal
### Teacher Routes:
**/teacher/punchIn:** POST endpoint for teachers to record a punch-in event.

**/teacher/punchOut:** POST endpoint for teachers to record a punch-out event.

**/teacher/attendance:** GET endpoint to retrieve attendance records for the teacher.

**/teacher/classAttendance:** GET endpoint to retrieve class-wise attendance records for the teacher.

### Student Routes:
**/student/punchIn:** POST endpoint for students to record a punch-in event.

**/student/punchOut:** POST endpoint for students to record a punch-out event.

**/student/attendance:** GET endpoint to retrieve attendance records for the student.

### Login Routes:
**/login:** GET and POST endpoint for user authentication.

**/logout:** GET endpoint for user logout.

### Dashboard Routes:
**/dashboard/principal:** GET endpoint to access the principal's dashboard.

**/dashboard/teacher:** GET endpoint to access the teacher's dashboard.

**/dashboard/student:** GET endpoint to access the student's dashboard.

## Frontend Routes

**/login** : Login page

**/logout** : Logout page

**/dashboard/:role** : Dashboard w.r.t roles 

**/teacher/class-attendance** : Retrieve  attendance of specified class ,providing input classname, day, month, year by user whose role is teacher

**/teacher/general-attendance** : Retrieve attendance of teacher (user) by providing input userID,month and year

**/principal/teacher-attendance** : Retrieve attendance of specified teacher ID , providing input month and year by principal 

**/principal/students** : List of students with classes

**/principal/classes** : Add class and get list of all classes by principal

**/student/student-attendance** :  Retrieve attendance of student (user) by providing input userID,month and year

**/** : To avoid any route apart from above list by redirecting to login

## Credentials (Principal)

Username : user1

Password: user1




https://github.com/komalreddy3/Attendance-go/assets/82363361/786029a4-bb26-40a7-a0c3-4cf96a82fbf9



## Preview

![Screenshot from 2024-02-22 10-55-24](https://github.com/komalreddy3/Attendance-go/assets/82363361/5bb09ae6-a5ef-4b5c-9086-40b339c4da6a)

![Screenshot from 2024-02-22 10-55-49](https://github.com/komalreddy3/Attendance-go/assets/82363361/9a26bf74-c0f1-4a09-b5c9-10dbe8b0ca96)


![Screenshot from 2024-02-22 10-56-53](https://github.com/komalreddy3/Attendance-go/assets/82363361/dbfbedc1-cb87-4ce6-a25b-2094a4513bbd)



![Screenshot from 2024-02-22 10-58-43](https://github.com/komalreddy3/Attendance-go/assets/82363361/1a32deaf-1134-4681-80a9-5f275f669cdd)


![Screenshot from 2024-02-22 10-58-58](https://github.com/komalreddy3/Attendance-go/assets/82363361/71b0bbff-473e-4db0-ac06-0d5d102490f1)


![Screenshot from 2024-02-22 10-59-08](https://github.com/komalreddy3/Attendance-go/assets/82363361/c4a05c01-42b4-4256-9950-69e96e604775)

![Screenshot from 2024-02-22 10-59-32](https://github.com/komalreddy3/Attendance-go/assets/82363361/9b9a0641-1e30-490e-b858-ae30bf39af72)


![Screenshot from 2024-02-22 10-59-43](https://github.com/komalreddy3/Attendance-go/assets/82363361/8f709275-f96c-4b74-a618-b13a6a217780)

![Screenshot from 2024-02-22 11-02-19](https://github.com/komalreddy3/Attendance-go/assets/82363361/71e85f1b-1750-4f84-ad8e-99787f0f8f90)

![Screenshot from 2024-02-22 11-02-47](https://github.com/komalreddy3/Attendance-go/assets/82363361/05aeb0af-c9f5-416b-a9cf-02049ae877d4)

![Screenshot from 2024-02-22 11-02-57](https://github.com/komalreddy3/Attendance-go/assets/82363361/d4e10c9a-27eb-4361-88a1-f28b8a661143)






















