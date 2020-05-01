package json_reader

type NumberReader struct {
	Data float64
}

func (j *NumberReader) Add(key interface{}, value JsonReader) {
	panic("implement me")
}

func (j *NumberReader) Set(value interface{}) {
	j.Data = value.(float64)
}

func (j *NumberReader) Get(key interface{}) JsonReader {
	panic("implement me")
}

func (j *NumberReader) Number() float64 {
	return j.Data
}

func (j *NumberReader) String() string {
	panic("implement me")
}
