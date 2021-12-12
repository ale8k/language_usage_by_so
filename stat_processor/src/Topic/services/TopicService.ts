import { EachMessagePayload } from "kafkajs";
import { consumer, producer } from "../../../services/kafka";
import { Topic } from "../Model/TopicModel";




export class TopicService {
    
    static getTopic = async (topicName: string, fromBeginning: boolean ): Promise<Topic[]> => {
        const TopicData: EachMessagePayload[] = []
        await producer.connect();
        await consumer.subscribe({topic: `${topicName}-questions`, fromBeginning: fromBeginning || false,})
    
        await consumer.run({
            eachMessage: (payload: EachMessagePayload): Promise<void> => {
                TopicData.push(payload);
                return Promise.resolve()
            }
        })

        return new Promise((resolve: any, reject: any) => {
            return TopicData
        })

    }

}