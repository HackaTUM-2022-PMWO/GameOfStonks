import * as React from "react";
import { SVGProps } from "react";
const SvgCornerDownLeft = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="corner-down-left_svg__feather corner-down-left_svg__feather-corner-down-left"
    {...props}
  >
    <path d="m9 10-5 5 5 5" />
    <path d="M20 4v7a4 4 0 0 1-4 4H4" />
  </svg>
);
export default SvgCornerDownLeft;
