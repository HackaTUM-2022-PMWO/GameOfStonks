import { Routes } from "../../router/router";
import { Link } from "react-router-dom";
import React from "react";

export type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement>;

export function Button({ className, ...props }: ButtonProps) {
  return (
    <button
      className={`${
        className ?? ""
      } bg-background text-primary font-medium hover:font-bold ease-in duration-200 drop-shadow py-2 px-4 rounded max-w-fit`}
      {...props}
    >
      {props.children}
    </button>
  );
}
