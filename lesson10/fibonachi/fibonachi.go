package fibonachi

var fibCache map[uint64]uint64

func Fib(a uint64, useCache bool) uint64 {
	if a == 0 {
		return 0
	} else {
		if a == 1 {
			return 1
		}
	}
	if useCache {
		return FibWithCache(a-1) + FibWithCache(a-2)
	}
	return Fib(a-1, useCache) + Fib(a-2, useCache)
}

func FibWithCache(a uint64) uint64 {
	_, cacheExists := fibCache[a]
	if !cacheExists {
		fibCache[a] = Fib(a, true)
	}
	return fibCache[a]
}

func init() {
	fibCache = make(map[uint64]uint64)
}
