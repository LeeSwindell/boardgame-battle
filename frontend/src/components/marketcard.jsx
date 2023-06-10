import Card from './Card';

function MarketCard({ img }) {
  return (
    <Card src={img} alt="Cornish Pixies" className="object-contain" />
  );
}

export default MarketCard;
