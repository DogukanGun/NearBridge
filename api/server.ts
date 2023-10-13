import dotenv from "dotenv"
import mongoose from "mongoose"
import app from "./src/index"

dotenv.config();

mongoose.connect(process.env.MONGO_URI ?? "mongodb://127.0.0.1:27017/test")
    .then(result=>{        
        if(result){
            void app.listen({host:process.env.HOST,port:process.env.PORT})  
            return {connectionStatus:true}
        }
    })
    .then((status)=>{
        console.log(`connection status is ${status?.connectionStatus}`);
    })
    .catch(exception=>{
        console.error(exception)
    })
