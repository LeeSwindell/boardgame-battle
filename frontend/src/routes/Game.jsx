import {
  useState, createContext, useContext, useEffect, useRef,
} from 'react';
import CardContainer from '../components/CardContainer';
import DiscardPile from '../components/Discard';
import EndTurn from '../components/Endturn';
import Gameboard from '../components/Gameboard';
import Hand from '../components/Hand';
import InspectCard from '../components/Inspectcard';
import PlayArea from '../components/Playarea';
import PlayerInfo from '../components/Playerinfo';
// import { Route, Link, BrowserRouter as Router } from "react-router-dom"
import { api, gameapi } from '../api';

// use react context to pass inspect around
const InspectContext = createContext();

function InspectProvider({ children }) {
  const [inspectCard, setInspectCard] = useState();
  const value = { inspectCard, setInspectCard };
  return <InspectContext.Provider value={value}>{children}</InspectContext.Provider>;
}

function useInspect() {
  const context = useContext(InspectContext);
  if (context === undefined) {
    throw new Error('useInspect must be used within an InspectProvider');
  }
  return context;
}

export { useInspect };

const GamestateContext = createContext();

function GamestateProvider({ children }) {
  const socket = useRef(null);
  const [gamestate, setGamestate] = useState({});

  useEffect(() => {
    socket.current = new WebSocket('ws://localhost:8000/connectsocket');
    socket.current.onopen = () => console.log('lobby socket opened');
    socket.current.onclose = () => console.log('lobby socket closed');

    socket.current.onmessage = (event) => {
      const data = JSON.parse(event.data);
      switch (data.type) {
        case 'Gamestate':
          setGamestate(data.data);
          console.log(data.data);
          console.log(gamestate);
          break;
        default:
          break;
      }
    };

    // cleanup socket connection and send a request to backend when leaving page.
    return () => {
      socket.current.close();
    };
  }, []);

  return (
    <GamestateContext.Provider value={gamestate}>
      {children}
    </GamestateContext.Provider>
  );
}

function Game() {
  // // create socket, read values from gamestate to set health.
  // const socket = useRef(null);
  // const [gs, setGs] = useState(null);

  // useEffect(() => {
  //   socket.current = new WebSocket('ws://localhost:8000/connectsocket');
  //   socket.current.onopen = () => console.log('lobby socket opened');
  //   socket.current.onclose = () => console.log('lobby socket closed');

  //   socket.current.onmessage = (event) => {
  //     const data = JSON.parse(event.data);
  //     switch (data.type) {
  //       case 'Gamestate':
  //         setGs(data.data);
  //         console.log(data.data);
  //         console.log(gs);
  //         break;
  //       default:
  //         break;
  //     }
  //   };

  //   // cleanup socket connection and send a request to backend when leaving page.
  //   return () => {
  //     socket.current.close();
  //   };
  // }, []);

  return (
    <GamestateProvider>
      <InspectProvider>
        <InspectCard />
        <div className="flex flex-row justify-between">
          <div className="flex flex-col space-y-4 w-auto h-auto">
            <Gameboard />
            <PlayArea />
            <Hand />
          </div>
          <div className="flex flex-col border justify-between items-center">
            <div>
              <PlayerInfo />
              <PlayerInfo />
              <PlayerInfo />
              <PlayerInfo />
              <EndTurn />
            </div>
            <div className="flex flex-col">
              <div className="text-center">Discard Pile</div>
              <CardContainer card={<DiscardPile />} size="reg" extra="p-2 m-2" />
            </div>
          </div>
        </div>
      </InspectProvider>
    </GamestateProvider>
  );
}

export default Game;
