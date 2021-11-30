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

func processorWrapper(topicname string, tags string) {
	done := make(chan struct{})
	phpQp := &QuestionProcessor{
		KafkaWriter: kafka.Writer{
			Addr:         kafka.TCP("kafka:9092"),
			Topic:        topicname,
			Balancer:     &kafka.LeastBytes{},
			MaxAttempts:  30,
			BatchSize:    5, // what if < 5 were asked?
			BatchBytes:   1048576,
			BatchTimeout: time.Duration(time.Second * 5), // we don't care that much honestly
			ReadTimeout:  10,                             // what is this?
			//WriteTimeout: 10,
			RequiredAcks: 1,
			Async:        false,
			Completion: func(km []kafka.Message, err error) {
				if err != nil {
					log.Printf("kafka failed message deliver: %v", err)
				} else {
					log.SetPrefix("KAFKA: ")
					log.Printf("Batch sent successfully of size: %v\n", len(km))
				}
			},
			Compression: compress.Gzip,
			ErrorLogger: kafka.LoggerFunc(func(s string, i ...interface{}) {
				log.SetPrefix("KAFKA: ")
				log.Printf("failed, %s, see: %v", s, i)
			}),
			// Transport:   nil, create a kafka.Transport for TLS config
		},
		Tags:         tags,
		CollectEvery: time.Duration(time.Minute * 10),
	}
	go phpQp.ProcessAskedQuestions(done)
}

func StartServer() {
	processorWrapper("php-questions", "php")
	processorWrapper("go-questions", "go")
	processorWrapper("java-questions", "java")
	processorWrapper("javascript-questions", "javascript")
	processorWrapper("typescript-questions", "typescript") // alex: top
	processorWrapper("nodejs-questions", "node.js")
	processorWrapper("ruby-questions", "ruby")
	processorWrapper("python-questions", "python") // chloe: top
	processorWrapper("csharp-questions", "c#")
	processorWrapper("dotnet-questions", ".net")
	processorWrapper("aspnet-questions", "asp.net")
	processorWrapper("cpp-questions", "c++")
	processorWrapper("c-questions", "c")
	processorWrapper("matlab-questions", "matlab")
	processorWrapper("sql-questions", "sql")
	processorWrapper("rust-questions", "rust")
	processorWrapper("haskell-questions", "haskell")
	processorWrapper("lua-questions", "lua")       // chloe: bottom
	processorWrapper("elixir-questions", "elixir") // alex: bottom
	processorWrapper("php-questions", "php")
	processorWrapper("swift-questions", "swift")
	processorWrapper("kotlin-questions", "kotlin")
	processorWrapper("scala-questions", "scala")
	processorWrapper("html-questions", "html")

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
