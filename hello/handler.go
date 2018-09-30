package function

import (
	"fmt"
)

// Handle a serverless request...
func Handle(req []byte) string {

	//return string(message)
	return fmt.Sprintf("👋 🌍 I ❤️ OpenFaaS")
}
