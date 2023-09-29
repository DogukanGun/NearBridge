import express from "express"
import bodyParser from "body-parser"
import chainRouter from "./router/chain.router"

const app = express()
app.use(bodyParser.json())
app.use(chainRouter)

export default app