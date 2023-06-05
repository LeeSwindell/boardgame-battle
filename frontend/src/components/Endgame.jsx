import { gameapi } from '../api';

function EndGame() {
  function onClick() {
    const gameid = localStorage.getItem('currentgameid');
    gameapi
      .get(`/${gameid}/shutdown`)
      .then(() => {
        //
      });
  }
  return (
    <button type="submit" className="flex flex-col bg-red-500 hover:bg-red-700 text-white w-40 h-16 justify-center items-center m-2 font-bold rounded" onClick={onClick}>
      End Game
    </button>
  );
}

export default EndGame;
