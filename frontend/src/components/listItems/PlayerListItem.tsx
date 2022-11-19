import React from "react";

export function PlayerListItem (props: {key:  React.Key | null | undefined, value: string}) {
    const avatar = require('./../../assets/avatar.jpg') // TODO: replace with player images

    return(
        <li className="flex flex-row space-x-4 content-center" key={props.key}>
            <img className="p-1 w-10 h-10 rounded-full ring-2 ring-gray-300 dark:ring-gray-500"
                 src={avatar} alt="avatar"/>
            <p>{props.value}</p>
        </li>

    )
}