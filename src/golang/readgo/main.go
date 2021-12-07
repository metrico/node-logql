package main

import (
	"bufio"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strings"

	logql "github.com/metrico/loki-apache/pkg/logql"
	logLoki "github.com/metrico/loki-apache/pkg/logql/log"
	labels "github.com/prometheus/prometheus/pkg/labels"
)

var spGlobal logLoki.StreamPipeline

func main() {

	syslogger, err := syslog.New(syslog.LOG_DEBUG, "udf-reader")
	if err != nil {
		panic(err)
	}

	/*
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
	*/

	log.SetOutput(syslogger)

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		inputText := strings.Replace(text, "\n", "", -1)
		inputData := strings.SplitN(inputText, "\t", 2)

		if len(inputData) == 2 {
			inputLine := inputData[0]
			inputString := inputData[1]
			log.Println("input inputLine:: ", inputLine)
			log.Println("input inputString:: ", inputString)
			fmt.Println(Parse(inputString, inputLine))
		} else {
			fmt.Println("bad input: ", len(inputData))
		}

	}
}

//export Parse
func Parse(inputString, inputLine string) string {
	expr, err := logql.ParseLogSelector(inputString)
	if err != nil {
		log.Fatalf("LogQL Parsing error: %s", err)
	}

	p, err := expr.Pipeline()
	if err != nil {
		log.Fatalf("Pipeline error: %s", err)
	}

	//global
	if spGlobal == nil {
		spGlobal = p.ForStream(labels.Labels{})
	}
	line, lbs, ok := spGlobal.Process([]byte(inputLine))
	if !ok && lbs != nil {
		log.Fatalf("Processing error: %s, %v", err, lbs)
		line = []byte("")
	}

	log.Println("Input output: ", string(line))

	return string(line)
}
