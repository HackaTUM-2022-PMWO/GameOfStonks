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
        <form className="flex flex-col justify-center items-center h-screen space-y-5">
            <h1>Welcome to the GAME OF STONKS!</h1>
            <h2>Tell us your name...</h2>
            <Input type={"text"} value={userName}
                   onChange={e => setUserName((e.target as unknown as HTMLTextAreaElement).value)}/>
            <RouterButton onClick={handleSubmit} route={Routes.Lobby} text={"start trading stonks"}/>
        </form>
    );
}

export default Onboard;
