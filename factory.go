package json_reader

import "encoding/json"

// data []byte JSON字符串
func Parse(data []byte) (JsonReader, error) {
	var j interface{}
	if err := json.Unmarshal(data, &j); err != nil {
		return nil, err
	}
	r := Make(j)
	return r, nil
}

func Make(i interface{}) JsonReader {
	var r JsonReader
	switch i.(type) {
	case []interface{}:
		r = new(ArrayReader)
		for k, v := range i.([]interface{}) {
			item := Make(v)
			r.Add(k, item)
		}
	case map[string]interface{}:
		r = new(ObjectReader)
		for k, v := range i.(map[string]interface{}) {
			item := Make(v)
			r.Add(k, item)
		}
	case float64:
		r = new(NumberReader)
		r.Set(i)
	case string:
		r = new(StringReader)
		r.Set(i)
	}
	return r
}
