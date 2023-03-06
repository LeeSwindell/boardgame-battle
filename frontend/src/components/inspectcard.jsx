import { useEffect, useRef } from "react";
import { useInspect } from "../App";
import CardContainer from "./cardcontainer";

function InspectCard() {
    const {inspectCard, setInspectCard} = useInspect();
    const ref = useRef(null);

    useEffect(() => {
        const handleClick = (event) => {
            event.preventDefault()
            if (ref.current && ref.current.contains(event.target)) {
                setInspectCard(undefined)
            }
        };

        document.addEventListener('click', handleClick);

        return () => {
            document.removeEventListener('click', handleClick);
        }
    }, []);
    
    if (inspectCard === undefined) {
        return null;
    }

    return (
        <div className="fixed w-full h-full backdrop-contrast-50" ref={ref}>
        <div className="flex w-full h-full justify-center items-center">
            <div className="border bg-white z-50 shadow-2xl">
                <CardContainer extra="w-80 h-80" card={inspectCard}/>
            </div>
        </div>
      </div>
    )
}

export default InspectCard;