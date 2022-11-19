import * as React from "react";
import { SVGProps } from "react";
const SvgArchive = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="archive_svg__feather archive_svg__feather-archive"
    {...props}
  >
    <path d="M21 8v13H3V8M1 3h22v5H1zM10 12h4" />
  </svg>
);
export default SvgArchive;
