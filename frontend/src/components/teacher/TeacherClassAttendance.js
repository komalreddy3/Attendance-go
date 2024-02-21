import React, {useEffect, useState} from 'react';
import axios from 'axios';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { saveAs } from 'file-saver';
import { jsPDF } from 'jspdf';
import 'jspdf-autotable';
const TeacherClassAttendanceByMonth = () => {

    const [className, setClass] = useState('');
    // const [day, setDay] = useState('');
    // const [month, setMonth] = useState('');
    // const [year, setYear] = useState('');
    const [selectedDate, setSelectedDate] = useState(new Date());
    const [aData, setAttendanceData] = useState('');

    const handleGetAttendance = async () => {
        try {
            const response = await axios.post('/api/teacher/classAttendance', {
                class: className,
                // day:parseInt(day),
                // month: parseInt(month),
                // year: parseInt(year),
                day: selectedDate.getDate(),
                month: selectedDate.getMonth() + 1, // Months are zero-based
                year: selectedDate.getFullYear(),
            }, {
                headers: {
                    'Content-Type': 'application/json',
                },
                withCredentials: true,
            });
            console.log(response)
            console.log(response.data)
            //aData=response.data
            console.log(response.data.attendance)
            console.log(typeof(response.data))
            // console.log(JSON.parse(response.data))
            // console.log(JSON.parse(JSON.stringify(response.data)))
            // console.log(JSON.parse(response))
            // const DataOfAttendance=JSON.parse(JSON.stringify(response.data))
            // console.log(DataOfAttendance)
            // console.log(typeof(DataOfAttendance))
            //console.log(JSON.parse(DataOfAttendance))

            setAttendanceData(response.data.attendance);
            console.log(aData)
        } catch (error) {
            console.error('Error during getting teacher attendance:', error);
        }
    };
    const convertToCSV = () => {
        const csvRows = [];
        const headers = ["Date", "Id", "Punch In", "Punch Out"];
        csvRows.push(headers.join(','));

        Object.keys(aData).forEach(day => {
            aData[day].forEach(entry => {
                const rowData = [
                    selectedDate.getDate() + "/" + (selectedDate.getMonth() + 1) + "/" + selectedDate.getFullYear(),
                    entry.id,
                    entry.firstPunchIn,
                    entry.lastPunchOut
                ];
                csvRows.push(rowData.join(','));
            });
        });

        const csvData = csvRows.join('\n');
        return csvData;
    };

    // Function to handle CSV download
    const handleDownloadCSV = () => {
        const csvData = convertToCSV();
        const blob = new Blob([csvData], { type: 'text/csv;charset=utf-8' });
        saveAs(blob, 'ClassAttendance.csv');
    };

    // Function to handle PDF download
    const handleDownloadPDF = () => {
        const doc = new jsPDF();
        doc.autoTable({ html: '#attendance-table' });
        doc.save('ClassAttendance.pdf');
    };

    return (
        <div>
            {/*<h2>Get Student Attendance by Month</h2>*/}
            {/*<label>*/}
            {/*    Student ID:*/}
            {/*    <input*/}
            {/*        type="text"*/}
            {/*        value={studentID}*/}
            {/*        onChange={(e) => setStudentID(e.target.value)}*/}
            {/*    />*/}
            {/*</label>*/}
            {/*<br />*/}
            {/*<label>*/}
            {/*    Month:*/}
            {/*    <input*/}
            {/*        type="text"*/}
            {/*        value={month}*/}
            {/*        onChange={(e) => setMonth(e.target.value)}*/}
            {/*    />*/}
            {/*</label>*/}
            {/*<br />*/}
            {/*<label>*/}
            {/*    Year:*/}
            {/*    <input*/}
            {/*        type="text"*/}
            {/*        value={year}*/}
            {/*        onChange={(e) => setYear(e.target.value)}*/}
            {/*    />*/}
            {/*</label>*/}
            {/*<br />*/}
            {/*<button onClick={handleGetAttendance}>Get Attendance</button>*/}

            <div>
                <h2> Class Attendance by Month</h2>
                <label>
                    Class Name:
                    <input
                        type="text"
                        value={className}
                        onChange={(e) => setClass(e.target.value)}
                    />
                </label>
                <br/>
                {/*<label>*/}
                {/*    Date:*/}
                {/*    <input*/}
                {/*        type="text"*/}
                {/*        value={day}*/}
                {/*        onChange={(e) => setDay(e.target.value)}*/}
                {/*    />*/}
                {/*</label>*/}
                {/*<br/>*/}
                {/*<label>*/}
                {/*    Month:*/}
                {/*    <input*/}
                {/*        type="text"*/}
                {/*        value={month}*/}
                {/*        onChange={(e) => setMonth(e.target.value)}*/}
                {/*    />*/}
                {/*</label>*/}
                {/*<br/>*/}
                {/*<label>*/}
                {/*    Year:*/}
                {/*    <input*/}
                {/*        type="text"*/}
                {/*        value={year}*/}
                {/*        onChange={(e) => setYear(e.target.value)}*/}
                {/*    />*/}
                {/*</label>*/}
                {/*<br/>*/}
                <label>
                    Date:
                    <DatePicker
                        selected={selectedDate}
                        onChange={(date) => setSelectedDate(date)}
                    />
                </label>
                <br/>
                <button onClick={handleGetAttendance}>Get Attendance</button>

            </div>

            {/* Display the attendance data if available */}
            {/*{aData && (*/}
            {/*    <div>*/}
            {/*    <h3>Attendance Data</h3>*/}
            {/*        /!* Map over each day and render attendance entries *!/*/}
            {/*        {Object.keys(aData).map((id) => (*/}
            {/*            <div key={id}>*/}
            {/*                /!*<p>Id: {id}</p>*!/*/}
            {/*                /!* Map over each attendance entry for the day *!/*/}
            {/*                {aData[id].map((entry, index) => (*/}
            {/*                    <div key={index}>*/}
            {/*                        <p>Id of the student: {entry.id}</p>*/}
            {/*                        <p>First Punch In: {entry.firstPunchIn}</p>*/}
            {/*                        <p>Last Punch Out: {entry.lastPunchOut}</p>*/}
            {/*                    </div>*/}
            {/*                ))}*/}
            {/*            </div>*/}
            {/*        ))}*/}
            {/*    </div>*/}
            {/*)}*/}
            {/*<div className="attendance-container">*/}
            {/*    /!* Map over each day and render attendance entries *!/*/}
            {/*    {Object.keys(aData).map((day) => (*/}
            {/*        <div key={day} className="attendance-card">*/}
            {/*            <h3>Date: {day}</h3>*/}
            {/*            /!* Map over each attendance entry for the day *!/*/}
            {/*            {aData[day].map((entry, index) => (*/}
            {/*                <div key={index}>*/}
            {/*                    <p>Id of the student: {entry.id}</p>*/}
            {/*                    <p>First Punch In: {entry.firstPunchIn}</p>*/}
            {/*                    <p>Last Punch Out: {entry.lastPunchOut}</p>*/}
            {/*                </div>*/}
            {/*            ))}*/}
            {/*        </div>*/}
            {/*    ))}*/}
            {/*</div>*/}
            {aData && (<div>
                <table style={{borderCollapse: 'collapse', width: '100%'}} id="attendance-table">
                    <thead style={{backgroundColor: '#f2f2f2'}}>
                    <tr>
                        <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Date</th>
                        <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Id</th>
                        <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Punch In</th>
                        <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Punch Out</th>
                    </tr>
                    </thead>
                    <tbody>
                    {/* Map over each attendance entry for the day */}
                    {Object.keys(aData).map(day => (
                        aData[day].map((entry, index) => (
                            <tr key={index}>
                                <td style={{
                                    padding: '8px',
                                    border: '1px solid #ddd'
                                }}>{selectedDate.getDate().toString() + "/" + (selectedDate.getMonth() + 1).toString() + "/" + selectedDate.getFullYear().toString()}</td>
                                <td style={{padding: '8px', border: '1px solid #ddd'}}>{entry.id}</td>
                                <td style={{padding: '8px', border: '1px solid #ddd'}}>{entry.firstPunchIn}</td>
                                <td style={{padding: '8px', border: '1px solid #ddd'}}>{entry.lastPunchOut}</td>
                            </tr>
                        ))
                    ))}
                    </tbody>
                </table>
                <button onClick={handleDownloadCSV}>Download CSV</button>
                <button onClick={handleDownloadPDF}>Download PDF</button>
            </div>)}

        </div>
    );
};

export default TeacherClassAttendanceByMonth;
