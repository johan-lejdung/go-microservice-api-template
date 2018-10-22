package goservice

// TestPayload contains the schema for the test payload consumed by the POST api
type TestPayload struct {
	Value string `json:"value"`
}
