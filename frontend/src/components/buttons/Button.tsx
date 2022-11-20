import React from "react";

export type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement>;

export function Button({ className, ...props }: ButtonProps) {
  return (
    <button
      className={`bg-accent2 text-primary font-medium hover:font-bold ease-in duration-200 shadow drop-shadow py-5 px-7 rounded-2xl max-w-fit ${
        className ?? ""
      }`}
      {...props}
    >
      {props.children}
    </button>
  );
}
