package json_reader

type JsonReader interface {
	Add(key interface{}, value JsonReader)
	Set(value interface{})
	Get(key interface{}) JsonReader
	Number() float64
	String() string
}
