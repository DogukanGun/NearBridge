import mongoose from "mongoose";

const Schema = mongoose.Schema;

const ChainSchema = new Schema(
    {
        chainName:{
            type:String,
            required:true
        },
        websocket:{
            type:String,
            required:true
        },
        chainID:{
            type:String,
            required:true
        },
        bridgeAddress:{
            type:String,
            required:true

        }
    }
)
const ChainModel = mongoose.model('Chain',ChainSchema)

export default ChainModel;