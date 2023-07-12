import Card from './Card';

function Location({ curControl, maxControl, imgPath }) {
  return (
    <div className="relative">
      <Card src={imgPath} alt="Location" />
      <div className="absolute text-white text-lg font-bold bottom-2 right-2 w-6 h-6 bg-teal-800">{maxControl - curControl}</div>
    </div>
  );
}

export default Location;
