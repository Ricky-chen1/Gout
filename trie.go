package gout

import "strings"

type node struct {
	path   string //完整路径
	part   string //路由中某一部分
	son    []*node
	isWild bool //是否为模糊匹配,即"/:" or "/*"
}

// 找出一个匹配的节点
func (n *node) matchSon(part string) *node {
	for _, son := range n.son {
		if son.part == part || son.isWild {
			return son
		}
	}
	return nil
}

// 找出子节点中所有匹配的节点
func (n *node) matchSons(part string) []*node {
	sons := make([]*node, 0)
	for _, son := range n.son {
		if son.part == part || son.isWild {
			sons = append(sons, son)
		}
	}
	return sons
}

func (n *node) insert(path string, parts []string, height int) {
	if len(parts) == height {
		n.path = path //到达路径末尾，打上记号
		return
	}

	part := parts[height]
	son := n.matchSon(part)

	//没有子节点创建一个当前子节点
	if son == nil {
		son = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.son = append(n.son, son)
	}

	son.insert(path, parts, height+1)
}

func (n *node) find(parts []string, height int) *node {
	if height == len(parts) || strings.HasPrefix(n.part, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	sons := n.matchSons(part)

	for _, son := range sons {
		res := son.find(parts, height+1)
		if res != nil {
			return res
		}
	}

	return nil
}
