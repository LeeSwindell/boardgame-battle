import { useInspect } from '../routes/Game';
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

  function playCardHandler(id) {
    // logger.log('playing card:', id);
    gameapi
      .post('/0/playcard', { id })
      .then(() => {
        // logger.log('playcard TestCard');
      });
  }

  function damageVillainHandler(id) {
    // logger.log('damaging villain: ', id);
    gameapi
      .get(`/0/damagevillain/${id}`)
      .then(() => {
        // logger.log(`damaged villain ${id}`, )
      });
  }

  function buyCardHandler(id) {
    gameapi
      .get(`/0/buycard/${id}`)
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
  };

  return (
    <button className={`flex items-center justify-center hover:shadow-lg ${Sizes[size]} ${extra} rounded`} onClick={onClick} onContextMenu={onClick} type="button">
      {card}
    </button>
  );
}

export default CardContainer;
