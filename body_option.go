package httprequester

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

type BodyReaderOption struct {
	body io.Reader
}

func BodyReader(body io.Reader) *BodyReaderOption {
	return &BodyReaderOption{
		body: body,
	}
}

func (o *BodyReaderOption) Apply(req *HTTPRequester) {
	req.body = o.body
}

func RequestBody(body any) *BodyReaderOption {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		err = errors.Join(errors.New("error marshaling body to JSON in http request"), err)
		panic(err)
	}

	return BodyReader(&buf)
}
