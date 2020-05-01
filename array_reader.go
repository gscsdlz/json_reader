package json_reader

type ArrayReader struct{}

func (j *ArrayReader) Get() JsonReader {
	panic("implement me")
}

func (j *ArrayReader) Number() float64 {
	panic("implement me")
}

func (j *ArrayReader) String() string {
	panic("implement me")
}
