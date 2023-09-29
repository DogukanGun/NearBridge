import app from "./index"

import mongoose from "mongoose"


mongoose.connect(process.env.MONGO_URI)
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
