package json_reader

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	reader, err := Parse([]byte(`{
	"alarm_list": [{
			"id": "123"
		}, {
			"id": "456"
		}
	]
}`))

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reader.Get("alarm_list").Get(0).Get("id").String())
}
