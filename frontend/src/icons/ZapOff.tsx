import * as React from "react";
import { SVGProps } from "react";
const SvgZapOff = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="zap-off_svg__feather zap-off_svg__feather-zap-off"
    {...props}
  >
    <path d="M12.41 6.75 13 2l-2.43 2.92M18.57 12.91 21 10h-5.34M8 8l-5 6h9l-1 8 5-6M1 1l22 22" />
  </svg>
);
export default SvgZapOff;
