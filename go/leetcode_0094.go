// https://leetcode.com/problems/binary-tree-inorder-traversal/?envType=problem-list-v2&envId=stack

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    //return inorderRecursive(root)
    return inorderIterative(root)
}

func inorderIterative(root *TreeNode) []int {
    result := []int{}

    if root == nil {
        return []int{}
    }

    stack := []*TreeNode{}
    stack = append(stack, root)
    down := true

    for len(stack) > 0 {
        head := stack[len(stack) - 1]
        if head.Left != nil && down {
            down = true
            stack = append(stack, head.Left)
        } else {
            down = false
            result = append(result, head.Val)
            stack = stack[:len(stack) - 1]
            if head.Right != nil {
                down = true
                stack = append(stack, head.Right)
            }
        
        }
    }

    return result
}

func inorderRecursive(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    result := []int{}
    if root.Left != nil {
        result = inorderTraversal(root.Left)
    }
    result = append(result, root.Val)
    if (root.Right != nil) {
        rightTraversal := inorderTraversal(root.Right)
        result = append(result, rightTraversal...)
    }
    return result

}
