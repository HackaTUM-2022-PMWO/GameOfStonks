import { Link } from "react-router-dom";
import { Airplay, AtSign, X } from "../../icons";
import SvgAtSign from "../../icons/AtSign";
import SvgChevronRight from "../../icons/ChevronRight";
import { getStonkUrl, Routes } from "../../router/router";
import { StonkInfo, StonkName } from "../../services/vo-stonks";
import { StonksAssetsMatch } from "../../assets/StonksAssetsMatch";

export type StonkPositionListProps = { stonks: Record<StonkName, number> };

// TODO: verify stonk position datatype with @bosastic
export const StonkPositionList = (props: StonkPositionListProps) => {
  return (
    <>
      <ul className="list-none">
        {Object.entries(props.stonks).map(([stonk, number], index, array) => {
          if (number === 0) {
            return null;
          }

          let img, path;
          if (typeof stonk === "string") {
            path = StonksAssetsMatch.filter((elem) => elem.stonk === stonk);
            img = path[0].img;
          }
          return (
            <Link to={getStonkUrl(stonk)}>
              <li className="flex items-center justify-between text-md gap-5 py-5 border-t-1">
                <div className="flex flex-row gap-6">
                  <img className="w-12 h-12" src={img} alt={"image"} />
                  <h3 className="flex items-center gap-1">
                    <span>{stonk}</span>
                  </h3>
                </div>
                <div className="flex items-center justify-end gap-5">
                  <span className="text-lg">
                    {number === -1 ? null : number}
                  </span>{" "}
                  <SvgChevronRight className="opacity-40" />
                </div>
              </li>
              {index < array.length - 1 && (
                <hr className="h-px bg-primary border-0 border-b-1 opacity-25" />
              )}
            </Link>
          );
        })}
      </ul>
    </>
  );
};
