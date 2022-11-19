import * as React from "react";
import { SVGProps } from "react";
const SvgCornerRightDown = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="corner-right-down_svg__feather corner-right-down_svg__feather-corner-right-down"
    {...props}
  >
    <path d="m10 15 5 5 5-5" />
    <path d="M4 4h7a4 4 0 0 1 4 4v12" />
  </svg>
);
export default SvgCornerRightDown;
