import * as React from "react";
import { SVGProps } from "react";
const SvgPercent = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="percent_svg__feather percent_svg__feather-percent"
    {...props}
  >
    <path d="M19 5 5 19" />
    <circle cx={6.5} cy={6.5} r={2.5} />
    <circle cx={17.5} cy={17.5} r={2.5} />
  </svg>
);
export default SvgPercent;
