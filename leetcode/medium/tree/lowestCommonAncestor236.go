package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	hashmap := make(map[*TreeNode]*TreeNode)
	mark := make(map[*TreeNode]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			hashmap[node.Left] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			hashmap[node.Right] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		mark[p] = true
		p = hashmap[p]
	}

	for q != nil {
		if _, has := mark[q]; has {
			return hashmap[q]
		}
		q = hashmap[q]
	}

	return nil

}
