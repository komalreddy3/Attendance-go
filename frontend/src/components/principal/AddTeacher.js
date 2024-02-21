// AddStudent.js
import React, {useEffect, useState} from 'react';
import axios from 'axios';

const AddTeacher = () => {
    const [teacherID, setTeacherID] = useState('');
    //const [className, setClassName] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [classNames,setClassNames]=useState([])
    const [selectedClasses, setSelectedClasses] = useState([]);
    const handleAddTeacher = async () => {
        // Split the className string into an array
        // const classNamesArray = className.split(',').map(className => className.trim());
        // console.log(classNamesArray)
        // console.log(className)
        if (teacherID===''){
            setErrorMessage('Teacher ID cant be null')
            return
        }
        if(selectedClasses===null){
            setErrorMessage('Please select class/es to add teacher')
            return
        }
        try {
            const response = await axios.post(
                '/api/principal/addTeacher',
                {
                    teacherID: teacherID,
                    class_names: selectedClasses,
                },
                {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    withCredentials: true,
                }
            );

            if (response.data.success) {
                setSuccessMessage('Teacher added successfully.');
                setErrorMessage('');
            } else {
                setErrorMessage('Failed to add Teacher. Please try again.');
                setSuccessMessage('');
            }
        } catch (error) {
            console.error('Error during adding student:', error);
            setErrorMessage('An error occurred while processing your request.');
            setSuccessMessage('');
        }
    };
    // useEffect(()=>{
    //     const interval = setInterval(() => {  const handleClasses = async () => {
    //         try {
    //             const response = await axios.get(
    //                 '/api/principal/Classes',
    //                 {
    //                     headers: {
    //                         'Content-Type': 'application/json',
    //                     },
    //                     // withCredentials: true,
    //                 }
    //             );
    //             console.log(response.data)
    //             setClassNames(response.data)
    //             //console.log(response.data.success)
    //             // if (response.data.success) {
    //             //     setSuccessMessage('Class added successfully.');
    //             //     setErrorMessage('');
    //             // } else {
    //             //     setErrorMessage('Failed to add Class. Please try again.');
    //             //     setSuccessMessage('');
    //             // }
    //         } catch (error) {
    //             console.error('Error during adding class:', error);
    //             setErrorMessage('An error occurred while processing your request.');
    //             setSuccessMessage('');
    //         }
    //     };
    //     handleClasses()
    // }, 1000);
    // },[])
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
    const handleCheckboxChange = (event) => {
        const value = event.target.value;
        if (event.target.checked) {
            setSelectedClasses([...selectedClasses, value]);
        } else {
            setSelectedClasses(selectedClasses.filter((className) => className !== value));
        }
    };
    return (
        <div style={{display: "flex", flexDirection: "column"}}>
            <h2>Add Teacher</h2>
            <label>
                Teacher ID:
                <input
                    style={{width: "-webkit-fill-available"}}
                    type="text"
                    value={teacherID}
                    onChange={(e) => setTeacherID(e.target.value)}
                />
            </label>
            <br/>
            {/*<label>*/}
            {/*    Class Names (seperated by comma):*/}
            {/*    <input*/}
            {/*        style={{width: "-webkit-fill-available"}}*/}
            {/*        type="text"*/}
            {/*        value={className}*/}
            {/*        onChange={(e) => setClassName(e.target.value)}*/}
            {/*    />*/}
            {/*</label>*/}
            {/*<br />*/}
            <label>
                Class/es:
                {/*<select*/}
                {/*    style={{width: "-webkit-fill-available"}}*/}
                {/*    multiple*/}
                {/*    value={selectedClasses}*/}
                {/*    onChange={(e) => setSelectedClasses(Array.from(e.target.selectedOptions, option => option.value))}*/}
                {/*>*/}
                {/*    {classNames!==null && classNames.map((className, index) => (*/}
                {/*        <option key={index} value={className}>{className}</option>*/}
                {/*    ))}*/}
                {/*</select>*/}
                <div style={{display:"flex",flexDirection:"column"}}>{classNames !== null && classNames.map((className, index) => (
                    <div key={index} style={{    display: "flex",
                        columnGap: "10px"}}>
                        <input
                            type="checkbox"
                            value={className}
                            checked={selectedClasses.includes(className)}
                            onChange={handleCheckboxChange}
                        />
                        {className}
                    </div>
                ))}</div>
            </label>
            <br/>
            <button onClick={handleAddTeacher}>Add Teacher</button>

            {/* Display success or error message */}
            {successMessage && <p style={{color: 'green'}}>{successMessage}</p>}
            {errorMessage && <p style={{color: 'red'}}>{errorMessage}</p>}
        </div>
    );
};

export default AddTeacher;
