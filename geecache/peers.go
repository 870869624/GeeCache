package geecache

type PeerPicker interface {
	//PeerPicker 的 PickPeer() 方法用于根据传入的 key 选择相应节点 PeerGetter。
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	// PeerGetter 的 Get() 方法用于从对应 group 查找缓存值。PeerGetter 就对应于上述流程中的 HTTP 客户端。
	Get(group string, key string) ([]byte, error)
}
