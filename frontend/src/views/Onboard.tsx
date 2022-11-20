import { Routes } from "../router/router";
import React, { useState } from "react";
import { Input } from "../components/inputs/Input";
import { useStonkState } from "../model/store";
import { useNavigate } from "react-router-dom";
import { Button } from "../components/buttons/Button";
import { Card } from "../components/card/Card";

function Onboard() {
  const [userName, setUserName] = useState("");
  const navigate = useNavigate();
  const rocket = require("./../assets/rocket.png");
  const register = useStonkState((state) => state.register);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (userName === "") {
      alert("Please specify a valid user name to participate in GameOfStonks!");
    } else {
      register(userName, navigate).then((resp) => {
        if (resp.ret_1 == null) {
          navigate(Routes.Lobby);
        } else {
          alert(
            "It seems like the all sessions are currently full, please wait, clear cookies and retry!"
          );
        }
      });
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex flex-col items-center space-y-5 min-h-screen"
    >
      <div>
        <Card className="flex flex-col space-y-7 lg:w-lg">
          <h1>Get Started</h1>
          <div className="flex flex-col space-y-2">
            <h4>Tell us your name:</h4>
            <Input
              className="text-5xl bg-transparent min-w-0 border-l-0 border-t-0 border-r-0 border-b-4 text-white"
              name="username"
              type="text"
              value={userName}
              onChange={(e) =>
                setUserName((e.target as unknown as HTMLTextAreaElement).value)
              }
            />
          </div>
          <Button type="submit">Start Game</Button>
        </Card>
        <Card className="pointer-events-none w-md p-0 lg:w-lg"></Card>
      </div>
      <Card className="w-md lg:w-lg space-y-5">
        <div className="flex flex-row gap-4">
          <h1>Welcome to the GAME OF STONKS!</h1>
          <img className="w-12 h-12" src={rocket} alt={"!"} />
        </div>
        <div>
          <h3>Project Idea</h3>
          <p>
            Have you heard of the person that traded his way from a single red
            paperclip all the way to a house? Well, it’s a true story and it
            inspired us to build a fun game about trading digital items among
            friends (don’t worry, we did not work on NFTs 😀). Besides building
            something fun, we believe that financial education is incredibly
            important and set out to code a digital exchange that entertains
            people from all ages while they learn about market mechanisms
            through playful interaction. We are calling it
            <b> GameOfStonks</b>.
          </p>
          <br />
          <h3>How it works</h3>
          <p>
            GameOfStonks is a multiplayer game that runs in your browser, where
            you and your friends trade the so-called “stonks” with each other
            and with the machine. You enter the game by logging in on our
            website and wait in the lobby until enough other players have joined
            so that the game can begin. Now you have 10min to place “bid” and
            “ask” orders on the exchange and trade to increase your net worth.
            In the humble beginnings of your career you have no more than a few
            paper clips and some coins to your name. But if you have the feel
            for the market, if you are a true wolf of wall street, these paper
            clips are not just paper clips to you, they are your chance to climb
            up the ladder! As soon as the timer hits zero, the game ends and the
            rich winners are celebrated. Are you up for the challenge?
          </p>
          <br />
          <iframe
            src="https://giphy.com/embed/bMycGOQLESDCEnLNUz"
            width="100%"
            height="100%"
          ></iframe>
        </div>
        {/* <div>
            <h4>Game Rules</h4>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
              eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut
              enim ad minim veniam, quis nostrud exercitation ullamco laboris
              nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
              reprehenderit in voluptate velit esse cillum dolore eu fugiat
              nulla pariatur. Excepteur sint occaecat cupidatat non proident,
              sunt in culpa qui officia deserunt mollit anim id est laborum.
            </p>
          </div>
          <div>
            <h4>Credits</h4>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
              eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut
              enim ad minim veniam, quis nostrud exercitation ullamco laboris
              nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
              reprehenderit in voluptate velit esse cillum dolore eu fugiat
              nulla pariatur. Excepteur sint occaecat cupidatat non proident,
              sunt in culpa qui officia deserunt mollit anim id est laborum.
            </p>
          </div> */}
      </Card>
    </form>
  );
}

export default Onboard;
