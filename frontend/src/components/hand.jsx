import CardContainer from "./CardContainer";
import MarketCard from "./Marketcard";

function Hand () {
    return (
        <div className="flex flex-row space-x-2 p-2 justify-center border">
            <CardContainer card={<MarketCard />} size="reg" />
            <CardContainer card={<MarketCard />} size="reg" />
            <CardContainer card={<MarketCard />} size="reg" />
            <CardContainer card={<MarketCard />} size="reg" />
            <CardContainer card={<MarketCard />} size="reg" />
            <CardContainer card={<MarketCard />} size="reg" />
            <CardContainer card={<MarketCard />} size="reg" />
        </div>
    )
}

export default Hand;