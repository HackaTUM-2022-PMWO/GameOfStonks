import React from "react";

export const colorsForIndex = [
  "#FCE700",
  "#EA047E",
  "#FF6D28",
  "#00F5FF",
  "#EB6440",
];

const colors = [
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#FCE700]",
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#EA047E]",
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#FF6D28]",
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#00F5FF]",
  "p-1 w-10 h-10 rounded-full ring-2 ring-[#EB6440]",
];

export function PlayerListItem(props: { idx: number; value: string }) {
  const avatar = `/avatar/${props.idx}.jpg`;

  return (
    <li className="flex inline-flex flex-row bg-foreground rounded-3xl p-2 items-center space-x-3 px-3 pr-5 content-center">
      <img
        className={
          "p-1 w-10 h-10 rounded-full ring-2 ring-[" + colors[props.idx]
        }
        src={avatar}
        alt="avatar"
      />
      <p>{props.value}</p>
    </li>
  );
}

export function PlayerTag(props: { idx: number; value: string }) {
  const avatar = `/avatar/${props.idx}.jpg`;

  return (
    <span className="inline-flex flex-row bg-foreground rounded-3xl p-1 items-center space-x-2 px-1 pr-3 content-center">
      <img
        className={"p-1 w-7 h-7 rounded-full ring-2 ring-[" + colors[props.idx]}
        src={avatar}
        alt="avatar"
      />
      <p>{props.value}</p>
    </span>
  );
}
