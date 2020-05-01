package json_reader

type ObjectReader struct {
	Data map[string]JsonReader
}

func (j *ObjectReader) Add(key interface{}, value JsonReader) {
	if j.Data == nil {
		j.Data = make(map[string]JsonReader)
	}
	idx := key.(string)
	j.Data[idx] = value
}

func (j *ObjectReader) Set(value interface{}) {
	panic("implement me")
}

func (j *ObjectReader) Get(key interface{}) JsonReader {
	idx := key.(string)
	return j.Data[idx]
}

func (j *ObjectReader) Number() float64 {
	panic("implement me")
}

func (j *ObjectReader) String() string {
	panic("implement me")
}
