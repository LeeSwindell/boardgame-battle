import CardContainer from './CardContainer';
import MarketCard from './Marketcard';
import { gameapi } from '../api';
import { useGamestate } from '../routes/Game';

function Hand() {
  const { gamestate } = useGamestate();
  function playCardHandler() {
    console.log('playcard TestCard');
    gameapi
      .post('/0/playcard', { cardname: 'TestCard' })
      .then(() => {
        // console.log('playcard TestCard');
      });
  }

  return (
    <div className="flex flex-row space-x-2 p-2 justify-center border">
      <button type="submit" onClick={playCardHandler}>PlayCard</button>
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

export default Hand;
