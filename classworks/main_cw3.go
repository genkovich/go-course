package main

import "fmt"

const DataSize = 10

type MapNode struct {
	Key         string
	Value       string
	NextMapNode *MapNode
}

func (mn *MapNode) AddNewNextNode(mapNode *MapNode) {
	nodeToPut := mn

	for {
		if nodeToPut.NextMapNode == nil {
			nodeToPut.NextMapNode = mapNode
			return
		}

		nodeToPut = nodeToPut.NextMapNode
	}

}

type HashMap struct {
	Data []*MapNode
}

func NewHashMap() *HashMap {
	return &HashMap{
		Data: make([]*MapNode, DataSize),
	}
}

func (m *HashMap) PutValue(key, value string) {
	index := getIndex(key)

	node := &MapNode{
		Key:   key,
		Value: value,
	}

	if m.Data[index] != nil {
		m.Data[index].AddNewNextNode(node)
	} else {
		m.Data[index] = node
	}
}

func (m *HashMap) GetValueByKey(key string) (string, bool) {
	index := getIndex(key)

	if m.Data[index] == nil {
		return "", false
	}

	node := m.Data[index]
	for {
		if node.Key == key {
			return node.Value, true
		}

		if node.NextMapNode == nil {
			break
		}

		node = node.NextMapNode
	}

	return m.Data[index].Value, true
}

func getIndex(key string) int {
	return int(hash(key)) % DataSize
}

func hash(key string) (hash uint8) {
	hash = 0
	for _, ch := range key {
		hash += uint8(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return
}

func main() {

	var m = NewHashMap()

	for i := 0; i < DataSize*2; i++ {
		key := fmt.Sprintf("%v-Go", i)

		m.PutValue(key, key+key)
		fmt.Println(getIndex(key))
	}

	fmt.Println(m.GetValueByKey("1-Go"))
}
