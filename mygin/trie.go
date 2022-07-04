// trie树, 实现路由前缀树，可用于参数匹配，URL通配等

package mygin

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点
	isWild   bool    // 是否模糊匹配
}

// 找到匹配的子节点，场景是用在插入时使用，找到1个匹配的就立即返回
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 节点的插入，一边匹配一边插入的方法
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		// 如果已经匹配完了，那么将pattern赋值给该node，表示它是一个完整的url，同时递归结束
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// 节点的搜索
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 递归终止
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	// 获取所有可能的子路径
	children := n.matchChildren(part)

	for _, child := range children {
		// 每条路径用下一条part查找
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

// 查找所有完整url
func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		// 一层层递归查找节点
		child.travel(list)
	}
}
