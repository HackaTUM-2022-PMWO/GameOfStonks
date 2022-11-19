import React from "react";

export const Card = (props: React.HTMLProps<HTMLDivElement>) => {
  return (
    <div className={`${props.className ?? ""} card m-10 p-10`} {...props} />
  );
};
