import * as React from "react";
import { SVGProps } from "react";
const SvgCode = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="code_svg__feather code_svg__feather-code"
    {...props}
  >
    <path d="m16 18 6-6-6-6M8 6l-6 6 6 6" />
  </svg>
);
export default SvgCode;
