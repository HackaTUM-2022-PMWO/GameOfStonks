import React from "react";

export type InputProps = React.InputHTMLAttributes<HTMLInputElement>;

export function Input(props: InputProps) {
  return (
    <input
      className={`${
        props.className ?? ""
      } bg-gray-50 text-gray-900 rounded-md px-2"`}
      {...props}
    />
  );
}
