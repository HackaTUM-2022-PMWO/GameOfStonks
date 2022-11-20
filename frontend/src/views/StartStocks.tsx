import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Routes } from "../router/router";
import { StonksAssetsMatch } from "../assets/StonksAssetsMatch";
import coin from "./../assets/coin.png";

function StartStocks() {
  const stonk = StonksAssetsMatch[1].img; // TODO: replace with dynamic icon image
  // const delay = (ms: number | undefined) => new Promise(
  //     resolve => setTimeout(resolve, ms)
  // );
  const navigate = useNavigate();
  const [counter, setCounter] = useState(5);

  useEffect(() => {
    async function redirectHome() {
      for (let i = 1; i <= 5; i++) {
        // await delay(1000);
        setCounter(counter - i);
      }
    }
    redirectHome().then((r) => navigate(Routes.Home));
  });

  return (
    <div className="flex flex-col justify-center items-center h-screen space-y-16">
      <h1>Your first Stonks & StonkCoins:</h1>
      <div className="grid grid-cols-2 space-x-10 justify-items-center space-y-5 items-center">
        <img className="w-32 h-32" src={stonk} alt={"stonk"} />
        <img className="w-32 h-32" src={coin} alt={"coin"} />
        <h3>5x Paperclip (PACP)</h3>
        <h3>1000x StonkCoins</h3>
      </div>
      <h1>{counter}</h1>
    </div>
  );
}

export default StartStocks;
