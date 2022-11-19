import { PropsWithChildren } from "react";

export const Container = (props: PropsWithChildren<{}>) => (
  <div className="p-7">{props.children}</div>
);
