import * as React from "react";
import { SVGProps } from "react";
const SvgCornerUpLeft = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="corner-up-left_svg__feather corner-up-left_svg__feather-corner-up-left"
    {...props}
  >
    <path d="M9 14 4 9l5-5" />
    <path d="M20 20v-7a4 4 0 0 0-4-4H4" />
  </svg>
);
export default SvgCornerUpLeft;
