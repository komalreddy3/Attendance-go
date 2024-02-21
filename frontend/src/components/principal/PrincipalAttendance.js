import React, {useEffect, useState} from 'react';
import axios from 'axios';
// import Calendar from 'react-calendar'; // Import the Calendar component
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import { saveAs } from 'file-saver';
import { jsPDF } from 'jspdf';
import 'jspdf-autotable';
import PrincipalSidebar from "./PrincipalSideBar";

const PrincipalTeacherAttendanceByMonth = () => {

    const [teacherID, setTeacherID] = useState('');
    // const [month, setMonth] = useState('');
    // const [year, setYear] = useState('');
    const [aData, setAttendanceData] = useState('');
    const [selectedDate, setSelectedDate] = useState(new Date()); // New state for selected date

    // Function to handle date change
    // const handleDateChange = (date) => {
    //     setSelectedDate(date);
    //     setMonth((date.getMonth() + 1).toString()); // Update month based on selected date
    //     setYear(date.getFullYear().toString()); // Update year based on selected date
    // };
    const handleGetAttendance = async () => {
        try {
            const response = await axios.post('/api/principal/teacherAttendance', {
                id: teacherID,
                // month: parseInt(month),
                // year: parseInt(year),
                month:(parseInt((selectedDate.getMonth() + 1).toString())),
                year:(parseInt((selectedDate.getFullYear().toString())))
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
        const headers = ["Date", "Class", "Punch In", "Punch Out"];
        csvRows.push(headers.join(','));

        Object.keys(aData).forEach(day => {
            aData[day].forEach(entry => {
                const rowData = [
                    selectedDate.getDate() + "/" + (selectedDate.getMonth() + 1) + "/" + selectedDate.getFullYear(),
                    entry.class,
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
        saveAs(blob, teacherID+ '-TeacherAttendance.csv');
    };

    // Function to handle PDF download
    const handleDownloadPDF = () => {
        const doc = new jsPDF();
        doc.autoTable({ html: '#attendance-table' });
        doc.save(teacherID+'-TeacherAttendance.pdf');
    };

    const [isNavVisible, setNavVisible] = useState(false);

    const handleClick = () => {
        setNavVisible(!isNavVisible);
    };
    return (

        <div>

            <div>
                <h2> Teacher Attendance by Month</h2>
                <label>
                    Teacher ID:
                    <input
                        type="text"
                        value={teacherID}
                        onChange={(e) => setTeacherID(e.target.value)}
                    />
                </label>
                <br/>
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
                {/*<label>*/}
                {/*    Select Date:*/}
                {/*    <Calendar onChange={handleDateChange} value={selectedDate}/>*/}
                {/*</label>*/}
                {/*<br/>*/}
                <label>
                    Month and Year:
                    <DatePicker
                        selected={selectedDate}
                        onChange={(date) => setSelectedDate(date)}
                        dateFormat="MM/yyyy"
                        showMonthYearPicker
                    />
                </label>
                <br/>
                <button onClick={handleGetAttendance}>Get Attendance</button>
            </div>

            {/* Display the attendance data if available */}
            {/*{aData && (*/}
            {/*    <div>*/}
            {/*        <h3>Attendance Data</h3>*/}
            {/*        /!* Map over each day and render attendance entries *!/*/}
            {/*        {Object.keys(aData).map((day) => (*/}
            {/*            <div key={day}>*/}
            {/*            <p>Date: {day}</p>*/}
            {/*                /!* Map over each attendance entry for the day *!/*/}
            {/*                {aData[day].map((entry, index) => (*/}
            {/*                    <div key={index}>*/}
            {/*                        <p>Class: {entry.class}</p>*/}
            {/*                        <p>First Punch In: {entry.firstPunchIn}</p>*/}
            {/*                        <p>Last Punch Out: {entry.lastPunchOut}</p>*/}
            {/*                    </div>*/}
            {/*                ))}*/}
            {/*            </div>*/}
            {/*        ))}*/}
            {/*    </div>*/}
            {/*)}*/}
            <div>
                {/* Map over each day and render attendance entries */}
                {Object.keys(aData).map((day) => (
                    <div key={day}>
                        {/*<h3>Date: {day}</h3>*/}
                        {/*/!* Map over each attendance entry for the day *!/*/}
                        {/*{aData[day].map((entry, index) => (*/}
                        {/*    <div key={index}>*/}
                        {/*        <p>Class: {entry.class}</p>*/}
                        {/*        <p>First Punch In: {entry.firstPunchIn}</p>*/}
                        {/*        <p>Last Punch Out: {entry.lastPunchOut}</p>*/}
                        {/*    </div>*/}
                        {/*))}*/}
                        <table style={{borderCollapse: 'collapse', width: '100%'}} id="attendance-table">
                            <thead style={{backgroundColor: '#f2f2f2'}}>
                            <tr>
                                <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Date</th>
                                <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Class</th>
                                <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Punch In
                                </th>
                                <th style={{padding: '8px', border: '1px solid #ddd', textAlign: 'left'}}>Punch
                                    Out
                                </th>
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
                                        }}>{day + "/" + (selectedDate.getMonth() + 1).toString() + "/" + selectedDate.getFullYear().toString()}</td>
                                        <td style={{padding: '8px', border: '1px solid #ddd'}}>{entry.class}</td>
                                        <td style={{
                                            padding: '8px',
                                            border: '1px solid #ddd'
                                        }}>{entry.firstPunchIn}</td>
                                        <td style={{
                                            padding: '8px',
                                            border: '1px solid #ddd'
                                        }}>{entry.lastPunchOut}</td>
                                    </tr>
                                ))
                            ))}
                            </tbody>
                        </table>
                        <button onClick={handleDownloadCSV}>Download CSV</button>
                        <button onClick={handleDownloadPDF}>Download PDF</button>
                    </div>
                ))}
            </div>
        </div>


    );
};

export default PrincipalTeacherAttendanceByMonth;
