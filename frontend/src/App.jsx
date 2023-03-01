import EndTurn from "./components/endturn"
import Gameboard from "./components/gameboard"
import Hand from "./components/hand"
import PlayArea from "./components/playarea"
import PlayerInfo from "./components/playerinfo"

function App() {
  return (
    <div className="flex flex-row justify-between">
      <div className="flex flex-col space-y-4 w-auto h-auto">
        <Gameboard />
        <PlayArea />
        <Hand />
      </div>
      <div className="flex flex-col border">
        <PlayerInfo />
        <PlayerInfo />
        <PlayerInfo />
        <PlayerInfo />
        <EndTurn />
      </div>
    </div>

  )
}

export default App