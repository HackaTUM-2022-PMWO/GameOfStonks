import React from "react";

export const Card = (props: React.HTMLProps<HTMLDivElement>) => {
  return (
    <div
      className={`${
        props.className ?? ""
      } bg-foreground  border border-background-200 m-10 p-10 rounded-3xl drop-shadow-2xl`}
      {...props}
    />
  );
};
