package terminal

import (
	"fmt"
	"io"

	"github.com/moby/term"
)

// Size 返回用户终端的当前宽度和高度。如果不是终端，返回 nil。
// 出错时，width 和 height 返回零值。
// 通常 w 必须是 stdout。Stderr 不起作用。
func Size(w io.Writer) (width, height int, err error) {
	outFd, isTerminal := term.GetFdInfo(w)
	if !isTerminal {
		return 0, 0, fmt.Errorf("given writer is no terminal")
	}
	winSize, err := term.GetWinsize(outFd)
	if err != nil {
		return 0, 0, err
	}
	return int(winSize.Width), int(winSize.Height), nil
}
