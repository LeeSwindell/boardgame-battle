import { useEffect, useRef } from 'react'
import { logger } from '../logger/logger';

let socketUrl = import.meta.env.VITE_PROD_SOCKET_API;
const prodMode = import.meta.env.VITE_PROD_MODE;
if (prodMode === 'dev') {
  socketUrl = import.meta.env.VITE_DEV_SOCKET_API;
}

interface UseLobbySocketProps {
    onOpen?: () => void,
    onClose?: () => void,
    onMessage?: (event: MessageEvent) => void,
}

function useLobbySocket({onOpen, onClose, onMessage}) {
    const socket = useRef(null);

    useEffect(() => {
        const username = localStorage.getItem('sessionid');
        // socket.current = new WebSocket(`ws://localhost:8000/connectsocket/${username}`);
        if (socketUrl === 'localhost:8000') {
          socket.current = new WebSocket(`ws://localhost:8000/connectsocket/${username}`);
        } else {
          socket.current = new WebSocket(`wss://${socketUrl}/connectsocket/${username}`);
        }

        socket.current.onopen = () => {
            logger.log('lobby socket opened');
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