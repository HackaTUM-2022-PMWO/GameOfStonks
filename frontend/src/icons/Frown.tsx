import * as React from "react";
import { SVGProps } from "react";
const SvgFrown = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="frown_svg__feather frown_svg__feather-frown"
    {...props}
  >
    <circle cx={12} cy={12} r={10} />
    <path d="M16 16s-1.5-2-4-2-4 2-4 2M9 9h.01M15 9h.01" />
  </svg>
);
export default SvgFrown;
