import MarketCard from './Marketcard';
import CardContainer from './CardContainer';

function PlayArea(props) {
  return (
    <div className="flex flex-row -space-x-16 justify-center">
      <CardContainer card={<MarketCard />} size="reg" />
      <CardContainer card={<MarketCard />} size="reg" />
      <CardContainer card={<MarketCard />} size="reg" />
      <CardContainer card={<MarketCard />} size="reg" />
      <CardContainer card={<MarketCard />} size="reg" />
      <CardContainer card={<MarketCard />} size="reg" />
      <CardContainer card={<MarketCard />} size="reg" />
      <CardContainer card={<MarketCard />} size="reg" />
      <CardContainer card={<MarketCard />} size="reg" />
    </div>
  );
}

export default PlayArea;
