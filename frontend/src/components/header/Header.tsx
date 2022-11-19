import {Link} from "react-router-dom";
import {useStonkState} from "../../model/store";

export function Header() {
    const logo = require('./../../assets/logo.png');
    // const gameStarted = useStonkState((state) => state.gameStarted);
    let gameStarted = true;
    return (
        <header>
            <div className="grid grid-cols-3 items-center shadow px-6 py-3 max-w-screen">
                <Link to={"/"} className="flex items-center">
                    <img src={logo} className="mr-3 h-6 sm:h-9"
                         alt="logo"/>
                    <h3>GameOfStonks</h3>
                </Link>
                {gameStarted && <h3 className="flex justify-center">
                    10:00
                </h3>}
                {gameStarted && <div className="flex justify-end">
                    <Search/>
                </div>}
            </div>
        </header>
    );
}