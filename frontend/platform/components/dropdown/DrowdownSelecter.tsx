export const DropdownSelecter = (props:DropdownSelecterProps) =>{

    return (
        <div className="w-1/3 mx-auto">
            <label htmlFor="selectboxes" className="block mb-2 text-sm font-medium text-white">{props.title}</label>
            <select onChange={(event)=>props.onChange(event.target.value)} id="selectboxes" className="bg-gray-50 text-white border border-gray-300 :text-white text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-400 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
            {
                props.data.map((index)=>{
                    return (<option className="text-white" value={index.chainID}>{index.chainName}</option>)
                })
            }
            </select>
        </div>
    )
}


interface DropdownSelecterProps{
    data:DropdownOptions[]
    title:string
    onChange:(value:string)=>void
}

export interface DropdownOptions{
    chainName:string
    chainID:string
}