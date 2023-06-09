import { gameapi } from '../api';
import { useGamestate } from '../routes/Game';

function EndTurn() {
  const { gameid } = useGamestate();
  function onClick() {
    gameapi
      .get(`/${gameid}/endturn`)
      .then(() => {
        //
      });
  }
  return (
    <button type="submit" className="flex flex-col bg-blue-500 hover:bg-blue-700 text-white w-40 h-16 justify-center items-center m-2 font-bold rounded" onClick={onClick}>
      End Turn
    </button>
  );
}

export default EndTurn;
