import { useEffect, useState } from "react"

const InputField = (props:InputFieldProps) => {
    const [inputText,setInputText] = useState<string>("")

    useEffect(()=>{
        props.onChange(inputText)
    },[inputText])
    return (
        <div className="w-1/3 mx-auto">
            {props.label !== null && <label htmlFor="default-input" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Default input</label>}
            <input value={inputText} onChange={(event)=>setInputText(event.target.value)} type="text" id="default-input" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"/>
        </div>
    )
}

interface InputFieldProps {
    placeholder:string
    label?:string
    onChange:(text:string)=>void
}

export default InputField;