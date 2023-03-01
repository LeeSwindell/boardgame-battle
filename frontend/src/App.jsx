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
      <div className="fixed w-full h-full border border-blue-500 visible">
        <div className="flex w-full h-full border border-green-500 justify-center items-center">
          <div className="border border-pink-500 w-64 h-64">
            <MarketCard />
          </div>
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