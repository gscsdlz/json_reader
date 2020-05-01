package json_reader

type StringReader struct {
	Data string
}

func (j *StringReader) Add(key interface{}, value JsonReader) {
	panic("implement me")
}

func (j *StringReader) Set(value interface{}) {
	j.Data = value.(string)
}

func (j *StringReader) Get(key interface{}) JsonReader {
	panic("implement me")
}

func (j *StringReader) Number() float64 {
	panic("implement me")
}

func (j *StringReader) String() string {
	return j.Data
}
