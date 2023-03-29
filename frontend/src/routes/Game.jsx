import { useState, createContext, useContext } from 'react';
import CardContainer from '../components/CardContainer';
import DiscardPile from '../components/Discard';
import EndTurn from '../components/Endturn';
import Gameboard from '../components/Gameboard';
import Hand from '../components/Hand';
import InspectCard from '../components/Inspectcard';
import PlayArea from '../components/Playarea';
import PlayerInfo from '../components/Playerinfo';
// import { Route, Link, BrowserRouter as Router } from "react-router-dom"

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

function Game() {
  return (
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
  );
}

export default Game;
