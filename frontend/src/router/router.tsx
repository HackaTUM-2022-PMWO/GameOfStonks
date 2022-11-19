import { createBrowserRouter } from "react-router-dom";
import Onboard from "../views/Onboard";
import Detail from "../views/Detail";
import Home from "../views/Home";
import Lobby from "../views/Lobby";
import Result from "../views/Result";
import Search from "../views/Search";
import StartStocks from "../views/StartStocks";
import Trade from "../views/Trade";
import React from "react";
import App from "../App";

export enum Routes {
  Onboard = "/",
  Detail = "/details/:stonkName",
  Home = "/home",
  Lobby = "/lobby",
  Result = "/result",
  Search = "/search",
  StartStocks = "/start",
  Trade = "/trade/:stonkName/:mode",
}

export const getStonkUrl = (stonkName: string) =>
  Routes.Detail.replace(":stonkName", stonkName);

export const getTradeUrl = (stonkName: string, mode: "buy" | "sell") =>
  Routes.Trade.replace(":stonkName", stonkName).replace(":mode", mode);

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: Routes.Onboard,
        element: <Onboard />,
      },
      {
        path: Routes.Detail,
        element: <Detail />,
      },
      {
        path: Routes.Home,
        element: <Home />,
      },
      {
        path: Routes.Lobby,
        element: <Lobby />,
      },
      {
        path: Routes.Result,
        element: <Result />,
      },
      {
        path: Routes.Search,
        element: <Search />,
      },
      {
        path: Routes.StartStocks,
        element: <StartStocks />,
      },
      {
        path: Routes.Trade,
        element: <Trade />,
      },
    ],
  },
]);
