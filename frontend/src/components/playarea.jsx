import MarketCard from "./marketcard";
import CardContainer from "./cardcontainer";

function PlayArea (props) {
    const { setInspect, setCardToInspect } = props;
    return (
        <div className="flex flex-row -space-x-16 justify-center">
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
        </div>
    )
}

export default PlayArea;