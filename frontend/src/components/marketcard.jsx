const handleClick = (e) => {
    e.preventDefault();
    if (e.nativeEvent.button === 0) {
      console.log('Left click');
    } else if (e.nativeEvent.button === 2) {
      console.log('Right click');
      
    }
  };

function MarketCard () {
    return (
        <button onClick={handleClick} onContextMenu={handleClick}>
            <div className="border w-32 h-32">
                Market Card
            </div>
        </button>
    )
}

export default MarketCard;