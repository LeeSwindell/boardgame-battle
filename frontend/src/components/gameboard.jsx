import DarkArtCard from "./darkart";
import Location from "./location";
import MarketCard from "./marketcard";
import Monster from "./monster";
import Villain from "./villain";

function Gameboard() {
    return (
        <div className="flex flex-col m-1 border">

             {/* Top Row */}
            <div className="flex flex-none p-1 space-x-8 justify-between border">
                <Location />
                <div className="flex space-x-2 justify-center">
                    <DarkArtCard />
                    <Monster />
                </div>
                <div className="flex flex-none space-x-2 justify-end">
                    <MarketCard />
                    <MarketCard />
                    <MarketCard />
                </div>
            </div>

            {/* Middle Row */}
            <div className="flex flex-none p-1 space-x-8 justify-between border">
                <Villain />
                <Villain />
                <Villain />
                <div className="flex flex-none space-x-2 justify-end">
                    <MarketCard />
                    <MarketCard />
                    <MarketCard />
                </div>
            </div>
        </div>
    )
}

export default Gameboard;