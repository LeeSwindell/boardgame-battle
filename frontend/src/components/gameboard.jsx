import CardContainer from "./cardcontainer";
import DarkArtCard from "./darkart";
import Location from "./location";
import MarketCard from "./marketcard";
import Monster from "./monster";
import Villain from "./villain";

function Gameboard(props) {
    const { setInspect, setCardToInspect } = props;

    return (
        <div className="flex flex-col m-1 border">
             {/* Top Row */}
            <div className="flex flex-none p-1 space-x-8 justify-between border">
                <CardContainer card={<Location />} size="wide" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                <div className="flex space-x-2 justify-center">
                    <CardContainer card={<DarkArtCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                    <CardContainer card={<Monster />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                </div>
                <div className="flex flex-none space-x-2 justify-end">
                    <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                    <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                    <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                </div>
            </div>

            {/* Middle Row */}
            <div className="flex flex-none p-1 space-x-8 justify-between border">
                <CardContainer card={<Villain />} size="wide" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                <CardContainer card={<Villain />} size="wide" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                <CardContainer card={<Villain />} size="wide" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                <div className="flex flex-none space-x-2 justify-end">
                    <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                    <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                    <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
                </div>
            </div>
        </div>
    )
}

export default Gameboard;