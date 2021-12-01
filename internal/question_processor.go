package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type QuestionProcessor struct {
	KafkaWriter  *kafka.Writer
	Tags         string
	CollectEvery time.Duration
}

// Queries all available pages until !HasMore
// uses a single slice to consistently append too, dynamically sizing
// throughout each iteration. It is unlikely no questions are asked,
// but in the event there are none this should go to zero value.
// 	- "We could return the current page if we wanted to be clever I guess
//	   but rate limiting on SE API prevents us. Oh well."
func (qp *QuestionProcessor) queryAllPages(dst []SOQuestionItem, currentPage int, startTime int, tags string) ([]SOQuestionItem, error) {
	data, err := GetCreatedQuestionsSync(currentPage, 100, startTime, tags)
	if err != nil {
		return nil, fmt.Errorf("something went wrong querying for statistics: %w", err)
	}
	// Should we consider a nil slice / empty slice?
	dst = append(dst, data.Items...)
	if data.HasMore {
		currentPage++
		return qp.queryAllPages(dst, currentPage, startTime, tags)
	} else {
		return dst, nil
	}
}

// Parses our SO items into kafka messages concurrency safe
// Delivers result back to a buffered channel with size of initial slice of SO items
func (qp *QuestionProcessor) parseQueryListToMessages(qryList []SOQuestionItem, mutex *sync.Mutex, ch chan<- kafka.Message, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println(len(qryList))
	for i := 0; i < len(qryList); i++ {

		data, _ := json.Marshal(qryList[i])
		ch <- kafka.Message{
			Key:   []byte(strconv.Itoa(qryList[i].QuestionID)),
			Value: data,
			// Headers: []protocol.Header{}, // maybe place tags as headers on here to differentiate / find overlaps?
		}
	}
	mutex.Unlock()
	wg.Done()
}

// Can make this a 'generic' method, but need to know more on reflection
// specially with reflection regarding slices, was only able to reflect.Slice type find
// and just wanted to move on honestly.
// When generics are released in 1.18, will solve issues like this!
func (qp *QuestionProcessor) divideQuestions(initialSlice []SOQuestionItem, size int) ([][]SOQuestionItem, int) {
	final := make([][]SOQuestionItem, 0)
	j := 0
	finalIndices := 0
	for i := 0; i < len(initialSlice); i += size {
		j += size
		if j > len(initialSlice) {
			j = len(initialSlice)
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices
		final = append(final, initialSlice[i:j])
		finalIndices += 1
	}
	return final, len(final)
}

// Ideally need a cache layer to check previous 10 minute of messages
// and to compare current id, answered and other relevant metadata worth updating
// and pushing a new message on. For now we just push everything every 10 minutes.
func (qp *QuestionProcessor) ProcessAskedQuestions(done chan struct{}) {
	tickaticka := time.NewTicker(qp.CollectEvery)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	// Wrapper to query consistently from now date
	runQuery := func() ([]SOQuestionItem, error) {
		date, _ := time.Parse("01-02-2006", time.Now().Format("01-02-2006"))
		finalDataQueryList := make([]SOQuestionItem, 0)
		finalDataQueryList, err := qp.queryAllPages(finalDataQueryList, 1, int(date.Unix()), qp.Tags)

		return finalDataQueryList, err
	}

	// Wrapper to process
	processQuestionsToKafka := func() {

		data, _ := runQuery()
		parsedMessages := make(chan kafka.Message, len(data))
		// TODO: Make batching size configurable alex
		spliced, size := qp.divideQuestions(data, 100)

		wg.Add(size)
		for i := 0; i < size; i++ {
			go qp.parseQueryListToMessages(spliced[i], &mutex, parsedMessages, &wg)
		}
		wg.Wait()
		close(parsedMessages)
		fmt.Printf("Channel size: %v \n", len(parsedMessages))

		// When would we need to actually cancel a fire-and-forget write...?
		msgSlice := make([]kafka.Message, 0)

		for msg := range parsedMessages {
			msgSlice = append(msgSlice, msg)
		}

		fmt.Printf("Messages to be sent size: %v\n", len(msgSlice))
		err := qp.KafkaWriter.WriteMessages(context.Background(), msgSlice...)
		qp.KafkaWriter.Close()
		if err != nil {
			fmt.Printf("error sending to kafka: %v\n", err)
		}
	}

	processQuestionsToKafka()
	for {
		select {
		case <-tickaticka.C:
			processQuestionsToKafka()
		case <-done:
			return
		}
	}

}
