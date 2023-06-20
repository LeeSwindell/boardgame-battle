import {
  useState, createContext, useContext, useEffect, useRef, useMemo,
} from 'react';
import { useNavigate, useParams } from 'react-router-dom';
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
import EndGame from '../components/Endgame';

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
  const { gameid } = useParams();
  const socket = useLobbySocket({
    onMessage: (event) => {
      const data = JSON.parse(event.data);
      switch (data.type) {
        case 'Gamestate':
          setGamestate(data.data);
          break;
        case 'UserInput':
          setUserInput({
            inputs: data.data,
            description: data.description,
            messageID: data.id,
          });
          logger.log('user input:::::', data);
          break;
        default:
          break;
      }
    },
  });
  const value = useMemo(
    () => ({
      gamestate, setGamestate, socket, userInput, setUserInput, gameid,
    }),
    [gamestate, socket, userInput, gameid],
  );
  const navigate = useNavigate();

  useEffect(() => {
    gameapi
      .get(`/${gameid}/getgamestate`)
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
  const {
    gamestate, userInput, setUserInput, gameid,
  } = useGamestate();

  function SubmitUserChoice(choice) {
    return (
      () => {
        const id = userInput.messageID;
        logger.log('sending user choice: ', id);
        api
          .post(`/game/${gameid}/submituserchoice`, { choice, id })
          .then(() => {
            setUserInput((currentInput) => {
              if (currentInput.messageID === id) {
                return null;
              }
              return currentInput;
            });
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
                <p className="p-2 w-full text-center font-bold">
                  {userInput.description}
                </p>
                {userInput.inputs.map((option, i) => <button key={option + i} type="submit" className="p-2 m-2 border rounded bg-blue-500 hover:bg-blue-700 text-white font-bold" onClick={SubmitUserChoice(option)}>{option}</button>)}
              </div>
            </div>
          </div>
          )
        }
        <InspectCard />
        <div className="flex flex-row justify-between">
          <div>
            <Gameboard />
            <div className="fixed ml-4 top-64 left-72 border w-[49rem] -z-10">
              <PlayArea />
            </div>
            <Hand />
          </div>
          <div className="border justify-between items-center fixed top-0 right-0">
            <div>
              {
                Object.entries(gamestate.players)
                  .map(([username]) => (<PlayerInfo key={username} username={username} />))
              }
              <EndTurn />
              {/* <EndGame /> */}
            </div>
            <div className="flex flex-col items-center">
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
