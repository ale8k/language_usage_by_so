package internal

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

// Queries all available pages until !HasMore
// uses a single slice to consistently append too, dynamically sizing
// throughout each iteration. It is unlikely no questions are asked,
// but in the event there are none this should go to zero value.
// 	- "We could return the current page if we wanted to be clever I guess
//	   but rate limiting on SE API prevents us. Oh well."
func queryAllPages(dst []SOQuestionItem, currentPage int, startTime int) ([]SOQuestionItem, error) {
	data, err := GetCreatedQuestionsSync(currentPage, 100, startTime, "go")
	if err != nil {
		return nil, fmt.Errorf("something went wrong querying for statistics: %w", err)
	}
	// Should we consider a nil slice / empty slice?
	dst = append(dst, data.Items...)
	if data.HasMore {
		currentPage++
		return queryAllPages(dst, currentPage, startTime)
	} else {
		return dst, nil
	}
}

func parseQueryListToMessages(qryList []SOQuestionItem, mutex *sync.Mutex, ch chan<- kafka.Message, wg *sync.WaitGroup) {
	mutex.Lock()
	for i := 0; i < len(qryList); i++ {
		data, _ := json.Marshal(qryList[i])
		ch <- kafka.Message{
			Key:   []byte(strconv.Itoa(qryList[i].QuestionID)),
			Value: data,
			// Headers: []protocol.Header{},
		}
	}
	mutex.Unlock()
	wg.Done()
}

func RenameMe(done chan struct{}) {
	date, _ := time.Parse("01-02-2006", time.Now().Format("01-02-2006"))
	finalDataQueryList := make([]SOQuestionItem, 0)
	finalDataQueryList, err := queryAllPages(finalDataQueryList, 1, int(date.Unix()))

	if err != nil {
		// skip this push
	}

	parsedMessages := make(chan kafka.Message, len(finalDataQueryList))
	// send channel to parse query list and have it fill buffer
	// split it up in divisions of 4 perhaps, idk, break slice down and copy
	// it is crucial u copy alex otherwise race conditions, or wrap inside parse in mutex?
	// mutex probably easily
	var mutex *sync.Mutex
	var wg *sync.WaitGroup
	wg.Add(1)
	parseQueryListToMessages(finalDataQueryList, mutex, parsedMessages, wg)
	wg.Wait()
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "test-topic", 0)

	// bytesWritten, err := conn.WriteMessages(msgBatch...)

	// log.Println(bytesWritten)

	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		fmt.Println("TODO")
	// 	case <-done:
	// 		return
	// 	}
	// }
}
