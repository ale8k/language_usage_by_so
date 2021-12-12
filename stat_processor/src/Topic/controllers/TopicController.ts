/**
 * Define Controllers which connects to the service.
 * @TODO 
 * Get response for specified language in the request.
 * @QUESTION
 * Should the instance be declared in a common file and then exported?
 */

import { Topic } from "../Model/TopicModel";
import { TopicService } from "../services/TopicService";

export class TopicController {
    static GetDataByTopic = async(topicName: string): Promise<Topic[]> => {
        let Data: Topic[];
        // const Data: Promise<Topic[]> = await TopicService.getTopic(topicName, true)

        // return Promise.resolve().then((): Promise<Topic[]> =>  {
        //     return Data
        // })
        return Promise.resolve().then((bob: void): Topic[] => {
            return Data
        })
    }
}