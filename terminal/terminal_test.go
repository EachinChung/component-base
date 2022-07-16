package terminal

import (
	"fmt"
	"os"
	"testing"
)

func TestTerminalSize(t *testing.T) {
	fmt.Println(Size(os.Stdout))
}
