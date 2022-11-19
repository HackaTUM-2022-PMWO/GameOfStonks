import React from "react";

export function PlayerListItem(props: { value: string }) {
  const avatar = require("./../../assets/avatar.jpg"); // TODO: replace with player images

  return (
    <li className="flex flex-row items-center space-x-4 content-center">
      <img
        className="p-1 w-10 h-10 rounded-full ring-2 ring-gray-300 dark:ring-gray-500"
        src={avatar}
        alt="avatar"
      />
      <p className="text-lg">{props.value}</p>
    </li>
  );
}
