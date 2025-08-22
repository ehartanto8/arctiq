package types

type Request struct {
	Prompt string `json:"prompt"`
}

type Response struct {
	Message string `json:"message"`
}
