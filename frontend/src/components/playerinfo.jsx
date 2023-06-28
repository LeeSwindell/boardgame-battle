import { useGamestate } from '../routes/Game';

function PlayerInfo({ username }) {
  const { gamestate } = useGamestate();
  if (gamestate && username !== undefined) {
    return (
      <div className={`flex flex-col p-2 m-2 w-40 h-16 border bg-white ${gamestate.currentturn === username ? 'border-blue-500' : ''} rounded`}>
        <div className="flex justify-between">
          <div>{gamestate.players[username].Name}</div>
          <div>
            HP:
            {gamestate.players[username].Health}
          </div>
        </div>
        <div className="flex justify-between">
          <div>
            atk:
            {gamestate.players[username].Damage}
          </div>
          <div>
            $
            {gamestate.players[username].Money}
          </div>
        </div>
      </div>
    );
  }
}

export default PlayerInfo;
