package model

import (
	"encoding/json"
	"fmt"
)

// Session struct
type Session_model struct {
	Userid string
	Username string
	Logined bool
}

func (s Session_model) String() string {
	return fmt.Sprintf("Userid: %s, Username: %s, Logined: %t", s.Userid, s.Username, s.Logined)
}

func (s Session_model) Json() ([]byte, error) {
	return json.Marshal(s)
}