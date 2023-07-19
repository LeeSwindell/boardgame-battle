import { useGamestate, useInspect } from '../routes/Game';
import MarketCard from './Marketcard';

function DiscardPile() {
  const { gamestate } = useGamestate();
  const { setInspectCard } = useInspect();
  const user = localStorage.getItem('sessionid');

  const onClick = (e) => {
    e.preventDefault();
    if (e.nativeEvent.button === 2) {
      setInspectCard(gamestate.players[user].Discard);
    }
  };

  return (
    <button className="w-32 h-40 rounded m-2 " type="button" onClick={onClick} onContextMenu={onClick}>
      {gamestate.players[user].Discard[gamestate.players[user].Discard.length - 1]
       && (
       <MarketCard
         img={gamestate.players[user].Discard[gamestate.players[user].Discard.length - 1].ImgPath}
       />
       )}
    </button>
  );
}

export default DiscardPile;
