import * as React from "react";
import { SVGProps } from "react";
const SvgCornerLeftUp = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="corner-left-up_svg__feather corner-left-up_svg__feather-corner-left-up"
    {...props}
  >
    <path d="M14 9 9 4 4 9" />
    <path d="M20 20h-7a4 4 0 0 1-4-4V4" />
  </svg>
);
export default SvgCornerLeftUp;
