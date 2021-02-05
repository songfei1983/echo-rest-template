package domain

type KeyValue struct {
	Key   string      `json:"key" validate:"min=3"`
	Value interface{} `json:"value"`
}
