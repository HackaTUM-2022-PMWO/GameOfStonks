import React from "react";

export function Input(props: { value: string | number | readonly string[] | undefined, type: React.HTMLInputTypeAttribute | undefined, onChange?: React.ChangeEventHandler<HTMLInputElement> | undefined }) {
    return (
        <input className="bg-gray-50 text-gray-900 rounded-md px-2" id="input" type={props.type} value={props.value}
               onChange={props.onChange}/>
    )
}