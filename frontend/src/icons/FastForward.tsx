import * as React from "react";
import { SVGProps } from "react";
const SvgFastForward = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="fast-forward_svg__feather fast-forward_svg__feather-fast-forward"
    {...props}
  >
    <path d="m13 19 9-7-9-7v14zM2 19l9-7-9-7v14z" />
  </svg>
);
export default SvgFastForward;
