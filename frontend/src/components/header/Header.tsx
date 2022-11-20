import { Link } from "react-router-dom";
import SvgSearch from "../../icons/Search";
import { RouterButton } from "../buttons/RouterButton";
import { Routes } from "../../router/router";
import SvgArchive from "../../icons/Archive";
import logo from "./../../assets/logo.png";
import { useStonkState } from "../../model/store";

function padTo2Digits(num: number) {
  return num.toString().padStart(2, "0");
}

export function Header() {
  const gameStarted = useStonkState((state) => state.gameStarted);
  const roundDuration = useStonkState((state) => state.roundDuration);

  // 👇️ get number of full minutes
  const minutes = Math.floor(roundDuration / 60);

  // 👇️ get remainder of seconds
  const seconds = roundDuration % 60;

  return (
    <header>
      <div className="grid grid-cols-3 items-center shadow px-6 py-3 max-w-screen">
        <Link
          to={gameStarted ? "/" : Routes.Home}
          className="flex items-center"
        >
          <img src={logo} className="mr-3 h-6 sm:h-9" alt="logo" />
          <h3>
            <b>GameOfStonks</b>
          </h3>
        </Link>
        {gameStarted && (
          <h3 className="flex justify-center">{`${padTo2Digits(
            minutes
          )}:${padTo2Digits(seconds)}`}</h3>
        )}
        {gameStarted && (
          <div className="flex justify-end gap-4">
            <RouterButton route={Routes.History}>
              <SvgArchive />
            </RouterButton>
            <RouterButton route={Routes.Search}>
              <SvgSearch />
            </RouterButton>
          </div>
        )}
      </div>
    </header>
  );
}
