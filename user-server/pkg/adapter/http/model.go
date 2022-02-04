package http

type errorResp struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type getUserResp struct {
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	Sentence string `json:"sentence"`
}
