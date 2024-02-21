// StudentPunchOut.js

import React, {useEffect, useState} from 'react';
import axios from 'axios';

const StudentPunchOut = () => {
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

                    if (response.data.allow === true) {
                        setErrorMessage('Please Punchin to Punchout');
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
    const handlePunchOut = async () => {
        try {
            const response = await axios.post('/api/student/punchOut', {
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
            <h2>Student Punch Out</h2>
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
            {/*<button onClick={handlePunchOut}>Punch Out</button>*/}
            {errorMessage === 'Please Punchin to Punchout' &&
                <button onClick={handlePunchOut} disabled={true} style={{display: "none"}}>Punch Out</button>}
            {errorMessage !== 'Please Punchin to Punchout' && <button onClick={handlePunchOut} >Punch Out</button>}
            {/* Display success or error message */}
            {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
        </div>
    );
};

export default StudentPunchOut;
