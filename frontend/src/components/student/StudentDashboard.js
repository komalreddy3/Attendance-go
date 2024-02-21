// StudentDashboard.js

import React, {useEffect} from 'react';
import StudentPunchIn from './StudentPunchIn';
import StudentPunchOut from "./StudentPunchOut";
import axios from "axios";

// import StudentSidebar from "./StudentSideBar";
// import Logout from "../Logout";

const StudentDashboard = () => {
    useEffect(() => {
        // Fetch dashboard data from the backend using the 'role' parameter
        const fetchData = async () => {
            try {
                const response = await axios.get('/api/dashboard/student', {
                    method: 'GET',
                    credentials: 'include',  // Ensure that credentials are included
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });
                console.log(response.status)
                if (response.status === 405) {
                    // Handle redirection on the frontend
                    window.location.href = "/login";
                    return;
                }
                //await fetch(`http://localhost:8080/dashboard/${role}`);
                // if (!response.ok) {
                //     throw new Error('Failed to fetch dashboard data');
                // }

                // const data = await response.json();
                // setDashboardData(data);
            } catch (error) {
                window.location.href = `/login`;
                console.error(error);
            }
        };

        fetchData();
    }, );
    return (
        <div>
            {/*<div style={{*/}
            {/*    display: "flex",*/}
            {/*    flexDirection: "column",*/}
            {/*    alignItems: "end"*/}
            {/*}}><Logout/></div>*/}
            {/*<h1>Student Dashboard</h1>*/}
            {/*<div style={{display: "flex"}}>*/}
            {/*    <StudentSidebar/>*/}
            {/*</div>*/}
            <div style={{
                display: "grid",
                gridTemplateColumns: "repeat(2, 1fr)",
                gap: "20px"
            }}>


                {/* Other content specific to student dashboard */}
                <StudentPunchIn/>
                <StudentPunchOut/>

                {/*<StudentAttendance/>*/}
            </div>
        </div>
    );
};

export default StudentDashboard;
