import React from "react";

export const Card = ({
  className,
  headline,
  ...props
}: React.HTMLProps<HTMLDivElement> & {
  headline?: string | React.ReactNode;
}) => {
  return (
    <div className="m-7">
      {typeof headline === "string" ? (
        <h2 className="text-3xl mt-5 mb-3 ml-5">{headline}</h2>
      ) : (
        headline
      )}
      <div
        className={`${
          className ?? ""
        } bg-accent2 border-accent rounded-3xl p-7`}
        {...props}
      />
    </div>
  );
};
