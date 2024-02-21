// TeacherSidebar.js

import React from 'react';
import { Link } from 'react-router-dom';
import Logout from "../Logout";
const styles = {
    sidebar: {
        width: '200px',
        height: '100vh',
        color: '#fff',
        padding: '20px',
        boxShadow: '2px 0 5px rgba(0, 0, 0, 0.1)',
        position: 'fixed',
        left: 0,
        top: 0,

        fontSize: "1.5rem",
        fontWeight: "400",
        lineHeight: "1.5",
        backgroundColor: "#212529"
    },
    link: {
        fontSize: '20px',
        borderRadius:'6px',
        padding: '4px',
        marginBottom: '10px',
        textDecoration: 'none',
        color: '#fff',
        transition: 'background-color 0.3s',
    }
};
const TeacherSidebar = ({ isNavVisible }) => {
    return (
        // <div style={styles.sidebar}>
        //     <h3>Sidebar</h3>
        //     <div style={{position: "relative",}}>
        //        < div id="hov" style={{
        //         ...styles.link,
        //     }}>
        //         <Link to="/teacher/general-attendance">Teacher Attendance</Link>
        //     </div>
        //     <div id="hov" style={{
        //         ...styles.link,
        //     }}>
        //         <Link to="/teacher/class-attendance">Class Attendance</Link>
        //     </div>
        //     <div id="hov" style={{
        //         ...styles.link,
        //     }}><Link to="/dashboard/teacher">Home</Link></div>
        //
        //     <div style={{
        //         display: "flex",
        //         flexDirection: "column",
        //         alignItems: "end",
        //         position: "absolute",
        //         bottom: "-73vh",
        //     }}><Logout/></div>
        //     </div>
        // </div>
        <div className={isNavVisible ? "nav-container showNav " : "nav-container hidden"}>

            <ul className="nav-list">
                <li className="list-item"><Link to="/teacher/general-attendance">Teacher Attendance</Link></li>
                <li className="list-item"><Link to="/teacher/class-attendance">Class Attendance</Link></li>
                <li className="list-item"><Link to="/dashboard/teacher">Home</Link></li>
                <li className="list-item"><Logout/></li>
            </ul>
        </div>
    );
};

export default TeacherSidebar;


