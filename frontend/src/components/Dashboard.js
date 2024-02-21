// Dashboard.js

import React, { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import StudentDashboard from "./student/StudentDashboard";
import TeacherDashboard from "./teacher/TeacherDashboard";
import PrincipalDashboard from "./principal/PrincipalDashboard";
import StudentSidebar from "./student/StudentSideBar";
import TeacherSidebar from "./teacher/TeacherSideBar";
import PrincipalSidebar from "./principal/PrincipalSideBar";
import {useState} from "react";
// import StudentSidebar from "./student/StudentSideBar";
// import TeacherSidebar from "./teacher/TeacherSideBar";
// import PrincipalSidebar from "./principal/PrincipalSideBar";
const Dashboard = () => {
    const { role } = useParams();
    //const [dashboardData, setDashboardData] = useState(null);

    console.log(role)
    useEffect(() => {
        // Fetch dashboard data from the backend using the 'role' parameter
        const fetchData = async () => {
            try {
                const response = await fetch(`/api/dashboard/${role}`, {
                    method: 'GET',
                    credentials: 'include',  // Ensure that credentials are included
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });

                if (role!=='principal'){
                    if(role!=='teacher'){
                        if(role!=='student'){
                            window.location.href = "/login";
                            return;
                        }
                    }

                }
                console.log(response.status)
                if (response.status === 405) {
                    // Handle redirection on the frontend
                    window.location.href = "/login";
                    return;
                }
                    //await fetch(`http://localhost:8080/dashboard/${role}`);
                if (!response.ok) {
                    throw new Error('Failed to fetch dashboard data');
                }

                // const data = await response.json();
                // setDashboardData(data);
            } catch (error) {
                console.error(error);
            }
        };

        fetchData();
    }, [role]);
    const [isNavVisible, setNavVisible] = useState(false);

    const handleClick = () => {
        setNavVisible(!isNavVisible);
    };
    return (
        // <div>
        //
        //     {role === 'student' &&
        //         <StudentSidebar />
        //     }
        //     {role === 'teacher' &&
        //         <TeacherSidebar />
        //     }
        //     {role === 'principal' &&
        //         <PrincipalSidebar />
        //     }
        //     {/*<div style={{*/}
        //     {/*    display: "flex",*/}
        //     {/*    flexDirection: "column",*/}
        //     {/*    alignItems: "end"*/}
        //     {/*}}><Logout/></div>*/}
        //
        //
        //     <h1>Welcome to the Dashboard, {role}!</h1>
        //     {/*{dashboardData && (*/}
        //     {/*    <div>*/}
        //     {/*        <h2>Dashboard Data</h2>*/}
        //     {/*        /!* Render your dashboard data here *!/*/}
        //     {/*        /!*<pre>{JSON.stringify(dashboardData, null, 2)}</pre>*!/*/}
        //     {/*    </div>*/}
        //     {/*)}*/}
        //     {role === 'student' && <StudentDashboard/>}
        //     {role === 'teacher' && <TeacherDashboard/>}
        //     {role === 'principal' && <PrincipalDashboard/>}
        // </div>
        <div>
            <button className="btn-nav" style={{ background: "none",
                border: "none"}} onClick={handleClick}>
                <div className="bar arrow-top-r"></div>
                <div className="bar arrow-middle-r"></div>
                <div className="bar arrow-bottom-r"></div>
            </button>

            {role === 'student' && <StudentSidebar isNavVisible={isNavVisible}/>}
            {role === 'teacher' && <TeacherSidebar isNavVisible={isNavVisible}/>}
            {role === 'principal' && <PrincipalSidebar isNavVisible={isNavVisible}/>}


            <div>
                <h1>Welcome to the Dashboard, {role}!</h1>

                {role === 'student' && <StudentDashboard/>}
                {role === 'teacher' && <TeacherDashboard/>}
                {role === 'principal' && <PrincipalDashboard/>}
            </div>
        </div>
    );
};

export default Dashboard;
