package internal

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
)

func Test_queryAllPages(t *testing.T) {
	type args struct {
		dst         []SOQuestionItem
		currentPage int
		startTime   int
		tags        string
	}
	tests := []struct {
		name    string
		args    args
		want    []SOQuestionItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := queryAllPages(tt.args.dst, tt.args.currentPage, tt.args.startTime, tt.args.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryAllPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryAllPages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseQueryListToMessages(t *testing.T) {
	type args struct {
		qryList []SOQuestionItem
		mutex   *sync.Mutex
		ch      chan<- kafka.Message
		wg      *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseQueryListToMessages(tt.args.qryList, tt.args.mutex, tt.args.ch, tt.args.wg)
		})
	}
}

func Test_divideQuestions(t *testing.T) {
	type args struct {
		initialSlice []SOQuestionItem
		size         int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]SOQuestionItem
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := divideQuestions(tt.args.initialSlice, tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideQuestions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("divideQuestions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProcessAskedQuestions(t *testing.T) {
	type args struct {
		done         chan struct{}
		tags         string
		collectEvery time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ProcessAskedQuestions(tt.args.done, tt.args.tags, tt.args.collectEvery)
		})
	}
}
