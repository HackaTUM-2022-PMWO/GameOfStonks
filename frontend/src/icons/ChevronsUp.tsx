import * as React from "react";
import { SVGProps } from "react";
const SvgChevronsUp = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="chevrons-up_svg__feather chevrons-up_svg__feather-chevrons-up"
    {...props}
  >
    <path d="m17 11-5-5-5 5M17 18l-5-5-5 5" />
  </svg>
);
export default SvgChevronsUp;
