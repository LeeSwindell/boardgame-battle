import { useEffect } from "react";
import CardContainer from "./cardcontainer";
import MarketCard from "./marketcard";

function InspectCard({ card }) {
    
    useEffect(() => {
        const handleClick = (event) => {
            event.preventDefault()
            console.log("should close inspect window")
        };

        document.addEventListener('click', handleClick);
        return () => {
            document.removeEventListener('click', handleClick);
        }
    }, []);
    
    return (
        <div className="fixed w-full h-full backdrop-contrast-50">
        <div className="flex w-full h-full justify-center items-center">
            <div className="border bg-white w-64 h-64 z-50 shadow-2xl">
                <CardContainer card={card} size="regzoom" />
            </div>
        </div>
      </div>
    )
}

export default InspectCard;