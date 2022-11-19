import * as React from "react";
import { SVGProps } from "react";
const SvgChevronsRight = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="chevrons-right_svg__feather chevrons-right_svg__feather-chevrons-right"
    {...props}
  >
    <path d="m13 17 5-5-5-5M6 17l5-5-5-5" />
  </svg>
);
export default SvgChevronsRight;
