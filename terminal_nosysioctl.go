// +build windows plan9 solaris

package goterm

import (
	"syscall"

	"github.com/AllenDang/w32"
)

func getWinsize() (*winsize, error) {
	ws := new(winsize)
	sb := w32.GetConsoleScreenBufferInfo(w32.HANDLE(uintptr(syscall.Stdout)))

	ws.Col = uint16(sb.DwSize.X)
	ws.Row = uint16(sb.DwSize.Y)

	return ws, nil
}
