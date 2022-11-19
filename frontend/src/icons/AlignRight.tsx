import * as React from "react";
import { SVGProps } from "react";
const SvgAlignRight = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="align-right_svg__feather align-right_svg__feather-align-right"
    {...props}
  >
    <path d="M21 10H7M21 6H3M21 14H3M21 18H7" />
  </svg>
);
export default SvgAlignRight;
