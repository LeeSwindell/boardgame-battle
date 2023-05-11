function Location({ curControl, maxControl, imgPath }) {
  return (
    <div>
      <p>
        Control:
        {curControl}
        /
        {maxControl}
      </p>
      <img src={imgPath} alt="great hall" />
    </div>
  );
}

export default Location;
