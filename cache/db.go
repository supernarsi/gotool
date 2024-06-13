package cache

import (
	"context"
	"time"
)

type cCtx context.Context
type FetcherFunc[T any] func(ctx cCtx) ([]*T, error)
type GetIdFunc[T any] func(item *T) uint
type GetKeyFunc[T any] func(item *T) string

func GetIdMap[T any](ctx cCtx, cacheIns cacheHandler, cKey string, fetcher FetcherFunc[T], getId GetIdFunc[T], duration time.Duration) map[uint]*T {
	var resultMap map[uint]T

	if !getFromCache(ctx, cacheIns, cKey, &resultMap) {
		items, err := fetcher(ctx)
		if err != nil {
			return nil
		}
		resultMap = make(map[uint]T, len(items))
		for _, item := range items {
			resultMap[getId(item)] = *item
		}
		if len(resultMap) > 0 {
			setItemIdMapToCache(ctx, cacheIns, cKey, resultMap, duration)
		}
	}

	ptrMap := make(map[uint]*T, len(resultMap))
	for id, item := range resultMap {
		newItem := item
		ptrMap[id] = &newItem
	}
	return ptrMap
}

func GetKeyMap[T any](ctx cCtx, cacheIns cacheHandler, key string, fetcher FetcherFunc[T], getKey GetKeyFunc[T], duration time.Duration) map[string]*T {
	var resultMap map[string]T

	if !getFromCache(ctx, cacheIns, key, &resultMap) {
		items, err := fetcher(ctx)
		if err != nil {
			return nil
		}

		resultMap = make(map[string]T, len(items))
		for _, item := range items {
			resultMap[getKey(item)] = *item
		}
		if len(resultMap) > 0 {
			setItemKeyMapToCache(ctx, cacheIns, key, resultMap, duration)
		}
	}

	ptrMap := make(map[string]*T, len(resultMap))
	for itemKey, item := range resultMap {
		newItem := item
		ptrMap[itemKey] = &newItem
	}
	return ptrMap
}

func getFromCache[T any](ctx cCtx, cacheIns cacheHandler, key string, dest *T) bool {
	cacheVal, err := cacheIns.Get(ctx, key)
	if err != nil {
		return false
	}
	if cacheVal == nil {
		return false
	}
	if err = cacheIns.Scan(dest); err != nil {
		return false
	}
	return true
}

func setItemIdMapToCache[T any](ctx cCtx, cacheIns cacheHandler, key string, value map[uint]T, duration time.Duration) bool {
	if len(value) == 0 {
		return true
	}
	if err := cacheIns.Set(ctx, key, value, duration); err != nil {
		return false
	}
	return true
}

func setItemKeyMapToCache[T any](ctx cCtx, cacheIns cacheHandler, key string, value map[string]T, duration time.Duration) bool {
	if len(value) == 0 {
		return true
	}
	if err := cacheIns.Set(ctx, key, value, duration); err != nil {
		return false
	}
	return true
}
