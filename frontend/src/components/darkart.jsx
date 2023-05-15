import { useGamestate, useInspect } from '../routes/Game';

function DarkArtCard() {
  const { gamestate } = useGamestate();
  const { setInspectCard } = useInspect();

  const onClick = (e) => {
    e.preventDefault();
    if (e.nativeEvent.button === 2 && gamestate.darkartsplayed) {
      setInspectCard(gamestate.darkartsplayed);
    }
  };

  return (
    <button className="flex w-32 h-32 rounded m-4 border items-center justify-center hover:shadow-lg" type="button" onClick={onClick} onContextMenu={onClick}>
      {gamestate.darkartsplayed.length > 0
       && <img src={gamestate.darkartsplayed[gamestate.darkartsplayed.length - 1].ImgPath} alt="dark art" />}
      {gamestate.darkartsplayed.length <= 0
       && <div>Dark Art</div>}
    </button>
  );
}

export default DarkArtCard;
