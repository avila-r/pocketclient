package validation

type Failure struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Data    []Err  `json:"data,omitempty"`
}

type Err struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (f Failure) Error() string {
	return f.Message
}
