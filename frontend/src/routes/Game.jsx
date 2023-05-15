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
import { api, gameapi } from '../api';
import { logger } from '../logger/logger';
import useLobbySocket from '../hooks/useLobbySocket';

// use react context to pass inspect around
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
  const [gamestate, setGamestate] = useState();
  const [userInput, setUserInput] = useState();
  const socket = useLobbySocket({
    onMessage: (event) => {
      const data = JSON.parse(event.data);
      switch (data.type) {
        case 'Gamestate':
          setGamestate(data.data);
          break;
        case 'UserInput':
          logger.log(userInput);
          setUserInput(data.data);
          break;
        default:
          break;
      }
    },
  });
  const value = useMemo(
    () => ({
      gamestate, setGamestate, socket, userInput, setUserInput,
    }),
    [gamestate, socket, userInput],
  );
  const navigate = useNavigate();

  useEffect(() => {
    gameapi
      .get('/0/getgamestate')
      .then((response) => {
        setGamestate(response.data);
      })
      .catch(() => {
        navigate('/');
      });
  }, []);

  // just to see how gamestate changes.
  useEffect(() => {
    logger.log(gamestate);
  }, [gamestate]);

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
  const { gamestate, userInput, setUserInput } = useGamestate();

  // FIX user game lobby id
  function SubmitUserChoice(choice) {
    return (
      () => {
        api
          .post('/game/0/submituserchoice', { choice })
          .then(() => {
            setUserInput(null);
          })
          .catch((res) => {
            logger.error(res);
          });
      }
    );
  }

  if (gamestate) {
    return (
      <>
        {
          userInput
          && (
          <div className="fixed w-full h-full backdrop-contrast-50">
            <div className="flex w-full h-full justify-center items-center">
              <div className="border bg-white z-50 shadow-2xl">
                <p className="p-2 w-full text-center font-bold">Choose One</p>
                {userInput.map((option, i) => <button key={option + i} type="submit" className="p-2 m-2 border rounded bg-blue-500 hover:bg-blue-700 text-white font-bold" onClick={SubmitUserChoice(option)}>{option}</button>)}
              </div>
            </div>
          </div>
          )
        }
        <InspectCard />
        <div className="flex flex-row justify-between">
          <div className="flex flex-col space-y-1 p-1 w-auto h-auto">
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
              <DiscardPile />
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
