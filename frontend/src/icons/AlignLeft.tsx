import * as React from "react";
import { SVGProps } from "react";
const SvgAlignLeft = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="align-left_svg__feather align-left_svg__feather-align-left"
    {...props}
  >
    <path d="M17 10H3M21 6H3M21 14H3M17 18H3" />
  </svg>
);
export default SvgAlignLeft;
