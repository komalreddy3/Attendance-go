// StudentPunchIn.js

import React, {useEffect, useState} from 'react';
import axios from 'axios';

const StudentPunchIn = () => {
    // const [userID, setUserID] = useState('');
    const [className, setClass] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    useEffect(()=> {
        const interval = setInterval(() => {
            const RemovePunchIn = async () => {
                try {
                    const response = await axios.post('/api/student/removePunchIn', {
                        // user_id: userID,
                        class: className,
                    }, {
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        withCredentials: true,
                    });

                    if (response.data.allow === false) {
                        setErrorMessage('Please Punchout to Punchin');
                        return
                    }
                    else{
                        setErrorMessage('')
                        setSuccessMessage('')
                        return
                    }

                } catch (error) {
                    console.error('Error during check punch-in:', error);
                }

            }
            RemovePunchIn()
        }, 1000);
        return () => clearInterval(interval);
    })
    const handlePunchIn = async () => {
        try {
            const response = await axios.post('/api/student/punchIn', {
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
                setSuccessMessage('Punch-in successful');
                console.log('Punch-in successful');
            }
            else {
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
            <h2>Student Punch In</h2>
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
            {errorMessage === 'Please Punchout to Punchin' &&
                <button onClick={handlePunchIn} disabled={true} style={{display: "none"}}>Punch In</button>}
            {errorMessage !== 'Please Punchout to Punchin' && <button onClick={handlePunchIn} >Punch In</button>}
            {/* Display success or error message */}
            {successMessage && <p style={{color: 'green'}}>{successMessage}</p>}
            {errorMessage && <p style={{color: 'red'}}>{errorMessage}</p>}
        </div>
    );
};

export default StudentPunchIn;
