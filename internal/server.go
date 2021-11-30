package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"strconv"

	"github.com/ale8k/language_usage_by_so/internal/handlers"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

var Server *http.Server

func StartServer() {
	done := make(chan struct{})
	phpQp := &QuestionProcessor{
		KafkaWriter: kafka.Writer{
			Addr:         kafka.TCP("kafka:9092"),
			Topic:        "php-questions",
			Balancer:     &kafka.LeastBytes{},
			MaxAttempts:  30,
			BatchSize:    5, // what if < 5 were asked?
			BatchBytes:   1048576,
			BatchTimeout: time.Duration(time.Second * 5), // we don't care that much honestly
			ReadTimeout:  10,                             // what is this?
			WriteTimeout: 10,
			RequiredAcks: 1,
			Async:        false,
			Completion: func(km []kafka.Message, err error) {
				if err != nil {
					log.Printf("kafka failed message deliver: %v", err)
				}
			},
			Compression: compress.Gzip,
			Logger:      nil,
			ErrorLogger: kafka.LoggerFunc(func(s string, i ...interface{}) {
				log.SetPrefix("KAFKA: ")
				log.Printf("failed, %s, see: %v", s, i)
			}),
			// Transport:   nil, create a kafka.Transport for TLS config
		},
		Tags:         "php",
		CollectEvery: time.Duration(time.Minute * 10),
	}
	go phpQp.ProcessAskedQuestions(done)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 9000
	}
	pprofEndpoints, err := strconv.ParseBool(os.Getenv("PPROF_ENDPOINTS"))
	if err != nil {
		log.Printf("Failed to set PPROF_ENDPOINTS")
		pprofEndpoints = false
	}

	mux := http.NewServeMux()

	if pprofEndpoints {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	Server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "", strconv.Itoa(port)),
		Handler: mux,
	}

	handlers.RegisterAllHandlers(mux, Server)

	log.Printf("Starting server on port %v... \n", port)
	log.Fatalf("Server failed to start: %v", Server.ListenAndServe())
}
