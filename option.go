package httprequester

type Option interface {
	Apply(r *HTTPRequester)
}
