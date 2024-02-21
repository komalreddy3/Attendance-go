import React, {useEffect} from 'react';
import axios from 'axios';



const Logout = () => {
    useEffect(() => {
      // window.history.pushState(null, document.title,`/login`);
        //window.history.replaceState(null, document.title, `/login`);
        const handlePopstate = (event) => {
           window.history.replaceState(null, document.title, `/login`);
           // //  window.location.replace(`/login`);
           //  window.location.replace(`/login`);
            window.history.go(1)
            //window.history.replaceState(null, document.title, `/login`);
        };
        window.addEventListener('popstate', handlePopstate);

        return () => {
            window.removeEventListener('popstate', handlePopstate);
        };
    }, []);
    const handleLogout = async () => {

        try {
           const response= await axios.post('/api/logout', {}, {
               headers: {
                   'Content-Type': 'application/json',
               },
               withCredentials: true,
           });
           console.log(response)
            //window.location.href = `/login`;
            // window.history.replaceState({}, '', `/login`);
            // window.location.replace(`/login`)

            // let history = usehistory();
            // history.replace("/login");
            sessionStorage.clear();
           localStorage.clear();
            window.location.href = `/login`;
            // window.history.pushState(null, '', `/login`);
            // window.location.replace(`/login`);
            if (response.data.success) {
                // Assuming the backend returns a success flag
                //const { token } = response.headers;
                // Store the token in localStorage or a secure cookie for subsequent requests

                // Redirect to the Dashboard with the user's role
                window.location.href = `/login`;
               // window.history.pushState(null, '', `/login`);
                //window.location.replace(`/login`);

            } else {

                console.error('logout error');
            }
            // Redirect or perform any necessary actions after successful logout
        } catch (error) {
            console.error('Error during logout:', error);
        }
    };

    return (
        <div >
            {/*<h2>Logout</h2>*/}
            <button onClick={handleLogout} style={{width: "189px",background:"none",border:"none"}}>Logout</button>

        </div>
    );
};

export default Logout;
