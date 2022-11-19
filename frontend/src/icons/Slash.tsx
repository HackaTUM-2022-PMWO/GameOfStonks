import * as React from "react";
import { SVGProps } from "react";
const SvgSlash = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="slash_svg__feather slash_svg__feather-slash"
    {...props}
  >
    <circle cx={12} cy={12} r={10} />
    <path d="m4.93 4.93 14.14 14.14" />
  </svg>
);
export default SvgSlash;
