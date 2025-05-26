package util

import (
	"context"
	"sync"
)

// GoForeach 使用协程池处理循环
// - ctx: 上下文
// - arr: 待循环处理的数据（任意类型的切片）
// - proc: 处理每个元素的逻辑函数
// - def: 协程处理异常或任务取消时返回的默认值
// - nu: 协程池容量 (goroutine pool size)
func GoForeach[T, R any](ctx context.Context, arr []T, proc func(c context.Context, i T) (R, error), def R, nu int) []R {
	n := len(arr)
	if n == 0 {
		return nil
	}

	// Adjust nu (concurrency level) for robustness and efficiency.
	// If nu is invalid (<=0), set it to 1 to ensure at least one worker and prevent panics.
	if nu <= 0 {
		nu = 1
	}
	// If nu is greater than the number of items, cap it at n.
	// This avoids creating an unnecessarily large semaphore channel.
	if nu > n {
		nu = n
	}

	var (
		ch     = make(chan struct{}, nu) // Semaphore channel to limit concurrency
		wg     sync.WaitGroup
		result = make([]R, n) // Pre-allocate result slice
	)

	wg.Add(n) // Add all tasks to WaitGroup upfront

	for i := 0; i < n; i++ {
		// Capture loop variables for the goroutine closure
		currentIndex := i
		currentValue := arr[i]

		// Try to acquire a semaphore slot or check for context cancellation.
		// This prevents the main goroutine from blocking indefinitely on ch <- struct{}{}
		// if the context is cancelled.
		select {
		case <-ctx.Done():
			// Context was cancelled.
			// For items from currentIndex to n-1, they won't be processed by a new goroutine.
			// Fill their results with the default value and mark them as "done" for the WaitGroup.
			for j := currentIndex; j < n; j++ {
				result[j] = def
				wg.Done() // Decrement WaitGroup for tasks that won't be launched
			}
			// Proceed to wait for already launched goroutines and then clean up.
			goto endLoop
		case ch <- struct{}{}:
			// Successfully acquired a slot from the semaphore. Launch the goroutine.
			go func(index int, value T) {
				defer func() {
					// Recover from panics within the proc function.
					if r := recover(); r != nil {
						// log.Printf("GoForeach: panic processing item %d: %v", index, r) // Example logging
						result[index] = def // Set default value on panic
					}
					<-ch      // Release the semaphore slot
					wg.Done() // Signal that this goroutine has finished
				}()

				// Double-check context before processing.
				// This handles cases where context might have been cancelled
				// after semaphore acquisition but before proc is called,
				// or if proc itself doesn't robustly check the context.
				if ctx.Err() != nil {
					result[index] = def
					return
				}

				// Execute the processing function.
				if val, err := proc(ctx, value); err != nil {
					// log.Printf("GoForeach: error processing item %d: %v", index, err) // Example logging
					result[index] = def // Set default value on error
				} else {
					result[index] = val
				}
			}(currentIndex, currentValue)
		}
	}

endLoop:
	wg.Wait() // Wait for all goroutines (either completed or marked done due to cancellation)
	close(ch) // Close the semaphore channel once all goroutines are done

	return result
}
