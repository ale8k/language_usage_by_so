import express, { Request, Response } from "express";
import { TopicController } from "./Topic/controllers/TopicController";

// import promClient from "prom-client";

const app =  express();

app.get("/test", (req, res) => {  
    res.send({message: "Connected to main route"});
})

app.get("/questions", (req: Request, res: Response) => {
    // const {topic} = req.query
    // const data =TopicController.GetDataByTopic(topic)

    res.send({Data: [], message: "Data has been sent"})
})



const PORT = process.env.PORT || 3000;

app.listen(PORT, () => console.log(`App has been registered and Listening at PORT ${PORT}`))
