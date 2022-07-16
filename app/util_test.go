package app

import "testing"

func Test_printWorkingDir(t *testing.T) {
	printWorkingDir()
}

func Test_formatBaseName(t *testing.T) {
	t.Log(formatBaseName("test"))
}
