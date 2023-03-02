import CardContainer from "./components/cardcontainer"
import DiscardPile from "./components/discard"
import EndTurn from "./components/endturn"
import Gameboard from "./components/gameboard"
import Hand from "./components/hand"
import InspectCard from "./components/inspectcard"
import PlayArea from "./components/playarea"
import PlayerInfo from "./components/playerinfo"
import { useState } from "react"

function App() {
  const [inspect, setInspect] = useState(false);
  const [cardToInspect, setCardToInspect] = useState(<div></div>)

  const showInspectCard = (card) => {
    if (inspect) {
      return <InspectCard card={card} setInspect={setInspect} setCardToInspect={setCardToInspect}/>
    }
  }

  return (
    <>
      {showInspectCard(cardToInspect)}
      <div className="flex flex-row justify-between">
        <div className="flex flex-col space-y-4 w-auto h-auto">
          <Gameboard setInspect={setInspect} setCardToInspect={setCardToInspect}/>
          <PlayArea setInspect={setInspect} setCardToInspect={setCardToInspect}/>
          <Hand setInspect={setInspect} setCardToInspect={setCardToInspect}/>
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
            <CardContainer card={<DiscardPile />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect} extra="p-2 m-2" />
          </div>
        </div>
      </div>
    </>
  )
}

export default App