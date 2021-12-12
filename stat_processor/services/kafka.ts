import {Kafka} from "kafkajs";


const kafka = new Kafka({
    clientId: "stat_processor",
    brokers: ["127.0.0.1:9093"]
})

export const producer = kafka.producer();
export const consumer = kafka.consumer({groupId:  "stat_processor"});