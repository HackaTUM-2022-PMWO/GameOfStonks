import React from "react";

export type InputProps = React.InputHTMLAttributes<HTMLInputElement> & {
  label?: React.ReactNode;
};

export function Input({ label, className, ...props }: InputProps) {
  return (
    <>
      <label>
        <small className="block text-left text-l opacity-40">{label}</small>
        <input
          className={`${
            className ?? ""
          } text-secondary max-w-sm rounded-md bg-gray-100 border-2 border-gray-300 py-1 px-3"`}
          {...props}
        />
      </label>
    </>
  );
}
