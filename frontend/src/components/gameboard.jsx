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
    <>
      {/* Sidebar with Villains/Location */}
      <div className="inset-y-0 left-0 top-4 w-16 fixed -z-10">
        <div className="flex flex-col space-y-2">
          {gamestate
        && <CardContainer card={<Location imgPath={gamestate.locations[gamestate.currentlocation].ImgPath} curControl={gamestate.locations[gamestate.currentlocation].CurControl} maxControl={gamestate.locations[gamestate.currentlocation].MaxControl} />} size="wide" />}
          {gamestate && gamestate.villains.map((v) => (
            <CardContainer key={v.Id} cardId={v.Id} cardType="villain" card={<Villain img={v.ImgPath} maxHp={v.MaxHp} curDamage={v.CurDamage} />} size="wide" />
          ))}
        </div>
      </div>

      <div className="inset-y-0 left-40 top-1 w-16 fixed -z-10">
        <div className="flex flex-col -space-y-1">
          <DarkArtCard />
          <CardContainer card={<Monster />} size="small" extra="m-2" />
        </div>
      </div>

      {/* Actual board */}
      <div className="fixed mx-1 top-1 left-72 -z-10">
        {/* Top Row */}
        {/* <div className="flex flex-none space-x-8 justify-between border items-center">

          <div className="flex space-x-2 justify-center">
            <DarkArtCard />
            <CardContainer card={<Monster />} size="small" extra="m-4" />
          </div>
          <div className="flex flex-none space-x-2 justify-end">
            Top Row Market
            {gamestate
          && gamestate.market.slice(0, 3).map((c) => <CardContainer key={c.Id} cardId={c.Id} cardType="market" card={<MarketCard img={c.ImgPath} />} size="reg" />)}
          </div>
        </div> */}

        {/* Middle Row */}
        <div className="flex flex-none ml-2 p-1 space-x-8 justify-between items-center">
          <div className="flex flex-none space-x-1 justify-end">
            {/* Middle Row Market */}
            {gamestate
          && gamestate.market.map((c) => <CardContainer key={c.Id} cardId={c.Id} cardType="market" card={<MarketCard img={c.ImgPath} />} size="reg" />)}
          </div>
        </div>
      </div>
    </>
  );
}

export default Gameboard;
