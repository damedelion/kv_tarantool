package dto

type Item struct {
	_msgpack struct{} `msgpack:",asArray"` //nolint: structcheck,unused

	Key   int    `json:"key"`
	Value string `json:"value"`
}
