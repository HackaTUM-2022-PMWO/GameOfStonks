import * as React from "react";
import { SVGProps } from "react";
const SvgRewind = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="rewind_svg__feather rewind_svg__feather-rewind"
    {...props}
  >
    <path d="m11 19-9-7 9-7v14zM22 19l-9-7 9-7v14z" />
  </svg>
);
export default SvgRewind;
