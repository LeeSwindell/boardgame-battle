import CardContainer from './CardContainer';
import MarketCard from './Marketcard';
import { useGamestate } from '../routes/Game';

function Hand() {
  const { gamestate } = useGamestate();
  const username = localStorage.getItem('sessionid');

  if (gamestate) {
    return (
      <div className="flex flex-row space-x-2 justify-center border">
        {gamestate.players[username].Hand.map((card) => (
          <CardContainer key={card.Id} cardId={card.Id} cardType="hand" card={<MarketCard img={card.ImgPath} />} size="reg" />
        ))}
      </div>
    );
  }
}

export default Hand;
