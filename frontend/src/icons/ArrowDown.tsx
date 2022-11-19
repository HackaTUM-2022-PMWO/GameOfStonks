import * as React from "react";
import { SVGProps } from "react";
const SvgArrowDown = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="arrow-down_svg__feather arrow-down_svg__feather-arrow-down"
    {...props}
  >
    <path d="M12 5v14M19 12l-7 7-7-7" />
  </svg>
);
export default SvgArrowDown;
