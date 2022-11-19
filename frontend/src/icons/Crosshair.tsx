import * as React from "react";
import { SVGProps } from "react";
const SvgCrosshair = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="crosshair_svg__feather crosshair_svg__feather-crosshair"
    {...props}
  >
    <circle cx={12} cy={12} r={10} />
    <path d="M22 12h-4M6 12H2M12 6V2M12 22v-4" />
  </svg>
);
export default SvgCrosshair;
