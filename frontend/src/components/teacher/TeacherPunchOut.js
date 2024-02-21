// TeacherPunchOut.js

import React, { useState } from 'react';
import axios from 'axios';

const TeacherPunchOut = () => {
    // const [userID, setUserID] = useState('');
    const [className, setClass] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handlePunchOut = async () => {
        try {
            const response = await axios.post('/api/teacher/punchOut', {
                // user_id: userID,
                class: className,
            }, {
                headers: {
                    'Content-Type': 'application/json',
                },
                withCredentials: true,
            });

            if (response.data.success) {
                // Handle success, if needed
                setErrorMessage('');
                setSuccessMessage('Punch-out successful');
                console.log('Punch-out successful');
            } else {
                setErrorMessage('Failed to punch out');
                setSuccessMessage('');
                console.error('Failed to punch out');
            }
        } catch (error) {
            console.error('Error during punch-out:', error);
        }
    };

    return (
        <div style={{display:"flex",flexDirection:"column"}}>
            <h2>Teacher Punch Out</h2>
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
            <button onClick={handlePunchOut}>Punch Out</button>
            {/* Display success or error message */}
            {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
        </div>
    );
};

export default TeacherPunchOut;
