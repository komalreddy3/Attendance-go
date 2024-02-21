// TeacherPunchIn.js

import React, {useEffect, useState} from 'react';
import axios from 'axios';

const TeacherPunchIn = () => {
    // const [userID, setUserID] = useState('');
    const [className, setClass] = useState('');
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
    const handlePunchIn = async () => {
        if(selectedClass===''){
            setErrorMessage('Please select the class you want to punch in')
            return
        }
        try {
            const response = await axios.post('/api/teacher/punchIn', {
                // user_id: userID,
                // class: className,
                class:selectedClass,
            }, {
                headers: {
                    'Content-Type': 'application/json',
                },
                withCredentials: true,
            });

            if (response.data.success) {
                // Handle success, if needed
                setErrorMessage('');
                setSuccessMessage('Punch-in successful');
                console.log('Punch-in successful');
            } else {
                setErrorMessage('Failed to punch in');
                setSuccessMessage('');
                console.error('Failed to punch in');
            }
        } catch (error) {
            console.error('Error during punch-in:', error);
        }
    };

    return (
        <div style={{display: "flex", flexDirection: "column"}}>
            <h2>Teacher Punch In</h2>
            {/*<label>*/}
            {/*    User ID:*/}
            {/*    <input*/}
            {/*        type="text"*/}
            {/*        value={userID}*/}
            {/*        onChange={(e) => setUserID(e.target.value)}*/}
            {/*    />*/}
            {/*</label>*/}
            {/*<br />*/}
            {/*<label>*/}
            {/*    Class:*/}
            {/*    <input*/}
            {/*        type="text"*/}
            {/*        value={className}*/}
            {/*        onChange={(e) => setClass(e.target.value)}*/}
            {/*    />*/}
            {/*</label>*/}
            {/*<br />*/}
            <label>
                Class:
                <div style={{display: "flex", flexDirection: "column"}}>{classNames !== null &&
                    classNames.map((className, index) => (
                        <div key={index} style={{display: 'flex', columnGap: '10px'}}>
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
            <br/>
            <button onClick={handlePunchIn}>Punch In</button>
            {/* Display success or error message */}
            {successMessage && <p style={{color: 'green'}}>{successMessage}</p>}
            {errorMessage && <p style={{color: 'red'}}>{errorMessage}</p>}
        </div>
    );
};

export default TeacherPunchIn;
