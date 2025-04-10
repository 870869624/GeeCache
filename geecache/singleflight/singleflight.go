package singleflight

import "sync"

type call struct { //代表正在进行中，或已经结束的请求。
	wg  sync.WaitGroup //锁避免重入
	val interface{}
	err error
}

type Group struct { // 是 singleflight 的主数据结构，管理不同 key 的请求(call)
	mu sync.Mutex
	m  map[string]*call
}

// Do 的作用就是，针对相同的 key，无论 Do 被调用多少次，函数 fn 都只会被调用一次，等待 fn 调用结束了，返回返回值或错误。
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		//没有小组
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}

// func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
// 	if c, ok := g.m[key]; ok {
// 		c.wg.Wait()
// 		return c.val, c.err
// 	}

// 	c := new(call)
// 	c.wg.Add(1)
// 	g.m[key] = c

// 	c.val, c.err = fn()
// 	c.wg.Done()

// 	delete(g.m, key)

// 	return c.val, c.err
// }
