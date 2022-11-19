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
  const register = useStonkState((state) => state.register);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (userName === "") {
      alert("Please specify a valid user name to participate in GameOfStonks!");
    } else {
      register(userName).then((resp) => {
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
      <div className="grid grid-cols-2">
        <Card className="w-md lg:w-lg space-y-5">
          <h1>Welcome to the GAME OF STONKS!</h1>
          <div>
            <h4>Project Idea</h4>
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
          </div>
        </Card>
        <div>
          <Card className="flex flex-col space-y-5 w-md lg:w-lg">
            <h1>Get Started</h1>
            <div className="flex flex-col space-y-2">
              <h4>Tell us your name:</h4>
              <Input
                className=""
                name="username"
                type="text"
                value={userName}
                onChange={(e) =>
                  setUserName(
                    (e.target as unknown as HTMLTextAreaElement).value
                  )
                }
              />
            </div>
            <Button type="submit">Start Game</Button>
          </Card>
          <Card className="pointer-events-none w-md lg:w-lg">
            <iframe
              className="h-[40vh]"
              src="https://giphy.com/embed/bMycGOQLESDCEnLNUz"
              width="100%"
              height="100%"
            ></iframe>
          </Card>
        </div>
      </div>
    </form>
  );
}

export default Onboard;
