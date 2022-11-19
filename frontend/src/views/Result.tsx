import useWindowSize from 'react-use/lib/useWindowSize'
import Confetti from 'react-confetti'
import {Currency} from "../components/Currency";
import {useEffect, useState} from "react";
import {Routes} from "../router/router";
import {RouterButton} from "../components/buttons/RouterButton";

function Result() {
    const {width, height} = useWindowSize()
    const first = require(`./../assets/awards/1.png`);
    const second = require(`./../assets/awards/2.png`);
    const third = require(`./../assets/awards/3.png`);
    const [nrConfetti, setNrConfetti] = useState(200);
    const delay = (ms: number | undefined) => new Promise(
        resolve => setTimeout(resolve, ms)
    );

    useEffect(() => {
        async function redirectHome() {
            await delay(10000)
        }

        redirectHome().then(r => setNrConfetti(0));
    })


    let results = [{name: "Greta", coins: 2239}, {name: "Elon", coins: 1899}, {
        name: "FreshHermann",
        coins: 1820
    }, {name: "Spong", coins: 289}, {name: "Trump", coins: 50}]

    return (
        <div className="flex flex-col gap-8 items-center">
            <div className="flex items-end justify-center min-h-[90vh]">
                <div className="grid items-end mx-5 grid-cols-3">
                    <div className="flex flex-col content-center text-center gap-4">
                        <h1>{results[1].name}</h1>
                        <div
                            className="flex flex-col items-center gap-8 h-[70vh] bg-accent2 shadow-md drop-shadow-xl rounded">
                            <img className=" w-[50%]" src={second} alt={"second place"}/>
                            <div className="flex">
                                <h1>{results[1].coins}</h1>
                                <Currency/>
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-col content-center text-center gap-4">
                        <h1>{results[0].name}</h1>
                        <div
                            className="flex flex-col items-center gap-8 h-[80vh] bg-accent2 shadow-md drop-shadow-xl rounded">
                            <img className=" w-[50%]" src={first} alt={"first place"}/>
                            <div className="flex">
                                <h1>{results[0].coins}</h1>
                                <Currency/>
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-col content-center text-center gap-4">
                        <h1>{results[2].name}</h1>
                        <div
                            className="flex flex-col items-center gap-8 h-[60vh] bg-accent2 shadow-md drop-shadow-xl rounded">
                            <img className=" w-[50%]" src={third} alt={"third place"}/>
                            <div className="flex">
                                <h1>{results[2].coins}</h1>
                                <Currency/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <RouterButton route={Routes.Onboard}>End Game</RouterButton>
            <Confetti
                width={width * 0.98}
                height={height}
                numberOfPieces={nrConfetti}
            />
        </div>
    );
}

export default Result;
