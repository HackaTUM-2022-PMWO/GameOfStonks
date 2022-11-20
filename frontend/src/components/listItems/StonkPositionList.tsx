import {Link} from "react-router-dom";
import {Airplay, AtSign, X} from "../../icons";
import SvgAtSign from "../../icons/AtSign";
import SvgChevronRight from "../../icons/ChevronRight";
import {getStonkUrl, Routes} from "../../router/router";
import {StonkInfo, StonkName} from "../../services/vo-stonks";
import {StonksAssetsMatch} from "../../assets/StonksAssetsMatch";

export type StonkPositionListProps = { stonks: (StonkInfo[] | StonkName[]) };

// TODO: verify stonk position datatype with @bosastic
export const StonkPositionList = (props: StonkPositionListProps) => {
    return (
        <div>
            <h2>Stonks</h2>
            <ul className="list-none">
                {props.stonks.map((stonk, index, array) => {
                    let img, path;
                    if (typeof stonk === "string") {
                        path = StonksAssetsMatch.filter(elem => elem.stonk === stonk)
                    } else {
                        path = StonksAssetsMatch.filter(elem => elem.stonk === stonk.Name)
                    }
                    img = path[0].img;
                    return (<Link to={typeof stonk === "string" ? getStonkUrl(stonk) : getStonkUrl(stonk.Name)}>
                            <li className="flex items-center justify-between text-lg gap-5 py-5 border-t-1">
                                <div className="flex flex-row gap-6">
                                    <img className="w-12 h-12" src={img} alt={"image"}/>
                                    <h3 className="flex items-center gap-1">
                                        <span>{typeof stonk === "string" ? stonk : stonk.Name}</span>
                                    </h3>
                                </div>
                                <div className="flex items-center justify-end gap-5">
                <span className="flex items-center gap-1">
                  <span className="opacity-40">
                    <X/>
                  </span>
                  <span>0</span>
                </span>
                                    <SvgChevronRight className="opacity-40"/>
                                </div>
                            </li>
                        <hr className="h-px bg-gray-400 border-0"/>
                        </Link>
                    )
                })}
            </ul>
        </div>
    );
};
