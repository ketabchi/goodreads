package api

type Request struct {
	Authentication bool   `xml:"authentication"`
	Key            string `xml:"key"`
	Method         string `xml:"method"`
}
