import React from "react";

export const Card = ({
  className,
  ...props
}: React.HTMLProps<HTMLDivElement>) => {
  return <div className={`${className ?? ""}  card rounded-lg m-10 p-10`} {...props} />;
};
