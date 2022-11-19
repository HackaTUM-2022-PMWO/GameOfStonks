import { PlayerListItem } from "../components/listItems/PlayerListItem";
import { Spinner } from "../components/spinner/Spinner";

function Lobby() {
  let nrPlayers = 3; // TODO: von zustand bekommen
  const players = ["Philipp", "Michi", "Wlad"]; // TODO: von zustand bekommen

  return (
    <div className="flex flex-col justify-center items-center h-screen space-y-5">
      <h1>Waiting for more STONK traders to join...</h1>
      <ul>
        {players.map((player) => (
          <PlayerListItem key={player} value={player} />
        ))}
      </ul>
      <Spinner />
      <h2>{nrPlayers} of 5</h2>
    </div>
  );
}

export default Lobby;
