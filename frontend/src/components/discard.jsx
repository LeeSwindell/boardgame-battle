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
    <button className="w-32 h-40 rounded p-2 m-2 border" type="button" onClick={onClick} onContextMenu={onClick}>
      {gamestate.players[user].Discard[0]
       && <MarketCard img={gamestate.players[user].Discard[0].ImgPath} />}
    </button>
  );
}

export default DiscardPile;
