package util

import (
	"context"
	"sync"
)

// GoForeach 使用协程池处理循环
// -ctx: 上下文
// -arr: 待循环处理的数据（任意类型的切片）
// -proc: 处理每个元素的逻辑函数
// -def: 协程处理异常时返回的默认值
// -nu: 协程池容量
func GoForeach[T, R any](ctx context.Context, arr []T, proc func(c context.Context, i T) (R, error), def R, nu int) []R {
	ch := make(chan struct{}, nu)
	wg := sync.WaitGroup{}
	total := len(arr)
	result := make([]R, total)
	if total == 0 {
		return result
	}

	wg.Add(total)
	for i := range arr {
		ch <- struct{}{}
		go func(k int) {
			defer func() {
				<-ch
				wg.Done()
			}()

			// 执行用户逻辑
			if val, err := proc(ctx, arr[k]); err != nil {
				result[k] = def
			} else {
				result[k] = val
			}
		}(i)
	}

	wg.Wait()
	close(ch)

	return result
}
