import * as React from "react";
import { SVGProps } from "react";
const SvgArrowDownCircle = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="arrow-down-circle_svg__feather arrow-down-circle_svg__feather-arrow-down-circle"
    {...props}
  >
    <circle cx={12} cy={12} r={10} />
    <path d="m8 12 4 4 4-4M12 8v8" />
  </svg>
);
export default SvgArrowDownCircle;
