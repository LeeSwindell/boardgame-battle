import { useEffect, useRef } from 'react';
import { useInspect } from '../routes/Game';
import CardContainer from './CardContainer';
import MarketCard from './Marketcard';

function InspectCard() {
  const { inspectCard, setInspectCard } = useInspect();
  const ref = useRef(null);

  useEffect(() => {
    const handleClick = (event) => {
      event.preventDefault();
      if (ref.current && ref.current.contains(event.target)) {
        setInspectCard(undefined);
      }
    };

    document.addEventListener('click', handleClick);

    return () => {
      document.removeEventListener('click', handleClick);
    };
  }, []);

  if (inspectCard === undefined) {
    return null;
  }

  if (Array.isArray(inspectCard)) {
    console.log('array of inspects: ', inspectCard);
    return (
      <div className="fixed w-full h-full backdrop-contrast-50" ref={ref}>
        <div className="flex w-full h-full justify-center items-center">
          <div className="border bg-white z-50 shadow-2xl">
            <div className="grid grid-cols-8">
              {inspectCard.map((card) => (
                <CardContainer key={card.Id} cardId={card.Id} cardType="hand" card={<MarketCard img={card.ImgPath} />} size="reg" />
              ))}
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="fixed w-full h-full backdrop-contrast-50" ref={ref}>
      <div className="flex w-full h-full justify-center items-center">
        <div className="border bg-white z-50 shadow-2xl">
          <CardContainer extra="w-80 h-80" card={inspectCard} />
        </div>
      </div>
    </div>
  );
}

export default InspectCard;
