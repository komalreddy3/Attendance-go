// Assuming this is in your main component where you set up routes
import {BrowserRouter as Router, Navigate, Route, Routes} from 'react-router-dom';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import TeacherClassAttendancePage from "./components/teacher/TeacherClassAttendancePage";
import TeacherAttendancePage from "./components/teacher/TeacherAttendancePage";
import PrincipalTeacherAttendancePage from "./components/principal/PrincipalTeacherAttendancePage";
import ListStudents from "./components/principal/ListStudents";
import ListTeachers from "./components/principal/ListTeachers";
import StudentAttendance from "./components/student/StudentAttendance";
import Logout from "./components/Logout";
import AddClass from "./components/principal/AddClass";

function App() {
    return (
        <Router basename="/">
            <Routes>
                <Route path="/login" element={<Login />} />
                <Route path="/logout" element={<Logout />} />
                <Route path="/dashboard/:role" element={<Dashboard />} />
                <Route path="/teacher/class-attendance" element={<TeacherClassAttendancePage />} />
                <Route path="/teacher/general-attendance" element={<TeacherAttendancePage />} />
                <Route path="/principal/teacher-attendance" element={<PrincipalTeacherAttendancePage />} />
                <Route path="/principal/students" element={<ListStudents />} />
                <Route path="/principal/teachers" element={<ListTeachers />} />
                <Route path="/principal/classes" element={<AddClass />} />
                <Route path="/student/student-attendance" element={<StudentAttendance />} />
                <Route path="*" element={<Navigate to="/login" />} />
            </Routes>
        </Router>
    );
}

export default App;

