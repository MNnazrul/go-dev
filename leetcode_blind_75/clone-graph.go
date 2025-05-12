/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    vis := make(map[*Node]*Node)

    var dfs func(cur *Node) *Node 
    dfs = func(cur *Node) *Node {
        if copy, ok := vis[cur]; ok {
            return copy
        }
        copy := &Node{Val: cur.Val}
        vis[cur] = copy

        for _, neighbors := range cur.Neighbors {
            copy.Neighbors = append(copy.Neighbors, dfs(neighbors))
        } 
        return copy
    }

    return dfs(node)
}
