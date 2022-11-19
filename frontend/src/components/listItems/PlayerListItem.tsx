import React from "react";

export function PlayerListItem (props: {key:  React.Key | null | undefined, idx: number, value: string}) {
    const avatar = require(`../../assets/avatar/${props.idx}.jpg`)
    const colors = ["#FCE700", "#EA047E", "#FF6D28", "#00F5FF", "#EB6440"]
    return(
        <li className="flex flex-row items-center space-x-4 content-center" key={props.key}>
            <img style={{}} className={`p-1 w-10 h-10 rounded-full ring-2 ring-gray-[${colors[props.idx % 5]}]`}
                 src={avatar} alt="avatar"/>
            <p>{props.value}</p>
        </li>
    )
}
