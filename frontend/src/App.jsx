import DiscardPile from "./components/discard"
import EndTurn from "./components/endturn"
import Gameboard from "./components/gameboard"
import Hand from "./components/hand"
import MarketCard from "./components/marketcard"
import PlayArea from "./components/playarea"
import PlayerInfo from "./components/playerinfo"

function App() {
  return (
    <>
      <div className="absolute w-full h-full border invisible">
        <div className="flex justify-center">
          <MarketCard />
        </div>
      </div>
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
          <DiscardPile />
        </div>
      </div>
    </>

  )
}

export default App