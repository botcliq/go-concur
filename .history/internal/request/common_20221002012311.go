package request

type Config struct {
	Concurrency int    `json:"concurrency"`
	Apis        []Apis `json:"apis"`
}
type Apis struct {
	Type    string   `json:"type"`
	Url     string   `json:"url"`
	Headers []Header `json:"url"`
}

type Header struct {
	Attr  string `json:"attribute"`
	Value string `json:"value"`
}
