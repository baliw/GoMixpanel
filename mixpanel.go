// This package implements the Mixpanel API in Go
//
package mixpanel

import (
	"encoding/json"
	//"encoding/base64"
	"fmt"
)

type Mixpanel struct {
	api_key    string
	api_secret string
	api_token  string
}

type track struct {
	Event      string            `json:"event"`
	Properties map[string]string `json:"properties"`
}

func Init(api_key string, api_secret string, api_token string) *Mixpanel {
	m := &Mixpanel{
		api_key:    api_key,
		api_secret: api_secret,
		api_token:  api_token,
	}
	return m
}

func (m *Mixpanel) Track(Event string, Properties map[string]string) {
	t := &track{
		Event:      Event,
		Properties: Properties,
	}
	j, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	// temp debug output
	fmt.Printf("%s\n", j)

	// base64 encode request
	// make request
}
