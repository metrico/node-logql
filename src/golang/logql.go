package main

import (
	"C"
	//	"fmt"
	"encoding/json"
	"log"

	logql "github.com/metrico/loki-apache/pkg/logql"
	logLoki "github.com/metrico/loki-apache/pkg/logql/log"
	labels "github.com/prometheus/prometheus/pkg/labels"
	promql_parser "github.com/prometheus/prometheus/promql/parser"
)

var spGlobal logLoki.StreamPipeline

//export ParseMetric
func ParseMetric(input *C.char) *C.char {
	inputString := C.GoString(input)
	expr, err := promql_parser.ParseMetric(inputString)

	if err != nil {
		log.Fatalf("Parsing error: %s", err)
	}

	j, err := json.Marshal(expr)
	if err != nil {
		log.Fatalf("JSON error: %s", err.Error())
	}
	return C.CString(string(j))
}

//export Parse
func Parse(input *C.char, logline *C.char) *C.char {
	inputString := C.GoString(input)
	inputLine := C.GoString(logline)
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
	if ok == false && lbs != nil {
		//log.Fatalf("Processing error: %s", err, lbs)
		line = []byte("")
	}
	return C.CString(string(line))
}

func main() {}
