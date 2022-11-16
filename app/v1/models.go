package v1

type HttpMessageMetadata struct {
	CorrelationId string `json:"correlation,omitempty"`
	SpanId        string `json:"span_id,omitempty"`
}

type LoginUserMessage struct {
	UserId       string              `json:"user_id"`
	UserPassword string              `json:"user_password"`
	MetaData     HttpMessageMetadata `json:"meta,omitempty"`
}
