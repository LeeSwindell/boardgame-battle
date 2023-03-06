import { useInspect } from "../App";

export const Sizes = {
    "reg": "w-32 h-32",
    "wide": "w-40 h-32",
    "regzoom": "w-64 h-64",
    "widezoom": "w-80, h-64"
}

function CardContainer({ card, size, extra }) {
    const { setInspectCard } = useInspect();

    const onClick = (e) => {
        e.preventDefault();
        if (e.nativeEvent.button === 2) {
            setInspectCard(card)
        }
    }

    return (
        <button className={`border hover:shadow-lg ${Sizes[size]} ${extra} rounded`} onClick={onClick} onContextMenu={onClick}>
            {card}
        </button>
    )
}

export default CardContainer;