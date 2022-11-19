import * as React from "react";
import { SVGProps } from "react";
const SvgXSquare = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="x-square_svg__feather x-square_svg__feather-x-square"
    {...props}
  >
    <rect x={3} y={3} width={18} height={18} rx={2} ry={2} />
    <path d="m9 9 6 6M15 9l-6 6" />
  </svg>
);
export default SvgXSquare;
