package request

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Gender  string   `json:"gender"`
	Hobbies []string `json:"hobbies"`
}
