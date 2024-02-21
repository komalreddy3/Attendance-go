// TeacherAttendancePage.js
import React, {useEffect, useState} from 'react';
import TeacherAttendanceByMonth from './TeacherAttendance';
// import {Link} from "react-router-dom";
import TeacherSidebar from "./TeacherSideBar";
import DatePicker from "react-datepicker";
import axios from "axios";

const TeacherAttendancePage = () => {
    const [red, setred] = useState(false);
    const [isNavVisible, setNavVisible] = useState(false);

    const handleClick = () => {
        setNavVisible(!isNavVisible);
    };
    useEffect(() => {
        // Fetch dashboard data from the backend using the 'role' parameter
        const fetchDat = async () => {
            try {
                //const role="student"
                const response = await axios.get('/api/dashboard/teacher', {
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
    return (
       <div>
           {/*{red===true && <Login />}*/}
           {red===false && (<div>
               {/*<TeacherSidebar/>*/}
               <button className="btn-nav" style={{
                   background: "none",
                   border: "none"
               }} onClick={handleClick}>
                   <div className="bar arrow-top-r"></div>
                   <div className="bar arrow-middle-r"></div>
                   <div className="bar arrow-bottom-r"></div>
               </button>
               <TeacherSidebar isNavVisible={isNavVisible}/>
               {/*<li><Link to="/dashboard/teacher">Dashboard</Link></li>*/}
               {/*<li style={{"listStyle": "none"}} id="hov2"><Link id="hov3" to="/dashboard/teacher">Dashboard</Link></li>*/}


               <TeacherAttendanceByMonth/>
               {/* Other content related to general attendance */}
           </div>)}
       </div>
    );
};

export default TeacherAttendancePage;
