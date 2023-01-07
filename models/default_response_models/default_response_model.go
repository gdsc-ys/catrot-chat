package default_response_models

type DefaultResponse200Model struct {
	Message string `json:"message"`
}

type DefaultResponse400Model struct {
	Error string `json:"error"`
}
