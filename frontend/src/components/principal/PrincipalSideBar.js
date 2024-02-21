// PrincipalSidebar.js
import React, {useState} from 'react';
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
const PrincipalSidebar = ({ isNavVisible }) => {
    const [scrollbarVisible, setScrollbarVisible] = useState(false); // State to track scrollbar visibility
    return (
        // <div onMouseEnter={() => setScrollbarVisible(true)}
        //      onMouseLeave={() => setScrollbarVisible(false)}>
        //    <p> click to sidebar</p>
        //    <div style={styles.sidebar}  className={`dashboard-container ${scrollbarVisible ? 'scrollbar-visible' : ''}`}>
        //     <h3>Sidebar</h3>
        //
        //
        //     <div style={{position: "relative",}}>
        //         <div id="hov" style={{
        //             ...styles.link,
        //         }}><Link to="/principal/teacher-attendance">Teacher Attendance</Link></div>
        //
        //         <div id="hov" style={{
        //             ...styles.link,
        //         }}><Link to="/principal/students">List of students</Link></div>
        //         <div id="hov" style={{
        //             ...styles.link,
        //         }}><Link to="/principal/teachers">List of Teachers</Link></div>
        //         <div id="hov" style={{
        //             ...styles.link,
        //         }}><Link to="/dashboard/principal">Home</Link></div>
        //         <div style={{
        //             display: "flex",
        //             flexDirection: "column",
        //             alignItems: "end",
        //             position: "absolute",
        //             bottom: "-69vh",
        //         }}><Logout/></div>
        //     </div>
        // </div>
        // </div>
        // <div style={{position: "relative",}}>
        //     <div id="hov" style={{
        //         ...styles.link,
        //     }}><Link to="/principal/teacher-attendance">Teacher Attendance</Link></div>
        //
        //     <div id="hov" style={{
        //         ...styles.link,
        //     }}><Link to="/principal/students">List of students</Link></div>
        //     <div id="hov" style={{
        //         ...styles.link,
        //     }}><Link to="/principal/teachers">List of Teachers</Link></div>
        //     <div id="hov" style={{
        //         ...styles.link,
        //     }}><Link to="/dashboard/principal">Home</Link></div>
        //     <div style={{
        //         display: "flex",
        //         flexDirection: "column",
        //         alignItems: "end",
        //         position: "absolute",
        //         bottom: "-69vh",
        //     }}><Logout/></div>
        // </div>
    <div className={isNavVisible ? "nav-container showNav " : "nav-container hidden"}>

        <ul className="nav-list">
            <li className="list-item"><Link to="/principal/teacher-attendance">Teacher Attendance</Link></li>
            <li className="list-item"><Link to="/principal/classes">Add Class</Link></li>
            <li className="list-item"><Link to="/principal/students">List of students</Link></li>
            <li className="list-item"><Link to="/principal/teachers">List of Teachers</Link></li>
            <li className="list-item"><Link to="/dashboard/principal">Home</Link></li>
            <li className="list-item"><Logout/></li>


        </ul>
    </div>
    )
        ;
};


export default PrincipalSidebar;
