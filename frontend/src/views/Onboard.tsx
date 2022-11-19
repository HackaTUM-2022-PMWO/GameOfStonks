import {RouterButton} from "../components/buttons/RouterButton";
import {Routes} from "../router/router";
import {useState} from "react";
import {Input} from "../components/inputs/Input";

function Onboard() {
    const [userName, setUserName] = useState('');

    const handleSubmit = () => {
        // alert(`The name you entered was: ${userName}`);
        // TODO: handle errors
    }

    return (
            <form className="grid grid-cols-1 justify-center justify-items-center space-y-5">
                <label>Tell us your name...</label>
                <Input type={"text"} value={userName} onChange={e => setUserName((e.target as unknown as HTMLTextAreaElement).value)}/>
                <RouterButton onClick={handleSubmit} route={Routes.Lobby} text={"start trading stonks"}/>
            </form>
    );
}

export default Onboard;
