package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Gender  string  `json:"gender"`
	Hobbies []Hobby `json:"hobbies"`
}

type Hobby struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Encounter struct {
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Place       string    `json:"place"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}
