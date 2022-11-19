import * as React from "react";
import { SVGProps } from "react";
const SvgCalendar = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="calendar_svg__feather calendar_svg__feather-calendar"
    {...props}
  >
    <rect x={3} y={4} width={18} height={18} rx={2} ry={2} />
    <path d="M16 2v4M8 2v4M3 10h18" />
  </svg>
);
export default SvgCalendar;
