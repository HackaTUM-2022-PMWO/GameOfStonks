import * as React from "react";
import { SVGProps } from "react";
const SvgSkipForward = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="skip-forward_svg__feather skip-forward_svg__feather-skip-forward"
    {...props}
  >
    <path d="m5 4 10 8-10 8V4zM19 5v14" />
  </svg>
);
export default SvgSkipForward;
