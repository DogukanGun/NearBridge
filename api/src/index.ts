import express from "express"
import bodyParser from "body-parser"
import chainRouter from "./router/chain.router"
import cors from "cors"

const app = express()
app.use(bodyParser.json())
app.use(cors())
app.use("/chain",chainRouter)

export default app