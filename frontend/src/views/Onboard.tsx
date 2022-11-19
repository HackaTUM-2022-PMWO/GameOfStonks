import {RouterButton} from "../components/buttons/RouterButton";
import {Routes} from "../router/router";
import React, {useState} from "react";
import {Input} from "../components/inputs/Input";
import {useStonkState} from "../model/store";
import {useNavigate} from "react-router-dom";
import {Button} from "../components/buttons/Button";
import {Card} from "../components/card/Card";

function Onboard() {
    const [userName, setUserName] = useState("");
    const navigate = useNavigate();
    const register = useStonkState((state) => state.register);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        register(userName).then((resp) => navigate(Routes.Lobby));
        // TODO: handle errors
    };

    return (
        <form
            onSubmit={handleSubmit}
            className="flex flex-col justify-center items-center h-screen space-y-5"
        >
            <Card className="max-w-md">
                <h1>Welcome to the GAME OF STONKS!</h1>
                <p>This game is exists to entertain, to learn and to push GameStop through the roof, STONKS! Invite your fellow mates and enjoy trading.</p>
            </Card>
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
