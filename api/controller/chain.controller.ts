import { Request, Response, NextFunction } from "express"
import ChainModel from "../data/chain.data"

interface ChainBody{
    chainName:string
    chainID:string
    websocket:string
}

export const saveNewChain = (req:Request<any,Response,ChainBody>, res:Response, next:NextFunction) =>{
    ChainModel.create({chainName:req.body.chainName,chainID:req.body.chainID,websocket:req.body.websocket})
        .then((response)=>{
            if(response){
                res.sendStatus(200)
            }
        })
}

export const getChains = (req:Request, res:Response, next:NextFunction) =>{
    ChainModel.find()
        .then((response)=>{
            if(response){
                res.json({data:response}).sendStatus(200)
            }
        })
}