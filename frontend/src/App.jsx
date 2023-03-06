import CardContainer from "./components/cardcontainer"
import DiscardPile from "./components/discard"
import EndTurn from "./components/endturn"
import Gameboard from "./components/gameboard"
import Hand from "./components/hand"
import InspectCard from "./components/inspectcard"
import PlayArea from "./components/playarea"
import PlayerInfo from "./components/playerinfo"
import { useState, createContext, useContext } from "react"

// use react context to pass inspect around
const InspectContext = createContext()

function InspectProvider({children}) {
  const [inspectCard, setInspectCard] = useState();
  const value = {inspectCard, setInspectCard}
  return <InspectContext.Provider value={value}>{children}</InspectContext.Provider>
}

function useInspect() {
  const context = useContext(InspectContext)
  if (context === undefined) {
    throw new Error('useInspect must be used within an InspectProvider')
  }
  return context
}

export { useInspect }

function App() {
  return (
    <InspectProvider>
      <InspectCard />
      <div className="flex flex-row justify-between">
        <div className="flex flex-col space-y-4 w-auto h-auto">
          <Gameboard/>
          <PlayArea/>
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
  )
}

export default App