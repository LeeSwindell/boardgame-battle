import CardContainer from './CardContainer';
import DarkArtCard from './Darkart';
import Location from './Location';
import MarketCard from './Marketcard';
import Monster from './Monster';
import Villain from './Villain';

function Gameboard(props) {
  return (
    <div className="flex flex-col m-1 border">
      {/* Top Row */}
      <div className="flex flex-none p-1 space-x-8 justify-between border">
        <CardContainer card={<Location />} size="wide" />
        <div className="flex space-x-2 justify-center">
          <CardContainer card={<DarkArtCard />} size="reg" />
          <CardContainer card={<Monster />} size="reg" />
        </div>
        <div className="flex flex-none space-x-2 justify-end">
          <CardContainer card={<MarketCard />} size="reg" />
          <CardContainer card={<MarketCard />} size="reg" />
          <CardContainer card={<MarketCard />} size="reg" />
        </div>
      </div>

      {/* Middle Row */}
      <div className="flex flex-none p-1 space-x-8 justify-between border">
        <CardContainer card={<Villain />} size="wide" />
        <CardContainer card={<Villain />} size="wide" />
        <CardContainer card={<Villain />} size="wide" />
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
