import MarketCard from './Marketcard';
import CardContainer from './CardContainer';
import { useGamestate } from '../routes/Game';

function PlayArea() {
  const { gamestate } = useGamestate();
  const username = localStorage.getItem('sessionid');

  if (gamestate) {
    return (
      <div className="flex flex-row space-x-2 p-2 justify-center border">
        {gamestate.players[username].PlayArea.Cards.map((card) => (
          <CardContainer key={card.Id} cardId={card.Id} cardType="hand" card={<MarketCard img={card.ImgPath} />} size="reg" />
        ))}
      </div>
    );
  }
}

export default PlayArea;
