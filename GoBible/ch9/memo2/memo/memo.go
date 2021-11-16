package memo

import "sync"

type Memo struct {
	f     Func
	cache map[string]result
	mu    sync.Mutex
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// not concurrency-safe !
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	defer memo.mu.Unlock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
