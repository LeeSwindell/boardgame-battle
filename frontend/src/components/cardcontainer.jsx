import { useInspect } from '../routes/Game';
import { gameapi } from '../api';

export const Sizes = {
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
    console.log('playing card:', id);
    gameapi
      .post('/0/playcard', { id })
      .then(() => {
        // console.log('playcard TestCard');
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
  };

  return (
    <button className={`border hover:shadow-lg ${Sizes[size]} ${extra} rounded`} onClick={onClick} onContextMenu={onClick} type="button">
      {card}
    </button>
  );
}

export default CardContainer;
