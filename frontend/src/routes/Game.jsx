import {
  useState, createContext, useContext, useEffect, useRef, useMemo,
} from 'react';
import { useNavigate } from 'react-router-dom';
import CardContainer from '../components/CardContainer';
import DiscardPile from '../components/Discard';
import EndTurn from '../components/Endturn';
import Gameboard from '../components/Gameboard';
import Hand from '../components/Hand';
import InspectCard from '../components/Inspectcard';
import PlayArea from '../components/Playarea';
import PlayerInfo from '../components/Playerinfo';
// import { Route, Link, BrowserRouter as Router } from "react-router-dom"
import { gameapi } from '../api';

// use react context to pass inspect around

// FIX! EXTRACT PROVIDERS TO A COMPONENT, ONE CONTEXT PER PROVIDER

const InspectContext = createContext();

function InspectProvider({ children }) {
  const [inspectCard, setInspectCard] = useState();
  const value = useMemo(() => ({ inspectCard, setInspectCard }), [inspectCard]);

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
  const [gamestate, setGamestate] = useState();
  const value = useMemo(() => ({ gamestate, setGamestate, socket }), [gamestate, socket]);
  const navigate = useNavigate();

  useEffect(() => {
    gameapi
      .get('/0/getgamestate')
      .then((response) => {
        setGamestate(response.data);
      })
      .catch(() => {
        console.log('error');
        navigate('/');
      });
  }, []);

  // just to see how gamestate changes.
  useEffect(() => {
    console.log(gamestate);
  }, [gamestate]);

  useEffect(() => {
    socket.current = new WebSocket('ws://localhost:8000/connectsocket');
    socket.current.onopen = () => console.log('lobby socket opened');
    socket.current.onclose = () => console.log('lobby socket closed');

    socket.current.onmessage = (event) => {
      const data = JSON.parse(event.data);
      switch (data.type) {
        case 'Gamestate':
          setGamestate(data.data);
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
    <GamestateContext.Provider value={value}>
      {children}
    </GamestateContext.Provider>
  );
}

function useGamestate() {
  const context = useContext(GamestateContext);
  if (context === undefined) {
    throw new Error('useGamestate must be used within GamestateProvider');
  }
  return context;
}

export { useGamestate };

function GameWithState() {
  const { gamestate } = useGamestate();
  if (gamestate) {
    return (
      <>
        <InspectCard />
        <div className="flex flex-row justify-between">
          <div className="flex flex-col space-y-4 w-auto h-auto">
            <Gameboard />
            <PlayArea />
            <Hand />
          </div>
          <div className="flex flex-col border justify-between items-center">
            <div>
              {
                Object.entries(gamestate.players)
                  .map(([username]) => (<PlayerInfo key={username} username={username} />))
              }
              <EndTurn />
            </div>
            <div className="flex flex-col">
              <div className="text-center">Discard Pile</div>
              <CardContainer card={<DiscardPile />} size="reg" extra="p-2 m-2" />
            </div>
          </div>
        </div>
      </>
    );
  }
}

function Game() {
  return (
    <GamestateProvider>
      <InspectProvider>
        <GameWithState />
      </InspectProvider>
    </GamestateProvider>
  );
}

export default Game;
