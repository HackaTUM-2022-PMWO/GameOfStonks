import * as React from "react";
import { SVGProps } from "react";
const SvgX = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="x_svg__feather x_svg__feather-x"
    {...props}
  >
    <path d="M18 6 6 18M6 6l12 12" />
  </svg>
);
export default SvgX;
