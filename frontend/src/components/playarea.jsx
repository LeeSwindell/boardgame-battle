import MarketCard from "./marketcard";
import CardContainer from "./cardcontainer";

function PlayArea () {
    return (
        <div className="flex flex-row relative -space-x-16 justify-center">
            <CardContainer card={<MarketCard />} size="reg" />
            <CardContainer card={<MarketCard />} size="reg" />
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

export default PlayArea;