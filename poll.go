package poll

// Voter interface to store polls
type Voter interface {
	Vote(string, int64) error
}

// Poll depic poll
type Poll struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Action         string   `json:"action"`
	Active         bool     `json:"active"`
	EventSubdomain string   `json:"event_subdomain"`
	Secuence       string   `json:"secuence"`
	Answers        []Answer `json:"answers"`

	voter Voter
}

// Answer poll answer
type Answer struct {
	ID          string `json:"id"`
	OptionLabel string `json:"option_label"`
	Total       int64  `json:"total"`
}
