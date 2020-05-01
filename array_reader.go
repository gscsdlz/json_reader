package json_reader

type ArrayReader struct {
	Data []JsonReader
}

func (j *ArrayReader) Add(key interface{}, value JsonReader) {
	idx := key.(int)
	if len(j.Data) <= idx {
		j.Data = append(j.Data, value)
	} else {
		j.Data[idx] = value
	}
}

func (j *ArrayReader) Set(value interface{}) {
	panic("implement me")
}

func (j *ArrayReader) Get(key interface{}) JsonReader {
	idx := key.(int)
	return j.Data[idx]
}

func (j *ArrayReader) Number() float64 {
	panic("implement me")
}

func (j *ArrayReader) String() string {
	panic("implement me")
}
