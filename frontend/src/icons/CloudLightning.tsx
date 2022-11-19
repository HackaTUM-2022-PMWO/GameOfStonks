import * as React from "react";
import { SVGProps } from "react";
const SvgCloudLightning = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="cloud-lightning_svg__feather cloud-lightning_svg__feather-cloud-lightning"
    {...props}
  >
    <path d="M19 16.9A5 5 0 0 0 18 7h-1.26a8 8 0 1 0-11.62 9" />
    <path d="m13 11-4 6h6l-4 6" />
  </svg>
);
export default SvgCloudLightning;
