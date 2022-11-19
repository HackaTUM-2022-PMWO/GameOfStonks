import * as React from "react";
import { SVGProps } from "react";
const SvgCloudDrizzle = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="cloud-drizzle_svg__feather cloud-drizzle_svg__feather-cloud-drizzle"
    {...props}
  >
    <path d="M8 19v2M8 13v2M16 19v2M16 13v2M12 21v2M12 15v2M20 16.58A5 5 0 0 0 18 7h-1.26A8 8 0 1 0 4 15.25" />
  </svg>
);
export default SvgCloudDrizzle;
