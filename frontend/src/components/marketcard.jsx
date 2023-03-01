import { useState } from "react";

function MarketCard () {
    const [showModal, setShowModal] = useState(false);
    
    const handleClick = (e) => {
        e.preventDefault();
        if (e.nativeEvent.button === 0) {
            console.log('Left click');
        } else if (e.nativeEvent.button === 2) {
            console.log('Right click');
            setShowModal(true)
        }
      };

    return (
        <button onClick={handleClick} onContextMenu={handleClick}>
            <div className="border w-32 h-32 hover:shadow-lg">
                Market Card
            </div>
            {showModal ? (
                <div className="absolute top-1/2 left-1/2 z-50 shadow-lg w-80 h-80 bg-blue opacity-50">
                    hi
                </div>
            ) : null }
        </button>
    )
}

export default MarketCard;