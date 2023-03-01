export const Sizes = {
    "reg": "w-32 h-32",
    "wide": "w-40 h-32",
    "regzoom": "w-64 h-64",
    "widezoom": "w-80, h-64"
}

function CardContainer(props) {
    const { card, size } = props
    return (
        <button className={`border hover:shadow-lg ${Sizes[size]} rounded`}>
            {card}
        </button>
    )
}

export default CardContainer;