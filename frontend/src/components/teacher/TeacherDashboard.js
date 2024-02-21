// TeacherDashboard.js

import React, {useEffect} from 'react';

import TeacherPunchIn from "./TeacherPunchIn"
import TeacherPunchOut from "./TeacherPunchOut"
import axios from "axios";

// import TeacherSidebar from "./TeacherSideBar";
// import {Link} from "react-router-dom";
// import Logout from "../Logout";
// import TeacherSidebar from "./TeacherSideBar";


const TeacherDashboard = () => {
    useEffect(() => {
        // Fetch dashboard data from the backend using the 'role' parameter
        const fetchData = async () => {
            try {
                const response = await axios.get('/api/dashboard/teacher', {
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
            {/*<div style={{display: "flex"}}>*/}
            {/*    <TeacherSidebar/>*/}
            {/*<div>*/}

                {/*<div style={{*/}
                {/*    display: "flex",*/}
                {/*    flexDirection: "column",*/}
                {/*    alignItems: "end"*/}
                {/*}}><Logout/></div>*/}
                <div>
                    {/*<h1>Teacher Dashboard</h1>*/}
                    {/*<li style={{"listStyle": "none"}} id="hov2"><Link id="hov3" to="/logout">Logout</Link> </li>*/}

                    {/*<div style={{display: "flex"}}>*/}
                    {/*    <TeacherSidebar/>*/}


                    {/*</div>*/}
                    <div style={{
                        display: "grid",
                        gridTemplateColumns: "repeat(2, 1fr)",
                        gap: "20px"
                    }}>


                        {/* Other content specific to student dashboard */}
                        <TeacherPunchIn/>
                        <TeacherPunchOut/>
                    </div>

                </div>
            {/*</div>*/}
        </div>
    );
};

export default TeacherDashboard;
