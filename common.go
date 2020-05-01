package json_reader

type JsonReader interface {
	Get() JsonReader
	Number() float64
	String() string
}
