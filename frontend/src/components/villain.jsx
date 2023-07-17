import Card from './Card';

function Villain({ maxHp, curDamage, img }) {
  return (
    <div className="relative">
      <Card src={img} alt="Villain" />
      <div className="absolute text-white text-lg font-bold bottom-2 right-2 w-6 h-6 bg-red-700">{maxHp - curDamage}</div>
    </div>
  );
}

export default Villain;
