# go-github-events [![Go Reference](https://pkg.go.dev/badge/github.com/bored-engineer/go-github-events.svg)](https://pkg.go.dev/github.com/bored-engineer/go-github-events)
The GitHub events schema converted to strongly typed Golang structs

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	github "github.com/bored-engineer/go-github-events"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	for dec.More() {
		var raw github.Event[github.RawEvent]
		if err := dec.Decode(&raw); err != nil {
			log.Fatalf("(*json.Decoder).Decode failed: %v", err)
		}
		switch *raw.Type {
		case "PushEvent":
			event, err := github.Cast[github.PushEvent](raw)
			if err != nil {
				log.Fatalf("github.Cast[github.PushEvent] failed: %v", err)
			}
			fmt.Printf("PushEvent: %d\n", *event.Payload.DistinctSize)
		}
	}
}
```