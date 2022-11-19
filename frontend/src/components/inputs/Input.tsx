import React from "react";

export type InputProps = React.InputHTMLAttributes<HTMLInputElement>;

export function Input({className, ...props}: InputProps) {
  return (
    <input
      className={`${
        className ?? ""
      } text-secondary max-w-sm rounded-md bg-gray-100 border border-2 border-gray-300 py-1 px-3"`}
      {...props}
    />
  );
}
