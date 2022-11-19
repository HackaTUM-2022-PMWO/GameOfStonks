import { Routes } from "../../router/router";
import { Link } from "react-router-dom";
import React from "react";
import { Button, ButtonProps } from "./Button";

export function RouterButton({
  route,
  ...props
}: {
  route: Routes;
} & ButtonProps) {
  return (
    <Link to={route}>
      <Button {...props} />
    </Link>
  );
}
