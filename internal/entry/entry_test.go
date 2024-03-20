package entry

import (
	"fmt"
	"testing"
)

func Test_Regx(t *testing.T) {
	content := "[PERP] [DEBUG] [2024-03-19 18:00:20]: Wait NTQQ startup: dial tcp 127.0.0.1:8000: connect: connection refused"
	fmt.Println(DateTimeRegex.FindString(content))
}
