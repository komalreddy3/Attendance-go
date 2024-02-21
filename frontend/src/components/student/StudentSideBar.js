// StudentSidebar.js
import React from 'react';
import { Link } from 'react-router-dom';
import Logout from "../Logout";

const StudentSidebar = ({ isNavVisible }) => {

    // const styles = {
    //     sidebar: {
    //         width: '200px',
    //         height: '100vh',
    //         color: '#fff',
    //         padding: '20px',
    //         boxShadow: '2px 0 5px rgba(0, 0, 0, 0.1)',
    //         position: 'fixed',
    //         left: 0,
    //         top: 0,
    //         fontSize: "1.5rem",
    //         fontWeight: "400",
    //         lineHeight: "1.5",
    //         backgroundColor: "#212529"
    //     },
    //     link: {
    //         fontSize: '20px',
    //         borderRadius:'6px',
    //         padding: '4px',
    //         marginBottom: '10px',
    //         textDecoration: 'none',
    //         color: '#fff',
    //         transition: 'background-color 0.3s',
    //     }
    // };
    const handleClick = () => {
        isNavVisible=!isNavVisible;
    };
    return (

        // <div style={styles.sidebar}>
        //     <h3>Sidebar</h3>
        //
        //     <div style={{position: "relative"}}>
        //         <div id="hov" style={{...styles.link}} className={isNavVisible ? "showNav hideNav" : "hidden"}>
        //             <Link to="/student/student-attendance">Student Attendance</Link>
        //         </div>
        //         <div id="hov" style={{...styles.link}} className={isNavVisible ? "showNav hideNav" : "hidden"}>
        //             <Link to="/dashboard/student">Home</Link>
        //         </div>
        //
        //     </div>
        // </div>
        <div>

            <div className={isNavVisible ? "nav-container showNav " : "nav-container hideNav"}>

                <ul className="nav-list">
                    <li className="list-item"><Link to="/student/student-attendance">Student Attendance</Link></li>
                    <li className="list-item"><Link to="/dashboard/student">Home</Link></li>
                    <li className="list-item"><Logout/></li>


                </ul>
            </div>
        </div>

    );
};

export default StudentSidebar;