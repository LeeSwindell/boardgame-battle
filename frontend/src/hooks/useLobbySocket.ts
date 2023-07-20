import { useEffect, useRef } from 'react'
import { logger } from '../logger/logger';

let socketUrl = import.meta.env.VITE_PROD_SOCKET_API;
const prodMode = import.meta.env.VITE_PROD_MODE;
if (prodMode === 'dev') {
  socketUrl = import.meta.env.VITE_DEV_SOCKET_API;
}
if (prodMode === 'run') {
  socketUrl = import.meta.env.VITE_SOCKET_URL;
}

interface UseLobbySocketProps {
    onOpen?: () => void,
    onClose?: () => void,
    onMessage?: (event: MessageEvent) => void,
}

function useLobbySocket({onOpen, onClose, onMessage}, gameid) {
    const socket = useRef(null);

    useEffect(() => {
        const username = localStorage.getItem('sessionid');
        // socket.current = new WebSocket(`ws://localhost:8000/connectsocket/${username}`);
        if (socketUrl === 'localhost:8000/lm') {
          socket.current = new WebSocket(`ws://${socketUrl}/connectsocket/${username}/${gameid}`);
        } else {
          socket.current = new WebSocket(`wss://${socketUrl}/connectsocket/${username}/${gameid}`);
        }

        socket.current.onopen = () => {
            logger.log('lobby socket opened - id:', gameid);
            onOpen && onOpen();
        }
        socket.current.onclose = () => {
            logger.log('lobby socket closed');
            onClose && onClose();
        }
        socket.current.onmessage = (event: MessageEvent) => {
            onMessage && onMessage(event);
        };
    
        // cleanup socket connection and send a request to backend when leaving page.
        return () => {
          socket.current.close();
        };
      }, []);

    return socket;
}

export default useLobbySocket;