import { DropdownOptions } from "@/components/dropdown/DrowdownSelecter"
import axios from "axios"

const chainServiceRoute = "/chain"
const getAllChains = "/all"

interface GetChainsResponse{
    data:DropdownOptions[]
}

export const getChains = async() =>{
    const response = await axios.get<GetChainsResponse>(process.env.NEXT_PUBLIC_API_URL + chainServiceRoute + getAllChains)
    return response.data.data
}