import React, { useEffect, useState } from 'react';
import axios from 'axios';

const AddStudent = () => {
    const [studentID, setStudentID] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [selectedClass, setSelectedClass] = useState('');
    const [classNames, setClassNames] = useState([]);

    useEffect(() => {
        const fetchClasses = async () => {
            try {
                const response = await axios.get('/api/principal/Classes', {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });
               // if (response.data.success) {
                    const newClassNames = response.data;
                    if (!isEqualArrays(newClassNames, classNames)) {
                        setClassNames(newClassNames);
                        // }
                        // } else {
                        //     setErrorMessage('Failed to fetch classes. Please try again.');
                        // }
                    }
                    } catch (error) {
                console.error('Error fetching classes:', error);
                setErrorMessage('An error occurred while fetching classes.');
            }
        };

        fetchClasses();

        const isEqualArrays = (array1, array2) => {
            if(array1==null || array2===null) return true;
            if (array1.length !== array2.length) {
                return false;
            }
            for (let i = 0; i < array1.length; i++) {
                if (array1[i] !== array2[i]) {
                    return false;
                }
            }
            return true;
        };
    }, [classNames]); // Only re-run the effect if classNames changes

    const handleAddStudent = async () => {
        if (studentID===''){
            setErrorMessage('Student ID cant be null')
            return
        }
        if(selectedClass===''){
            setErrorMessage('Please select the class to add student')
            return
        }
        try {
            const response = await axios.post(
                '/api/principal/addStudent',
                {
                    studentID: studentID,
                    class_name: selectedClass,
                },
                {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    withCredentials: true,
                }
            );

            if (response.data.success) {
                setSuccessMessage('Student added successfully.');
                setErrorMessage('');
            } else {
                setErrorMessage('Failed to add student. Please try again.');
                setSuccessMessage('');
            }
        } catch (error) {
            console.error('Error during adding student:', error);
            setErrorMessage('An error occurred while processing your request.');
            setSuccessMessage('');
        }
    };

    return (
        <div style={{ display: 'flex', flexDirection: 'column' }}>
            <h2>Add Student</h2>
            <label>
                Student ID:
                <input
                    style={{ width: '-webkit-fill-available' }}
                    type="text"
                    value={studentID}
                    onChange={(e) => setStudentID(e.target.value)}
                />
            </label>
            <br />
            <label>
                Class:
                <div style={{display:"flex",flexDirection:"column"}}>{classNames !== null &&
                    classNames.map((className, index) => (
                        <div key={index} style={{ display: 'flex', columnGap: '10px' }}>
                            <input
                                type="radio"
                                name="className"
                                value={className}
                                checked={selectedClass === className}
                                onChange={(e) => setSelectedClass(e.target.value)}
                            />
                            {className}
                        </div>
                    ))}</div>
            </label>
            <br />

            <button onClick={handleAddStudent}>Add Student</button>

            {/* Display success or error message */}
            {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
        </div>
    );
};

export default AddStudent;
