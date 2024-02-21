// AddStudent.js
import React, { useState } from 'react';
import axios from 'axios';
import PrincipalSidebar from "./PrincipalSideBar";

const AddClass = () => {
    const [classId, setClassId] = useState(0);
    const [className, setClassName] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [classNames,setClassNames]=useState([])

    const handleAddClass = async () => {
        try {
            const response = await axios.post(
                '/api/principal/Class',
                {
                    class_name: className,
                },
                {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    withCredentials: true,
                }
            );
            console.log(response.data)
            console.log(response.data[0])
            if (response.data===0){
                setErrorMessage('Class already exists');
            }
            // if (response.data.success) {
            //     setSuccessMessage('Class added successfully.');
            //     setErrorMessage('');
            // } e
            else {
                setErrorMessage('');
                setSuccessMessage('Class inserted successfully');
            }
        } catch (error) {
            console.error('Error during adding class:', error);
            setErrorMessage('An error occurred while processing your request.');
            setSuccessMessage('');
        }
    };
    const handleClasses = async () => {
        try {
            const response = await axios.get(
                '/api/principal/Classes',
                {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    // withCredentials: true,
                }
            );
            console.log(response.data)
            setClassNames(response.data)
            console.log(response.data.success)
            // if (response.data.success) {
            //     setSuccessMessage('Class added successfully.');
            //     setErrorMessage('');
            // } else {
            //     setErrorMessage('Failed to add Class. Please try again.');
            //     setSuccessMessage('');
            // }
        } catch (error) {
            console.error('Error during adding class:', error);
            setErrorMessage('An error occurred while processing your request.');
            setSuccessMessage('');
        }
    };
    const [isNavVisible, setNavVisible] = useState(false);

    const handleClick = () => {
        setNavVisible(!isNavVisible);
    };
    return (
        <div>
            <button className="btn-nav" style={{
                background: "none",
                border: "none"
            }} onClick={handleClick}>
                <div className="bar arrow-top-r"></div>
                <div className="bar arrow-middle-r"></div>
                <div className="bar arrow-bottom-r"></div>
            </button>


            <PrincipalSidebar isNavVisible={isNavVisible}/>
            <div style={{display: "flex", flexDirection: "column"}}>
                <h2>Add Class</h2>

                <br/>
                <label>
                    Class Name:
                    <input
                        style={{width: "-webkit-fill-available"}}
                        type="text"
                        value={className}
                        onChange={(e) => setClassName(e.target.value)}
                    />
                </label>
                <br/>
                <button onClick={handleAddClass}>Add Class</button>

                {/* Display success or error message */}
                {successMessage && <p style={{color: 'green'}}>{successMessage}</p>}
                {errorMessage && <p style={{color: 'red'}}>{errorMessage}</p>}
                <button onClick={handleClasses}>View Classes</button>
                <ul>
                    {classNames !== null && classNames.map((classi) => (
                        <li key={classi}>
                            {classi}
                        </li>

                    ))}
                </ul>
            </div>
        </div>
    );
};

export default AddClass;
