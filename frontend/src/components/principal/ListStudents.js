// PrincipalStudents.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
// import {Link} from "react-router-dom";
import PrincipalSidebar from "./PrincipalSideBar";
import { saveAs } from 'file-saver';
import { jsPDF } from 'jspdf';
import 'jspdf-autotable';
import StudentSidebar from "../student/StudentSideBar";
import TeacherSidebar from "../teacher/TeacherSideBar";

const PrincipalStudents = () => {
    const [red, setred] = useState(false);
    const [students, setStudents] = useState([]);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchPrincipalStudents = async () => {
            try {
                const response = await axios.get('/api/principal/Students', {
                    withCredentials: true,
                });
                 console.log(response)
                if(response.data==null){
                    setError('No students available');
                }

                if (response.status === 200) {
                    setStudents(response.data);
                    setError('');
                } else {
                    setError('Failed to fetch list of students.');
                }
            } catch (error) {
                setred(true)
                window.location.href = `/login`;
                console.error('Error during fetching list of students:', error);
                setError('An error occurred while processing your request.');
            }
        };

        fetchPrincipalStudents();
    }, []); // Empty dependency array ensures useEffect runs only once on component mount
    const convertToCSV = () => {
        const csvRows = [];
        const headers = ["Username", "Class"];
        csvRows.push(headers.join(','));

        students.forEach(student => {
            const rowData = [
                student.username,
                student.classnames.join(', ')
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
        saveAs(blob, 'students.csv');
    };

    // Function to handle PDF download
    const handleDownloadPDF = () => {
        const doc = new jsPDF();
        doc.autoTable({ html: '#students-table', startY: 10 });
        doc.save('students.pdf');
    };
    const [isNavVisible, setNavVisible] = useState(false);

    const handleClick = () => {
        setNavVisible(!isNavVisible);
    };
    return (

       <div>
           {red===false &&
           (<div>
           {/*<PrincipalSidebar/>*/}
               <button className="btn-nav" style={{ background: "none",
                   border: "none"}} onClick={handleClick}>
                   <div className="bar arrow-top-r"></div>
                   <div className="bar arrow-middle-r"></div>
                   <div className="bar arrow-bottom-r"></div>
               </button>


               <PrincipalSidebar isNavVisible={isNavVisible}/>
           {/*<li style={{"listStyle": "none"}} id="hov2"><Link id="hov3" to="/dashboard/principal">Dashboard</Link></li>*/}
           <h2>List of Students</h2>
           {students === null && <p style={{color: 'red'}}>No student available</p>}

           {/* Display the list of students */}
           {/*<ul>*/}
           {/*    {students.map((student) => (*/}
           {/*        <li key={student.username}>*/}
           {/*            {student.username}*/}
           {/*        </li>*/}
           {/*    ))}*/}
           {/*</ul>*/}

           {students !== null && students !== undefined && (
               // <ul>
               //     {students.map((student) => (
               //         <li key={student.username}>
               //             {student.username}
               //         </li>
               //
               //     ))}
               // </ul>
               <table style={{borderCollapse: 'collapse', width: '100%'}} id="students-table">
                   <thead style={{backgroundColor: '#f2f2f2'}}>
                   <tr>
                       <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Username</th>
                       <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Class</th>
                   </tr>
                   </thead>
                   <tbody>
                   {students.map((student) => (
                       <tr key={student.username}>
                           <td style={{padding: '8px', border: '1px solid #ddd'}}>{student.username}</td>
                           <td style={{padding: '8px', border: '1px solid #ddd'}}>
                               <ul>
                                   {student.classnames.map((classname, index) => (
                                       <li key={index}>{classname}</li>
                                   ))}
                               </ul>
                           </td>
                       </tr>
                   ))}
                   </tbody>
               </table>

           )}

           <button onClick={handleDownloadCSV}>Download CSV</button>
           <button onClick={handleDownloadPDF}>Download PDF</button>
       </div>)
       }
           {/*{red === true && <Login/>}*/}
       </div>


    )
        ;
};

export default PrincipalStudents;