package main

import (
	"fmt"
)

type Node struct {
	next *Node // Nodeのポインタ型
	prev *Node // Nodeのポインタ型（双方向用）
	name string
}

// （単方向）
// m->next ,n->next ,p->next
// m->next->n->next
// m->next->p->next->n->next

// （双方向）
// prev<-m->next ,prev<-n->next ,prev<-p->next
// prev<-m->next/prev<-n->next
// prev<-m->next/prev<-p->next/prev<-n->next
func (m *Node) insert(n *Node) {
	n.next = m.next
	if m.next != nil { // 双方向用
		m.next.prev = n
	}
	m.next = n
	n.prev = m // 双方向用
}

func (m *Node) erase(n *Node) {
	if n == nil {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (m *Node) printList() {
	cur := m.next

	for {
		fmt.Println("name.next", cur.name)
		if cur.next == nil {
			break
		}
		cur = cur.next
	}

	for {
		fmt.Println("name.prev", cur.name)
		if cur.prev == nil {
			break
		}
		cur = cur.prev
	}
}

func main() {
	m := map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}
	fmt.Println("map", m)

	languages := make(map[string]string)
	languages["go"] = "golang"
	languages["rb"] = "ruby"
	languages["js"] = "javascript"

	fmt.Println("languages", languages)

	for key, value := range languages { // Hashテーブルに格納されているので、要素をランダムに取得
		fmt.Println(key, value)
	}

	// 複数の異なるキーが同じバケットに入ることを衝突（collision）と呼ぶ

	// 連結リスト（単方向）
	names := [6]string{
		"yamamoto",
		"watanabe",
		"ito",
		"takahashi",
		"suzuki",
		"sato",
	}

	master := new(Node)
	master.name = "master"

	for _, name := range names {
		node := new(Node)
		node.name = name

		master.insert(node)

		if name == "ito" {
			master.erase(node)
		}
	}

	master.printList()
}
