import {Routes} from "../router/router";
import {Link} from "react-router-dom";

export function RouterButton(props: {route: Routes, text: string}){
    return (
        <Link className="flex bg-teal text-white font-medium hover:font-bold ease-in duration-200 drop-shadow py-2 px-4 rounded max-w-fit" to={props.route}>{props.text}</Link>
    )
}