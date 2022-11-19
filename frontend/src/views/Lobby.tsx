import {PlayerListItem} from "../components/listItems/PlayerListItem";
import {Spinner} from "../components/spinner/Spinner";
import {useStonkState} from "../model/store";

function Lobby() {
    const players = useStonkState((state) => state.sessionUsers);
    return (
        <div className="flex flex-col justify-center items-center h-screen space-y-5">
            <h1>Waiting for more STONK traders to join...</h1>
            {players != null && <ul>
                {players.map((player, index) => (
                    <PlayerListItem key={player.Name} idx={index} value={player.Name}/>
                ))}
            </ul>}
            <Spinner/>
            <h2>{players == null ? "0" : "" + players.length} of 5</h2>
        </div>
    );
}

export default Lobby;
