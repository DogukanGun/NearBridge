const Button = (props:ButtonProps) =>{
    return(
        <button type="button" onClick={()=>props.onClick()} className="text-white bg-gradient-to-br from-green-400 to-blue-600 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-green-200 dark:focus:ring-green-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2">{props.text}</button>
    )
}

interface ButtonProps{
    text:string
    onClick:()=>void
}

export default Button;