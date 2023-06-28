import { useEffect, useRef } from 'react';

function TestServer() {
  const socketUrl = 'www.gamewithyourfriends.dev/lm';
  const socket = useRef(null);
  // Create the socket connection
  useEffect(() => {
    console.log('testing socket hook');
    console.log(`wss://${socketUrl}/connectsocket`);
    socket.current = new WebSocket(`wss://${socketUrl}/connectsocket`);

    socket.current.onopen = () => {
      console.log('lobby socket opened');
    };

    socket.current.onclose = () => console.log('lobby socket closed');
  }, []);

  return (
    <div>Success???</div>
  );
}

export default TestServer;
