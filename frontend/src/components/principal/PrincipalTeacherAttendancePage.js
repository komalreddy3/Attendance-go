// PrincipalTeacherAttendancePage.js
import React, {useEffect, useState} from 'react';
import PrincipalAttendance from "./PrincipalAttendance";
// import {Link} from "react-router-dom";
import PrincipalSidebar from "./PrincipalSideBar";
import axios from "axios";


const TeacherAttendancePage = () => {
    const [red, setred] = useState(false);
    useEffect(() => {
        // Fetch dashboard data from the backend using the 'role' parameter
        const fetchDat = async () => {
            try {
                //const role="student"
                const response = await axios.get('/api/dashboard/principal', {
                    method: 'GET',
                    credentials: 'include',  // Ensure that credentials are included
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    withCredentials: true,
                });
                console.log("called")
                console.log(response.status)

                //await fetch(`http://localhost:8080/dashboard/${role}`);
                // if (!response.ok) {
                //     throw new Error('Failed to fetch dashboard data');
                // }

                // const data = await response.json();
                // setDashboardData(data);
            } catch (error) {
                window.location.href = `/login`;
                setred(true)
                console.error(error);
            }
        };

        fetchDat();
    });
    const [isNavVisible, setNavVisible] = useState(false);

    const handleClick = () => {
        setNavVisible(!isNavVisible);
    };
    return (
        <div>
            <button className="btn-nav" style={{
                background: "none",
                border: "none"
            }} onClick={handleClick}>
                <div className="bar arrow-top-r"></div>
                <div className="bar arrow-middle-r"></div>
                <div className="bar arrow-bottom-r"></div>
            </button>
            <PrincipalSidebar isNavVisible={isNavVisible}/>
            {red === false && (<div>
                {/*<PrincipalSidebar/>*/}
                {/*<li style={{"listStyle":"none"}} id="hov2"><Link id="hov3" to="/dashboard/principal">Dashboard</Link></li>*/}
                <PrincipalAttendance/>
                {/* Other content related to general attendance */}
            </div>)}
            {/*{red===true && <Login />}*/}
        </div>
    );
};

export default TeacherAttendancePage;
