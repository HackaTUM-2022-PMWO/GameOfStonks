import * as React from "react";
import { SVGProps } from "react";
const SvgCornerRightUp = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="corner-right-up_svg__feather corner-right-up_svg__feather-corner-right-up"
    {...props}
  >
    <path d="m10 9 5-5 5 5" />
    <path d="M4 20h7a4 4 0 0 0 4-4V4" />
  </svg>
);
export default SvgCornerRightUp;
