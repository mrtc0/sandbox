package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"context"
	"encoding/json"

	"github.com/hpcloud/tail"
	"github.com/open-policy-agent/golang-opa-wasm/opa"
)

func main() {
	var alert interface{} = map[string]interface{}{}

	policy, err := ioutil.ReadFile("./policy.wasm")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	rego, err := opa.New().WithPolicyBytes(policy).Init()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	defer rego.Close()

	ctx := context.Background()

	t, err := tail.TailFile("./alerts.json", tail.Config{Follow: true, ReOpen: true})
	for line := range t.Lines {
		err = json.Unmarshal([]byte(line.Text), &alert)
		if err != nil {
			os.Exit(1)
		}

		resultBool, err := opa.EvalBool(ctx, rego, &alert)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("ignore? -> %#v\n", resultBool)
	}
}
