import {Routes} from "../router/router";
import {Link} from "react-router-dom";

export function RouterButton(props: {route: Routes, text: string}){
    return (
        <Link className="hover:scale-110 ease-in duration-500 bg-white drop-shadow-md hover:bg-green-4 hover:text-white text-mobile-standard lg:text-standard text-green-5 py-2 lg:py-4 px-4 lg:px-8 border border-green-5 hover:border-green-4 rounded shadow max-w-fit" target="_blank" rel="noreferrer" to={props.route}>{props.text}</Link>
    )
}