import { Link } from "react-router-dom";
import { Airplay, AtSign, X } from "../../icons";
import SvgAtSign from "../../icons/AtSign";
import SvgChevronRight from "../../icons/ChevronRight";
import { getStonkUrl, Routes } from "../../router/router";
import {StonkInfo, StonkName} from "../../services/vo-stonks";

export type StonkPositionListProps = { stonks: (StonkInfo[] | StonkName[]) };

// TODO: verify stonk position datatype with @bosastic
export const StonkPositionList = (props: StonkPositionListProps) => {
  return (
    <>
      <h2>Stonks</h2>
      <ul className="list-none">
        {props.stonks.map((stonk) => (
          <Link to={typeof stonk === "string" ? getStonkUrl(stonk) : getStonkUrl(stonk.Name)}>
            <li className="flex items-center justify-between text-lg gap-5 py-5 border-t-1">
              <span className="flex items-center gap-1">
                <span>{typeof stonk === "string" ? stonk : stonk.Name}</span>
              </span>
              <div className="flex items-center justify-end gap-5">
                <span className="flex items-center gap-1">
                  <span className="opacity-40">
                    <X />
                  </span>
                  <span>0</span>
                </span>

                <SvgChevronRight className="opacity-40" />
              </div>
            </li>
          </Link>
          // <li key={stonk.Name} className="p-4 border-t-1">
          //   {stonk.Name}
          // </li>
        ))}
        <Link to={getStonkUrl("pencil")}>
          <li className="flex items-center justify-between text-lg gap-5 py-5 border-t-1">
            <span className="flex items-center gap-1">
              {/* <b>{<SvgAtSign />}</b> */}
              <span>test</span>
            </span>
            <div className="flex items-center justify-end gap-5">
              <span className="flex items-center gap-1">
                <span className="opacity-40">
                  <X />
                </span>
                <span>5</span>
              </span>

              <SvgChevronRight className="opacity-40" />
            </div>
          </li>
        </Link>
      </ul>
    </>
  );
};
