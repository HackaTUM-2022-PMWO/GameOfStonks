import { RouterButton } from "../components/buttons/RouterButton";
import { Routes } from "../router/router";
import React, { useState } from "react";
import { Input } from "../components/inputs/Input";
import { useStonkState } from "../model/store";
import { useNavigate } from "react-router-dom";
import { Button } from "../components/buttons/Button";

function Onboard() {
  const [userName, setUserName] = useState("");
  const navigate = useNavigate();
  const register = useStonkState((state) => state.register);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    register(userName).then((resp) => navigate(Routes.Lobby));
    // alert(`The name you entered was: ${userName}`);
    // TODO: handle errors
  };

  return (
    <div className="flex h-screen">
      <form
        className="grid grid-cols-1 m-auto justify-center justify-items-center space-y-5"
        onSubmit={handleSubmit}
      >
        <h1>Welcome to the GAME OF STONKS!</h1>
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
    </div>
  );
}

export default Onboard;
