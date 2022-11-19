import * as React from "react";
import { SVGProps } from "react";
const SvgSpeaker = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={24}
    height={24}
    fill="none"
    stroke="currentColor"
    strokeWidth={2}
    strokeLinecap="round"
    strokeLinejoin="round"
    className="speaker_svg__feather speaker_svg__feather-speaker"
    {...props}
  >
    <rect x={4} y={2} width={16} height={20} rx={2} ry={2} />
    <circle cx={12} cy={14} r={4} />
    <path d="M12 6h.01" />
  </svg>
);
export default SvgSpeaker;
