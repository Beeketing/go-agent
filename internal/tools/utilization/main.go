package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Beeketing/go-agent"
	"github.com/Beeketing/go-agent/internal/utilization"
)

func main() {
	util := utilization.Gather(utilization.Config{
		DetectAWS:    true,
		DetectAzure:  true,
		DetectDocker: true,
		DetectPCF:    true,
		DetectGCP:    true,
	}, newrelic.NewDebugLogger(os.Stdout))

	js, err := json.MarshalIndent(util, "", "\t")
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%s\n", js)
	}
}
