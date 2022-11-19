import * as React from "react";
import { SVGProps } from "react";
const SvgVolumeX = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="volume-x_svg__feather volume-x_svg__feather-volume-x"
    {...props}
  >
    <path d="M11 5 6 9H2v6h4l5 4V5zM23 9l-6 6M17 9l6 6" />
  </svg>
);
export default SvgVolumeX;
