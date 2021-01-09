package poll

import "encoding/json"

// Encode Encode a poll
func (p Poll) Encode() ([]byte, error) {
	return json.Marshal(p)
}
