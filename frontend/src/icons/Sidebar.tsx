import * as React from "react";
import { SVGProps } from "react";
const SvgSidebar = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="sidebar_svg__feather sidebar_svg__feather-sidebar"
    {...props}
  >
    <rect x={3} y={3} width={18} height={18} rx={2} ry={2} />
    <path d="M9 3v18" />
  </svg>
);
export default SvgSidebar;
