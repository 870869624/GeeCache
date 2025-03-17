package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int            //虚拟节点倍数 replicas
	keys     []int          //哈希环 keys
	hashMap  map[int]string //虚拟节点与真实节点的映射表 hashMap，键是虚拟节点的哈希值，值是真实节点的名称
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(fmt.Sprintf("%d%s", i, key))))
			m.keys = append(m.keys, hash) //讲节点添加到哈希环中
			m.hashMap[hash] = key         //添加虚拟节点与真实节点的映射关系
			fmt.Println(m.keys, m.hashMap, "-------")
		}
	}

	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]]
}
