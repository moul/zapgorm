package zapgorm

import (
	"fmt"
	"testing"

	"go.uber.org/zap/zaptest"
)

func Test(t *testing.T) {
	logger := zaptest.NewLogger(t)
	fmt.Println(logger)
}
