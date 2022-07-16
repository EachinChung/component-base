package managers

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eachinchung/component-base/shutdown"
)

// Name 定义关机管理器名称。
const Name = "PosixSignalManager"

// PosixSignalManager 实现了添加到 GracefulShutdown 的 Manager 接口。
// 使用 NewPosixSignalManager 初始化。
type PosixSignalManager struct {
	signals []os.Signal
}

// NewPosixSignalManager 初始化 PosixSignalManager。
// 作为参数，你可以提供 os.Signal-s 来监听，如果没有给出，它将默认为 SIGINT 和 SIGTERM。
func NewPosixSignalManager(sig ...os.Signal) *PosixSignalManager {
	if len(sig) == 0 {
		sig = make([]os.Signal, 2)
		sig[0] = os.Interrupt
		sig[1] = syscall.SIGTERM
	}

	return &PosixSignalManager{
		signals: sig,
	}
}

// GetName 返回此 Manager 的名称。
func (posixSignalManager *PosixSignalManager) GetName() string {
	return Name
}

// Start 开始监听 posix 信号。
func (posixSignalManager *PosixSignalManager) Start(gs shutdown.GSInterface) error {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, posixSignalManager.signals...)

		// 阻塞直到接收到信号。
		<-c

		gs.StartShutdown(posixSignalManager)
	}()

	return nil
}

// ShutdownStart 什么也不做。
func (posixSignalManager *PosixSignalManager) ShutdownStart() error {
	return nil
}

// ShutdownFinish 使用 os.Exit(0) 退出应用程序。
func (posixSignalManager *PosixSignalManager) ShutdownFinish() error {
	os.Exit(0)
	return nil
}
