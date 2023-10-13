import express from "express";
import * as chainController from '../controller/chain.controller'
const chainRouter = express.Router()

chainRouter.get("/all",chainController.getChains)
chainRouter.post("/new",chainController.saveNewChain)

export default chainRouter