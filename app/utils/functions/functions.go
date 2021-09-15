package functions

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// RandCode 验证码
func RandCode(num int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var code strings.Builder
	for i := 0; i < num; i++ {
		_, _ = fmt.Fprintf(&code, "%d", numeric[rand.Intn(r)])
	}
	return code.String()
}
