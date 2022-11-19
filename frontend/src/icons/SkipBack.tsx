import * as React from "react";
import { SVGProps } from "react";
const SvgSkipBack = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="skip-back_svg__feather skip-back_svg__feather-skip-back"
    {...props}
  >
    <path d="M19 20 9 12l10-8v16zM5 19V5" />
  </svg>
);
export default SvgSkipBack;
