// src/components/Login.js
import React, {useEffect, useState} from 'react';
import axios from 'axios';
import "./login.css"
// Set up default Axios configuration
// axios.defaults.baseURL = 'http://localhost:9000';
// axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    useEffect(() => {
        //window.history.pushState(null, document.title, window.location.href);
        //window.history.replaceState(null, document.title, `/login`);
        //window.history.replaceState(null, document.title, `/login`);
        const handlePopstate = (event) => {
            window.history.replaceState(null, document.title, `/login`);
            // window.location.replace(`/login`);
            // window.location.replace(`/login`)
           // window.history.replaceState(null, document.title, `/login`);
            window.history.go(1)
        };
        window.addEventListener('popstate', handlePopstate);

        return () => {
            window.removeEventListener('popstate', handlePopstate);
        };
    }, []);
    const handleLogin = async (e) => {
        e.preventDefault();
        const backendServiceName = process.env.react_app_addr;
        console.log(backendServiceName)


        try {
            // const response = await axios.post('http://localhost:8080/login', {
            //'http://localhost:9000/login'
           // const url =`http://${backendServiceName}:9000`
            //http://localhost:9000/login
            const response = await axios.post("/api/login", {
                username,
                password,
                role,
            }, {
                headers: {
                    'Content-Type': 'application/json',
                },
                withCredentials: true,
            });
            // const response = await fetch( 'http://'+backendServiceName+':9000/login', {
            //     method:"POST"
            // });
            // // const response = await fetch('http://api:8080/login',{method: "POST"});
            console.log(response)
            if (response.data.success) {
                // Assuming the backend returns a success flag
                //const { token } = response.headers;
                // Store the token in localStorage or a secure cookie for subsequent requests

                // Redirect to the Dashboard with the user's role

                window.location.href = `/dashboard/${role}`;
            } else {
                setErrorMessage('Invalid login credentials');
                setSuccessMessage('');
                console.error('Invalid login credentials');
            }

        } catch (error) {
            console.error('Error during login:', error);
        }
    };

    return (
        <div>
            <h1>Login Page</h1>
            <form onSubmit={handleLogin}>
                <label>
                    Username:
                    <input
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </label>
                <br/>
                <label>
                    Password:
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </label>
                <br/>
                {/*<label>*/}
                {/*    Role:*/}
                {/*    <input*/}
                {/*        type="text"*/}
                {/*        placeholder="principal..teacher..student"*/}
                {/*        value={role}*/}
                {/*        onChange={(e) => setRole(e.target.value)}*/}
                {/*    />*/}
                {/*</label>*/}
                <label>
                    Role:
                    <select
                        value={role}
                        onChange={(e) => setRole(e.target.value)}
                    >
                        <option value="">Select a role</option>
                        <option value="principal">Principal</option>
                        <option value="teacher">Teacher</option>
                        <option value="student">Student</option>
                    </select>
                </label>
                <br/>
                <button type="submit">Login</button>
            </form>
            {/* Display success or error message */}
            {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
        </div>
    );
};

export default Login;
