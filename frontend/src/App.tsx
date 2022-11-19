import React from 'react';
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import Onboard from "./views/Onboard";
import Detail from "./views/Detail";
import Home from "./views/Home";
import Lobby from "./views/Lobby";
import Result from "./views/Result";
import Search from "./views/Search";
import StartStocks from "./views/StartStocks";
import Trade from "./views/Trade";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Onboard/>,
    },
    {
        path: "/detail",
        element: <Detail/>
    },
    {
        path: "/home",
        element: <Home/>
    },
    {
        path: "/lobby",
        element: <Lobby/>
    },
    {
        path: "/result",
        element: <Result/>
    },
    {
        path: "/search",
        element: <Search/>
    },
    {
        path: "/start",
        element: <StartStocks/>
    },
    {
        path: "/trade",
        element: <Trade/>
    },
]);

function App() {
  return (
    <div>
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
