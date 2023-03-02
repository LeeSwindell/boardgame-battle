import { useEffect, useRef } from "react";
import CardContainer from "./cardcontainer";

function InspectCard(props) {
    const { card, setInspect, setCardToInspect } = props
    const ref = useRef(null);

    useEffect(() => {
        const handleClick = (event) => {
            event.preventDefault()
            if (ref.current && ref.current.contains(event.target)) {
                console.log("should close inspect window")
                setInspect(false)
                setCardToInspect()
            }
        };

        document.addEventListener('click', handleClick);

        return () => {
            document.removeEventListener('click', handleClick);
        }
    }, []);
    
    return (
        <div className="fixed w-full h-full backdrop-contrast-50" ref={ref}>
        <div className="flex w-full h-full justify-center items-center">
            <div className="border bg-white w-64 h-64 z-50 shadow-2xl">
                <CardContainer card={card} size="regzoom" />
            </div>
        </div>
      </div>
    )
}

export default InspectCard;