import { PlayerListItem } from "../components/listItems/PlayerListItem";
import { Spinner } from "../components/spinner/Spinner";
import { useStonkState } from "../model/store";

function Lobby() {
  const players = useStonkState((state) => state.sessionUsers); // TOOD: replace with SSE or getPlayers method
  return (
    <div className="flex flex-col justify-center items-center h-screen space-y-5">
      <h1 className="text-5xl">Waiting for more traders</h1>
      <p>Trading will begin soon</p>
      <h2>{players == null ? "0" : "" + players.length} of 5</h2>

      {players != null && (
        <ul className="space-y-3 text-3xl">
          {players.map((player, index) => (
            <PlayerListItem key={player.Name} idx={index} value={player.Name} />
          ))}
        </ul>
      )}
      <Spinner />
    </div>
  );
}

export default Lobby;
