import express from "express";
import {Kafka} from "kafkajs";
import promClient from "prom-client";

const kafka = new Kafka({
    clientId: "stat_processor",
    brokers: ["127.0.0.1:9093"]
})

const producer = kafka.producer();
const consumer = kafka.consumer({groupId:  "stat_processor"});

const run = async () => {
    await producer.connect();

    await consumer.connect()
    await consumer.subscribe({topic: "python-questions", fromBeginning: true})

    await consumer.run({
        eachMessage: async (message) => {
            console.log(message)
        }
    })
}


const app =  express();

app.get("/main", (req, res) => {
    run()
    res.send({message: "Connected to main route"});
})


const PORT = process.env.PORT || 3100;

app.listen(PORT)