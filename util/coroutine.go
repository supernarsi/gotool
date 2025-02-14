package util

import (
	"context"
	"sync"
)

// GoForeach 使用协程池处理循环
// - ctx: 上下文
// - arr: 待循环处理的数据（任意类型的切片）
// - proc: 处理每个元素的逻辑函数
// - def: 协程处理异常时返回的默认值
// - nu: 协程池容量
func GoForeach[T, R any](ctx context.Context, arr []T, proc func(c context.Context, i T) (R, error), def R, nu int) []R {
	if len(arr) == 0 {
		return nil
	}

	var (
		ch     = make(chan struct{}, nu)
		wg     sync.WaitGroup
		result = make([]R, len(arr))
	)

	wg.Add(len(arr))
	for i, v := range arr {
		ch <- struct{}{}
		go func(index int, value T) {
			defer func() {
				<-ch
				wg.Done()
			}()

			if val, err := proc(ctx, value); err != nil {
				result[index] = def
			} else {
				result[index] = val
			}
		}(i, v)
	}

	wg.Wait()
	close(ch)

	return result
}
