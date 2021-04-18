package main

import (
	"context"
	"io"
	"strconv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/open-policy-agent/opa/rego"
)

func eval(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if req.Header.Get("Content-Type") != "application/json" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  body := make([]byte, length)
  length, err = req.Body.Read(body)
  if err != nil && err != io.EOF {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

	var jsonBody map[string]interface{}
  err = json.Unmarshal(body[:length], &jsonBody)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

	ignore := IsIgnore(jsonBody)
	if ignore {
    w.WriteHeader(http.StatusBadRequest)
		return
	} else {
    w.WriteHeader(http.StatusOK)
		return
	}
}

func IsIgnore(payload map[string]interface{}) bool {
	ctx := context.Background()

	r := rego.New(
		rego.Query("data.wazuh.ignore"),
		rego.Load([]string{"wazuh.rego"}, nil))

	query, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Print(err)
		return false
	}

	rs, err := query.Eval(ctx, rego.EvalInput(payload))
	if err != nil {
		log.Print(err)
		return false
	}

	ignore, ok := rs[0].Expressions[0].Value.(bool)
	if !ok {
		return false
	}

	if ignore {
		fmt.Println(rs)
		return true
	}

	return false

}

func main() {
	http.HandleFunc("/eval", eval)
	http.ListenAndServe(":8080", nil)
}
