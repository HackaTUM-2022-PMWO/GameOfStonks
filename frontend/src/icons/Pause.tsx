import * as React from "react";
import { SVGProps } from "react";
const SvgPause = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="pause_svg__feather pause_svg__feather-pause"
    {...props}
  >
    <path d="M6 4h4v16H6zM14 4h4v16h-4z" />
  </svg>
);
export default SvgPause;
