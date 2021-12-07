package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	logql "github.com/metrico/loki-apache/pkg/logql"
	logLoki "github.com/metrico/loki-apache/pkg/logql/log"
	labels "github.com/prometheus/prometheus/pkg/labels"
)

var spGlobal logLoki.StreamPipeline

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		inputString := strings.Replace(text, "\n", "", -1)
		//I don't know what is inputLine ?
		inputLine := ""

		fmt.Println(Parse(inputString, inputLine))
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
	if ok == false && lbs != nil {
		//log.Fatalf("Processing error: %s", err, lbs)
		line = []byte("")
	}
	return string(line)
}
