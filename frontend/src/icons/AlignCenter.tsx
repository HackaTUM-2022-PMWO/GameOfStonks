import * as React from "react";
import { SVGProps } from "react";
const SvgAlignCenter = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="align-center_svg__feather align-center_svg__feather-align-center"
    {...props}
  >
    <path d="M18 10H6M21 6H3M21 14H3M18 18H6" />
  </svg>
);
export default SvgAlignCenter;
