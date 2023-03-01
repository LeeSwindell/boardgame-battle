import CardContainer from "./cardcontainer";
import MarketCard from "./marketcard";

function InspectCard() {
    return (
        <div className="fixed w-full h-full backdrop-contrast-50 invisible">
        <div className="flex w-full h-full justify-center items-center">
            <div className="border border-pink-500 bg-white w-64 h-64 z-50 shadow-2xl">
                <CardContainer card={<MarketCard />} size="regzoom" />
            </div>
        </div>
      </div>
    )
}

export default InspectCard;