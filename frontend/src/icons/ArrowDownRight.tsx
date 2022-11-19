import * as React from "react";
import { SVGProps } from "react";
const SvgArrowDownRight = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="arrow-down-right_svg__feather arrow-down-right_svg__feather-arrow-down-right"
    {...props}
  >
    <path d="m7 7 10 10M17 7v10H7" />
  </svg>
);
export default SvgArrowDownRight;
