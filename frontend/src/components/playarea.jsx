import MarketCard from "./marketcard";

function PlayArea () {
    return (
        <div className="flex flex-row relative -space-x-16 justify-center">
            <MarketCard className=""/>
            <MarketCard className=""/>
            <MarketCard className=""/>
            <MarketCard />
            <MarketCard />
            <MarketCard />
            <MarketCard />
            <MarketCard />
            <MarketCard />
        </div>
    )
}

export default PlayArea;