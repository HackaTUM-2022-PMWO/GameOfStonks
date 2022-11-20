import { Link } from "react-router-dom";
import SvgSearch from "../../icons/Search";
import { RouterButton } from "../buttons/RouterButton";
import { Routes } from "../../router/router";
import SvgArchive from "../../icons/Archive";
import logo from "./../../assets/logo.png";
import { useStonkState } from "../../model/store";

export function Header() {
  const gameStarted = useStonkState((state) => state.gameStarted);
  return (
    <header>
      <div className="grid grid-cols-3 items-center shadow px-6 py-3 max-w-screen">
        <Link
          to={gameStarted ? "/" : Routes.Home}
          className="flex items-center"
        >
          <img src={logo} className="mr-3 h-6 sm:h-9" alt="logo" />
          <h3>GameOfStonks</h3>
        </Link>
        {gameStarted && <h3 className="flex justify-center">10:00</h3>}
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
