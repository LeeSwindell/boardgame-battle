import CardContainer from "./cardcontainer";
import MarketCard from "./marketcard";

function Hand (props) {
    const { setInspect, setCardToInspect } = props;

    return (
        <div className="flex flex-row space-x-2 p-2 justify-center border">
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect} />
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
            <CardContainer card={<MarketCard />} size="reg" setInspect={setInspect} setCardToInspect={setCardToInspect}/>
        </div>
    )
}

export default Hand;