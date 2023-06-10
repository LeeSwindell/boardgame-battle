function Card({ alt, src }) {
  return (
    <img src={src} alt={alt} className="overflow-hidden rounded-lg" />
  );
}

export default Card;
