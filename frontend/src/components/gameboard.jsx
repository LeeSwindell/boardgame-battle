import { useGamestate } from '../routes/Game';
import CardContainer from './CardContainer';
import DarkArtCard from './Darkart';
import Location from './Location';
import MarketCard from './Marketcard';
import Monster from './Monster';
import Villain from './Villain';

function Gameboard(props) {
  const { gamestate } = useGamestate();

  return (
    <div className="flex flex-col">
      {/* Top Row */}
      <div className="flex flex-none p-1 space-x-8 justify-between border items-center">
        {gamestate
        && <CardContainer card={<Location imgPath={gamestate.locations[gamestate.currentlocation].ImgPath} curControl={gamestate.locations[gamestate.currentlocation].CurControl} maxControl={gamestate.locations[gamestate.currentlocation].MaxControl} />} size="wide" />}
        <div className="flex space-x-2 justify-center">
          <DarkArtCard />
          <CardContainer card={<Monster />} size="small" extra="m-4" />
        </div>
        <div className="flex flex-none space-x-2 justify-end">
          <CardContainer card={<MarketCard />} size="reg" />
          <CardContainer card={<MarketCard />} size="reg" />
          <CardContainer card={<MarketCard />} size="reg" />
        </div>
      </div>

      {/* Middle Row */}
      <div className="flex flex-none p-1 space-x-8 justify-between border items-center">
        {gamestate && gamestate.villains.map((v) => (
          <CardContainer key={v.Id} cardId={v.Id} cardType="villain" card={<Villain img={v.ImgPath} maxHp={v.MaxHp} curDamage={v.CurDamage} />} size="wide" />
        ))}
        <div className="flex flex-none space-x-2 justify-end">
          <CardContainer card={<MarketCard />} size="reg" />
          <CardContainer card={<MarketCard />} size="reg" />
          <CardContainer card={<MarketCard />} size="reg" />
        </div>
      </div>
    </div>
  );
}

export default Gameboard;
