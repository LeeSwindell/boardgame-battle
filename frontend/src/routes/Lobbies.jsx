import { useEffect, useState } from "react";
import axios from 'axios'

const baseUrl = 'localhost:8000'

function Lobbies() {
    const [lobbies, setLobbies] = useState([])
    useEffect(async () => {
        const result = await axios(
          baseUrl + '/lobbies',
        );
    
        setLobbies(result.data);
      }, []);


    return (
        <div>List of lobbies</div>
    )
}

export default Lobbies;