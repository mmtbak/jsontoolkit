package jsonnode

type Node struct {
}

func (n *Node) Get(key string) *Node {
	return nil
}

func (n *Node) GetInt(key string) int {
	return 0
}

func ParseJson(data []byte) *Node {
	return nil
}
