import { logger } from '../logger/logger';
import Card from './Card';

const profLookup = {
  Arithmancy: 'arithmancy.jpg',
  'Care of Magical Creatures': 'careofmagicalcreatures.jpg',
  Charms: 'charms.jpg',
  'Defense Against the Dark Arts': 'defenseagainstthedarkarts.jpg',
  Divination: 'divination.jpg',
  'Flying Lessons': 'flyinglessons.jpg',
  Herbology: 'herbology.jpg',
  'History of Magic': 'historyofmagic.jpg',
  Potions: 'potions.jpg',
  Transfiguration: 'transfiguration.jpg',
};

function Proficiency({ name }) {
  logger.log(name);
  return (
    <div className="w-40 h-32 relative">
      <Card src={`/images/charcards/${profLookup[name]}`} alt="Villain" />
    </div>
  );
}

export default Proficiency;
