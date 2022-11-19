import {RouterButton} from "../components/buttons/RouterButton";
import {Routes} from "../router/router";
import React, {useState} from "react";
import {Input} from "../components/inputs/Input";
import {useStonkState} from "../model/store";
import {useNavigate} from "react-router-dom";
import {Button} from "../components/buttons/Button";

function Onboard() {
    const [userName, setUserName] = useState("");
    const navigate = useNavigate();
    const register = useStonkState((state) => state.register);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        register(userName).then((resp) => navigate(Routes.Lobby));
        // alert(`The name you entered was: ${userName}`);
        // TODO: handle errors
    };

    return (
        <form
            onSubmit={handleSubmit}
            className="flex flex-col justify-center items-center h-screen space-y-5"
        >
            <h1>Welcome to the GAME OF STONKS!</h1>
            <div>
                <iframe src="https://giphy.com/embed/bMycGOQLESDCEnLNUz" width="100%" height="100%"
                        frameBorder="0" className="giphy-embed" allowFullScreen></iframe>
            </div>
            <h2>Tell us your name...</h2>
            <Input
                name="username"
                type="text"
                value={userName}
                onChange={(e) =>
                    setUserName((e.target as unknown as HTMLTextAreaElement).value)
                }
            />
            <Button type="submit">start trading stonks</Button>
        </form>
    );
}

export default Onboard;
