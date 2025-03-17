package geecache

type ByteView struct {
	b []byte //存储真实的缓存值
}

// 要求被缓存对象必须实现 Value 接口，即 Len() int 方法，返回其所占的内存大小
func (v ByteView) Len() int {
	return len(v.b)
}

// b 是只读的，使用 ByteSlice() 方法返回一个拷贝，防止缓存值被外部程序修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}
func (v ByteView) String() string {
	return string(v.b)
}
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
