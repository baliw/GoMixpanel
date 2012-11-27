// Unit Test code for gomixpanel.
//
// Use `go test` on the commandline to test this package
//
//
package mixpanel

import (
	"fmt"
	"testing"
)

func TestMixpanel(*testing.T) {
	fmt.Printf("Testing Mixpanel...\n")

	m := Init("", "", "")
	properties := map[string]string{
		"hello": "world",
		"red":   "blue",
	}
	m.Track("test", properties)
}
