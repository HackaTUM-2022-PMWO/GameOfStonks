import * as React from "react";
import { SVGProps } from "react";
const SvgTv = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="tv_svg__feather tv_svg__feather-tv"
    {...props}
  >
    <rect x={2} y={7} width={20} height={15} rx={2} ry={2} />
    <path d="m17 2-5 5-5-5" />
  </svg>
);
export default SvgTv;
