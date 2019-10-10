/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {

    if p == nil && q == nil{ // Both are nil, ended at same time, good. Only time they match are when we get through entire tree to all nil and they all match
        return true
    }

    if q == nil || p == nil { // One is nil but one isnt
        return false
    }

    if p.Val != q.Val{ // Current values dont match
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
