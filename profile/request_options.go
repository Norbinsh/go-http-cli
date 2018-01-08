package profile

// RequestOptions is a representation of a request that can be loaded from a profile.
type RequestOptions struct {
	Body    string
	Headers map[string][]string
	Method  string
	URL     string
}
