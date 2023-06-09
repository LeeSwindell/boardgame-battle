import MarketCard from './Marketcard';
import CardContainer from './CardContainer';
import { useGamestate } from '../routes/Game';
import { ConsoleLogger, logger } from '../logger/logger';

function PlayArea() {
  const { gamestate } = useGamestate();
  const username = localStorage.getItem('sessionid');

  if (gamestate) {
    logger.log(gamestate);
    return (
      <div className="flex flex-row space-x-2 justify-center border h-40">
        {gamestate.players[username].PlayArea.map((card) => (
          <CardContainer key={card.Id} cardId={card.Id} cardType="hand" card={<MarketCard img={card.ImgPath} />} size="reg" />
        ))}
      </div>
    );
  }
}

export default PlayArea;
