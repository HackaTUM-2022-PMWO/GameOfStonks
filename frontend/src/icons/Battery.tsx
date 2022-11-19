import * as React from "react";
import { SVGProps } from "react";
const SvgBattery = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="battery_svg__feather battery_svg__feather-battery"
    {...props}
  >
    <rect x={1} y={6} width={18} height={12} rx={2} ry={2} />
    <path d="M23 13v-2" />
  </svg>
);
export default SvgBattery;
