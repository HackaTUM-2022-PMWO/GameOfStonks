import * as React from "react";
import { SVGProps } from "react";
const SvgXCircle = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="x-circle_svg__feather x-circle_svg__feather-x-circle"
    {...props}
  >
    <circle cx={12} cy={12} r={10} />
    <path d="m15 9-6 6M9 9l6 6" />
  </svg>
);
export default SvgXCircle;
