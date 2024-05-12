package htask

// RunTask 运行一个循环任务, 此任务有如下特征
// 1. 在第一次运行的时候, 会同步执行, 此时可以阻塞等待 error
// 2. 从第二运行开始会启动协程异步执行
// 适用于如下场景:
// 1. 耗时的任务执行, 但首次执行时的报错需要返回给前端页面用作提示
func RunTask(call func(errCh chan error)) error {
	errCh := make(chan error)
	finishCh := make(chan struct{})
	go func() {
		call(errCh)
		finishCh <- struct{}{}
		close(errCh)
		close(finishCh)
	}()
	// 首次执行
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-finishCh:
		return nil
	}
	// 后续开启协程执行
	go func() {
		for {
			select {
			case err := <-errCh:
				if err != nil {
					return
				}
			case <-finishCh:
				return
			}
		}
	}()
	return nil
}
