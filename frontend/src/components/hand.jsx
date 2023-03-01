import MarketCard from "./marketcard";

function Hand () {
    return (
        <div className="flex flex-row space-x-2 justify-center border">
            <MarketCard />
            <MarketCard />
            <MarketCard />
            <MarketCard />
            <MarketCard />
            <MarketCard />
            <MarketCard />
        </div>
    )
}

export default Hand;