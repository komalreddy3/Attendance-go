// PrincipalStudents.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
// import {Link} from "react-router-dom";
import PrincipalSidebar from "./PrincipalSideBar";
import Login from "../Login";
import { saveAs } from 'file-saver';
import { jsPDF } from 'jspdf';
import 'jspdf-autotable';

const PrincipalTeachers = () => {
    const [red, setred] = useState(false);
    const [teachers, setTeachers] = useState([]);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchPrincipalTeachers = async () => {
            try {
                const response = await axios.get('/api/principal/Teachers', {
                    withCredentials: true,
                });
                if(response.data==null){
                    setError('No teachers available');
                }
                console.log(response)

                if (response.status === 200) {
                    setTeachers(response.data);
                    setError('');
                } else {
                    setError('Failed to fetch list of teachers.');
                }
            } catch (error) {
                console.error('Error during fetching list of teachers:', error);
                window.location.href = `/login`;
                setred(true)
                setError('An error occurred while processing your request.');
            }
        };

        fetchPrincipalTeachers();
    }, []); // Empty dependency array ensures useEffect runs only once on component mount
    const convertToCSV = () => {
        const csvRows = [];
        const headers = ["Username", "Classes"];
        csvRows.push(headers.join(','));

        teachers.forEach(teacher => {
            const rowData = [
                teacher.username,
                teacher.classnames.join(', ')
            ];
            csvRows.push(rowData.join(','));
        });

        const csvData = csvRows.join('\n');
        return csvData;
    };

    // Function to handle CSV download
    const handleDownloadCSV = () => {
        const csvData = convertToCSV();
        const blob = new Blob([csvData], { type: 'text/csv;charset=utf-8' });
        saveAs(blob, 'teachers.csv');
    };

    // Function to handle PDF download
    const handleDownloadPDF = () => {
        const doc = new jsPDF();
        doc.autoTable({ html: '#teachers-table', startY: 10 });
        doc.save('teachers.pdf');
    };
    const [isNavVisible, setNavVisible] = useState(false);

    const handleClick = () => {
        setNavVisible(!isNavVisible);
    };
    return (
        <div>
            {red === false && (<div>
                {/*<PrincipalSidebar/>*/}
                {/*<li style={{"listStyle": "none"}} id="hov2"><Link id="hov3" to="/dashboard/principal">Dashboard</Link></li>*/}
                <button className="btn-nav" style={{
                    background: "none",
                    border: "none"
                }} onClick={handleClick}>
                    <div className="bar arrow-top-r"></div>
                    <div className="bar arrow-middle-r"></div>
                    <div className="bar arrow-bottom-r"></div>
                </button>


                <PrincipalSidebar isNavVisible={isNavVisible}/>
                <h2>List of Teachers</h2>

                {teachers === null && <p style={{color: 'red'}}>No teacher available</p>}

                {/* Display the list of students */}
                {/*<ul>*/}
                {/*    {teachers.map((teacher) => (*/}
                {/*        <li key={teacher.username}>*/}
                {/*            {teacher.username}*/}
                {/*        </li>*/}
                {/*    ))}*/}
                {/*</ul>*/}
                {teachers !== null && teachers !== undefined && (
                    // <ul>
                    //     {teachers.map((teacher) => (
                    //         <li key={teacher.username}>
                    //             {teacher.username}
                    //         </li>
                    //     ))}
                    // </ul>
                    <table style={{borderCollapse: 'collapse', width: '100%'}} id="teachers-table">
                        <thead style={{backgroundColor: '#f2f2f2'}}>
                        <tr>
                            <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Username</th>
                            <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Classes</th>
                        </tr>
                        </thead>
                        <tbody>
                        {teachers.map((teacher) => (
                            <tr key={teacher.username}>
                                <td style={{padding: '8px', border: '1px solid #ddd'}}>{teacher.username}</td>
                                <td style={{padding: '8px', border: '1px solid #ddd'}}>
                                    <ul>
                                        {teacher.classnames.map((classname, index) => (
                                            <li key={index}>{classname}</li>
                                        ))}
                                    </ul>
                                </td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                )}
            </div>)}
            {/*{red === true && <Login/>}*/}
            <button onClick={handleDownloadCSV}>Download CSV</button>
            <button onClick={handleDownloadPDF}>Download PDF</button>
        </div>


    );
};

export default PrincipalTeachers;