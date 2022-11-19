import * as React from "react";
import { SVGProps } from "react";
const SvgArrowRight = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="arrow-right_svg__feather arrow-right_svg__feather-arrow-right"
    {...props}
  >
    <path d="M5 12h14M12 5l7 7-7 7" />
  </svg>
);
export default SvgArrowRight;
