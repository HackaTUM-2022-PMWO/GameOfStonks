import React from "react";

export type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement>;

export function Button({ className, ...props }: ButtonProps) {
  return (
    <button
      className={`${
        className ?? ""
      } bg-accent2 text-primary font-medium hover:font-bold ease-in duration-200 shadow drop-shadow py-2 px-4 rounded max-w-fit`}
      {...props}
    >
      {props.children}
    </button>
  );
}
