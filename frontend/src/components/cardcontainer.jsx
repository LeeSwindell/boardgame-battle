import { useGamestate, useInspect } from '../routes/Game';
import { gameapi } from '../api';
import { logger } from '../logger/logger';

export const Sizes = {
  small: 'w-32 h-32',
  reg: 'w-32 h-40',
  wide: 'w-40 h-32',
  regzoom: 'w-64 h-64',
  widezoom: 'w-80, h-64',
};

function CardContainer({
  card, size, extra, cardId, cardType,
}) {
  const { setInspectCard } = useInspect();
  const { gameid } = useGamestate();

  function playCardHandler(id) {
    // logger.log('playing card:', id);
    gameapi
      .post(`/${gameid}/playcard`, { id })
      .then(() => {
        // logger.log('playcard TestCard');
      });
  }

  function damageVillainHandler(id) {
    // logger.log('damaging villain: ', id);
    gameapi
      .get(`/${gameid}/damagevillain/${id}`)
      .then(() => {
        // logger.log(`damaged villain ${id}`, )
      });
  }

  function buyCardHandler(id) {
    gameapi
      .get(`/${gameid}/buycard/${id}`)
      .then(() => {
        // logger.log(`damaged villain ${id}`, )
      });
  }

  function UseProficiencyHandler() {
    // logger.log('damaging villain: ', id);
    gameapi
      .get(`/${gameid}/useproficiency`)
      .then(() => {
        // logger.log(`damaged villain ${id}`, )
      });
  }

  const onClick = (e) => {
    e.preventDefault();
    if (e.nativeEvent.button === 2) {
      setInspectCard(card);
    }
    if (cardType === 'hand' && e.nativeEvent.button === 0) {
      playCardHandler(cardId);
    }
    if (cardType === 'villain' && e.nativeEvent.button === 0) {
      damageVillainHandler(cardId);
    }
    if (cardType === 'market' && e.nativeEvent.button === 0) {
      buyCardHandler(cardId);
    }
    if (cardType === 'proficiency' && e.nativeEvent.button === 0) {
      logger.log('use prof handler');
      UseProficiencyHandler();
    }
  };

  return (
    <button className={`flex items-center justify-center hover:shadow-lg hover:z-20 ${cardType === 'hand' && 'hover:-translate-y-4'} ${Sizes[size]} ${extra}`} onClick={onClick} onContextMenu={onClick} type="button">
      {card}
    </button>
  );
}

export default CardContainer;
