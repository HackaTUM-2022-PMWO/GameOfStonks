import { Routes } from "../../router/router";
import { Link } from "react-router-dom";
import React from "react";

export type ButtonProps = Omit<
  React.ButtonHTMLAttributes<HTMLButtonElement>,
  "className"
>;

export function Button(props: ButtonProps) {
  return (
    <button
      className="bg-teal text-white font-medium hover:font-bold ease-in duration-200 drop-shadow py-2 px-4 rounded max-w-fit"
      {...props}
    >
      {props.children}
    </button>
  );
}
