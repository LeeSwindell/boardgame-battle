function createSocket() {
  const socket = new WebSocket('ws://localhost:8000/connectsocket');

  return socket;
}

export default createSocket;
