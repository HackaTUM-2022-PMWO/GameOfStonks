import React from "react";

const colors = [
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#FCE700]",
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#EA047E]",
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#FF6D28]",
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#00F5FF]",
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#EB6440]",
];

export function PlayerListItem(props: {
  key: React.Key | null | undefined;
  idx: number;
  value: string;
}) {
  const avatar = require(`../../assets/avatar/${props.idx}.jpg`);

  return (
    <li className="flex flex-row items-center space-x-4 content-center">
      <img className={colors[props.idx]} src={avatar} alt="avatar" />
      <p>{props.value}</p>
    </li>
  );
}
