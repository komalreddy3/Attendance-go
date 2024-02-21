// PrincipalDashboard.js

import React, {useEffect, useState} from 'react';


// import PrincipalSidebar from "./PrincipalSideBar";

import AddStudent from "./AddStudent";
import AddTeacher from "./AddTeacher";
import 'boxicons'

import axios from "axios";
import {Link} from "react-router-dom";
import AddClass from "./AddClass";

const PrincipalDashboard = () => {
    const [students, setStudents] = useState([]);
    const [teachers, setTeachers] = useState([]);

    useEffect(() => {
        // Fetch dashboard data from the backend using the 'role' parameter
        const fetchData = async () => {
            try {
                const response = await axios.get('/api/dashboard/principal', {
                    method: 'GET',
                    credentials: 'include',  // Ensure that credentials are included
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });
                console.log(response.status)
                // if (response.status === 405) {
                //     // Handle redirection on the frontend
                //     window.location.href = "/login";
                //     return;
                // }
                //await fetch(`http://localhost:8080/dashboard/${role}`);
                // if (!response.ok) {
                //     throw new Error('Failed to fetch dashboard data');
                // }

                // const data = await response.json();
                // setDashboardData(data);
            } catch (error) {
                console.log(error)
               window.location.href = `/login`;
                console.error(error);
            }
        };

        fetchData();

    }, );
    useEffect(()=>{
        const interval = setInterval(() => { const fetchPrincipalStudents = async () => {
            try {
                const response = await axios.get('/api/principal/Students', {
                    withCredentials: false,
                });
                console.log(response)
                if(response.data==null){
                    //setError('No students available');
                }

                if (response.status === 200) {
                    setStudents(response.data);
                    //  setError('');
                } else {
                    // setError('Failed to fetch list of students.');
                }
            } catch (error) {
                // setred(true)
                //  window.location.href = `/login`;
                // console.error('Error during fetching list of students:', error);
                // setError('An error occurred while processing your request.');
            }
        };
            fetchPrincipalStudents();
        }, 500);
        return () => clearInterval(interval);

    },[])
    useEffect(()=>{
        const interval = setInterval(() => {   const fetchPrincipalTeachers = async () => {
            try {
                const response = await axios.get('/api/principal/Teachers', {
                    withCredentials: false,
                });
                // if(response.data==null){
                //     setError('No teachers available');
                // }
                // console.log(response)

                if (response.status === 200) {
                    setTeachers(response.data);
                    //setError('');
                } else {
                    //setError('Failed to fetch list of teachers.');
                }
            } catch (error) {
                // console.error('Error during fetching list of teachers:', error);
                // window.location.href = `/login`;
                // setred(true)
                // setError('An error occurred while processing your request.');
            }
        };

        fetchPrincipalTeachers();
        }, 500);
        return () => clearInterval(interval);
    },[])
    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
        }}>
            {/*<div style={{*/}
            {/*    display: "flex",*/}
            {/*    flexDirection: "column",*/}
            {/*    alignItems: "end"*/}
            {/*}}><Logout/></div>*/}
            {/*<div style={{display: "flex"}}>*/}
            {/*    <PrincipalSidebar/>*/}
            {/*</div>*/}
            <div style={{
                display: "flex",

            }}>
                <div className="student-count-card">

                    {students !== null && (<div style={{display: "flex", justifyContent: "space-evenly"}}>
                        <box-icon type='solid' name='user' style={{
                            height: "100px",
                            width: "100px"
                        }}></box-icon>
                        <div className="content"><p style={{fontSize: "x-large"}}>No of Students</p><p style={{fontSize: "xx-large"}}>{students.length}</p></div>
                    </div>)}
                    {students === null && (<div style={{display: "flex", justifyContent: "space-evenly"}}>
                        <box-icon type='solid' name='user' style={{
                            height: "100px",
                            width: "100px"
                        }}></box-icon>
                        <div className="content"><p style={{fontSize: "x-large"}}>No of Students</p><p style={{fontSize: "xx-large"}}>0</p></div>
                    </div>)}

                    <div className="long-arrow-right"><Link to="/principal/students" className="arrow-button"
                                                            style={{
                                                                display: "block",
                                                                width: "45.25px",
                                                                height: "45.25px"
                                                            }}>List
                        of students</Link></div>
                    <p>Click Arrow to view more</p>

                </div>
                <div className="student-count-card">

                    {teachers !== null && (<div style={{display: "flex", justifyContent: "space-evenly"}}>
                        <box-icon type='solid' name='user' style={{
                            height: "100px",
                            width: "100px"
                        }}></box-icon>
                        <div className="content"><p style={{fontSize: "x-large"}}>No of Teachers</p><p style={{fontSize: "xx-large"}}>{teachers.length}</p></div>
                    </div>)}
                    {teachers === null && (<div style={{display: "flex", justifyContent: "space-evenly"}}>
                        <box-icon type='solid' name='user' style={{
                            height: "100px",
                            width: "100px"
                        }}></box-icon>
                        <div className="content"><p style={{fontSize: "x-large"}}>No of Teachers</p><p style={{fontSize: "xx-large"}}>0</p></div>
                    </div>)}

                    <div className="long-arrow-right"><Link to="/principal/teachers" className="arrow-button"
                                                            style={{
                                                                display: "block",
                                                                width: "45.25px",
                                                                height: "45.25px"
                                                            }}>List of Teachers</Link></div>
                    <p>Click Arrow to view more</p>

                </div>
            </div>
            <div style={{
                display: "grid",
                gridTemplateColumns: "repeat(2, 1fr)",
                gap: "20px"
            }}>


                {/* Other content specific to student dashboard */}
                <AddStudent />
                <AddTeacher />
                {/*<AddClass />*/}
            </div>


        </div>
    );
};

export default PrincipalDashboard;
