import { useGamestate } from '../routes/Game';

function PlayerInfo({ playerIndex }) {
  const { gamestate } = useGamestate();
  console.log('player info:', playerIndex);
  console.log(gamestate);
  if (gamestate && playerIndex >= 0) {
    console.log('rendering playerinfo');
    return (
      <div className="flex flex-col p-2 m-2 w-40 h-16 border rounded">
        <div className="flex justify-between">
          <div>{gamestate.players[playerIndex].Name}</div>
          <div>10</div>
        </div>
        <div className="flex justify-between">
          <div>money:99</div>
          <div>dmg:55</div>
        </div>
      </div>
    );
  }
}

export default PlayerInfo;
