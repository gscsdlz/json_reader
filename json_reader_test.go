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
	list := reader.Get("alarm_list")
	for i := 0; i < list.Len(); i++ {
		fmt.Println(list.Get(i).Get("id").String())
	}
}
