package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Decode[T any](r *http.Request, v *T) error {
	var (
		buf = new(bytes.Buffer)
		tee = io.TeeReader(r.Body, buf)
	)
	if err := json.NewDecoder(tee).Decode(v); err != nil {
		return fmt.Errorf("failed to decode request body:[%s] %w", buf.String(), err)
	}
	return nil
}
