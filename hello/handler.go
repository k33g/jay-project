package function

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	who string `json:"who"`
	what string `json:"what"`
	text string `json:"text"`
}

// Handle a serverless request...
func Handle(req []byte) string {

	message := json.Marshal(Message{
		who: "Jay",
		what: "🤖",
		text: "👋 I ❤️ OpenFaaS"
	})

	return string(message)
	//return fmt.Sprintf("%s", string(message))
}
