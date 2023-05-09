function Villain({ maxHp, curDamage, img }) {
  return (
    <div className="">
      <p>
        Health:
        {maxHp - curDamage}
        /
        {maxHp}
      </p>
      <img src={img} alt="cornish pixies" />
    </div>
  );
}

export default Villain;
