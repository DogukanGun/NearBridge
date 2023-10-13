"use client"

import Button from "@/components/button/Button";
import { Checkbox } from "@/components/checkbox/Checkbox";
import { DropdownOptions, DropdownSelecter } from "@/components/dropdown/DrowdownSelecter";
import InputField from "@/components/inputfield/InputField";
import { getChains } from "@/services/chainService";
import axios from "axios";
import { useEffect, useState } from "react";

const Bridge = () =>{

    const [chains,setChains] = useState<DropdownOptions[]>([])
    const [chain,setChain] = useState<DropdownOptions>()
    const [receiverContractImp,setReceiverContractImp] = useState<boolean>(false)
    const [destinationContractAddress,setDestinationContractAddress] = useState("")

    useEffect(()=>{
        getChains()
            .then((response)=>{
                if(response.length !== 0){
                    setChains(response)
                }
            })
    },[])

    const inputFieldOnChange = (text:string) =>{
        setDestinationContractAddress(text)
    }

    const chainOnChange = (text:string) =>{
        setChain(chains.filter(index=>index.chainID === text)[0])
    }

    const onChecked = (checked:boolean) =>{
        setReceiverContractImp(checked)
    }

    const createBridge = () =>{

    }

    return(
        <>
            <div className="w-full flex">
                <DropdownSelecter title="Select destination bridge" data={chains} onChange={chainOnChange}/>
            </div>
            <div className="w-full flex">
                <InputField placeholder="Destionation contract address" onChange={inputFieldOnChange}/>
            </div>
            <div className="w-full flex">
                <Checkbox label="I implemented the Receiver Contract" onChecked={onChecked}/>
            </div>
            <div className="w-full flex">
                <div className="mx-auto">
                    <Button onClick={createBridge} text="Create Bridge"/>
                </div>
            </div>
        </>
    )
}

export default Bridge;