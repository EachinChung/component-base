package shutdown

import "sync"

// Callback 是您必须实现的接口。
// OnShutdown 将在请求关机时调用。 该参数是请求关机的管理器的名称。
type Callback interface {
	OnShutdown(string) error
}

// Func 是一种辅助类型，因此您可以轻松提供匿名函数作为 Callback。
type Func func(string) error

// OnShutdown 定义触发关闭时需要运行的操作。
func (f Func) OnShutdown(shutdownManager string) error {
	return f(shutdownManager)
}

// Manager 是 Manager 实现的接口。
// GetName 返回 Manager 的名字。
// ShutdownManagers 在 Start 中开始监听关机请求。
// 当他们在 GSInterface 上调用 StartShutdown 时，
// 首先调用 ShutdownStart()，然后执行所有 Callback
// 一旦所有 Callback 返回，ShutdownFinish() 就会被调用。
type Manager interface {
	GetName() string
	Start(gs GSInterface) error
	ShutdownStart() error
	ShutdownFinish() error
}

// ErrorHandler 是一个接口，你可以传递给 SetErrorHandler 来处理异步错误。
type ErrorHandler interface {
	OnError(err error)
}

// ErrorFunc 是一个辅助类型，所以你可以很容易地提供匿名函数作为 ErrorHandlers。
type ErrorFunc func(err error)

// OnError 定义了发生错误时需要运行的动作。
func (f ErrorFunc) OnError(err error) {
	f(err)
}

// GSInterface 是 GracefulShutdown 实现的接口，
// 在请求关闭时传递给 Manager 以调用 StartShutdown。
type GSInterface interface {
	StartShutdown(sm Manager)
	ReportError(err error)
	AddShutdownCallback(shutdownCallback Callback)
}

// GracefulShutdown 是处理 Callback 和 Manager。用 New 初始化它。
type GracefulShutdown struct {
	callbacks    []Callback
	managers     []Manager
	errorHandler ErrorHandler
}

// New 初始化 GracefulShutdown。
func New() *GracefulShutdown {
	return &GracefulShutdown{
		callbacks: make([]Callback, 0, 10),
		managers:  make([]Manager, 0, 3),
	}
}

// Start 在所有添加的 ShutdownManager 上启动。
// ShutdownManagers 开始监听关机请求。
// 如果任何 ShutdownManager 返回错误，则返回错误。
func (gs *GracefulShutdown) Start() error {
	for _, manager := range gs.managers {
		if err := manager.Start(gs); err != nil {
			return err
		}
	}

	return nil
}

// AddShutdownManager 添加一个将监听关机请求的管理器。
func (gs *GracefulShutdown) AddShutdownManager(manager Manager) {
	gs.managers = append(gs.managers, manager)
}

// AddShutdownCallback 添加一个回调，将会被关机请求回调
//
// 你可以提供任何实现回调接口的东西，或者你可以提供这样的函数：
//	AddShutdownCallback(shutdown.Func(func() error {
//		// callback code
//		return nil
//	}))
func (gs *GracefulShutdown) AddShutdownCallback(shutdownCallback Callback) {
	gs.callbacks = append(gs.callbacks, shutdownCallback)
}

// SetErrorHandler 设置一个 ErrorHandler，
// 当在 Callback 或 Manager 中遇到错误时将调用该错误处理程序。
//
// 你可以提供任何实现 ErrorHandler 接口的东西，或者你可以提供这样的函数：
//	SetErrorHandler(shutdown.ErrorFunc(func (err error) {
//		// handle error
//	}))
func (gs *GracefulShutdown) SetErrorHandler(errorHandler ErrorHandler) {
	gs.errorHandler = errorHandler
}

// StartShutdown 是从管理器调用的，将启动关闭。
// 首先在 Manager 上调用 ShutdownStart，调用所有 ShutdownCallbacks，
// 等待回调完成并在 Manager 上调用 ShutdownFinish。
func (gs *GracefulShutdown) StartShutdown(sm Manager) {
	gs.ReportError(sm.ShutdownStart())

	var wg sync.WaitGroup
	for _, shutdownCallback := range gs.callbacks {
		wg.Add(1)
		go func(shutdownCallback Callback) {
			defer wg.Done()

			gs.ReportError(shutdownCallback.OnShutdown(sm.GetName()))
		}(shutdownCallback)
	}

	wg.Wait()

	gs.ReportError(sm.ShutdownFinish())
}

// ReportError 是一个函数，可以用来向ErrorHandler 报告错误。它用于 Manager。
func (gs *GracefulShutdown) ReportError(err error) {
	if err != nil && gs.errorHandler != nil {
		gs.errorHandler.OnError(err)
	}
}
