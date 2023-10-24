package models

type LiturgicalData struct {
	Year   string `json:"year"`
	Season string `json:"season"`
	Week   string `json:"week"`
	Day    string `json:"day"`
	Title  string `json:"title"`
	Psalms struct {
		Morning []string `json:"morning"`
		Evening []string `json:"evening"`
	} `json:"psalms"`
	Lessons struct {
		First  string `json:"first"`
		Second string `json:"second"`
		Gospel string `json:"gospel"`
	} `json:"lessons"`
}
