package request

type Config struct {
	Concurrency int    `json:"concurrency"`
	Total       int    `json:"total_requests"`
	Apis        []Apis `json:"apis"`
}
type Apis struct {
	Type    string   `json:"type"`
	Url     string   `json:"url"`
	Headers []Header `json:"headers"`
	Payload string   `json:"payload"`
}

type Header struct {
	Attr  string `json:"attribute"`
	Value string `json:"value"`
}
