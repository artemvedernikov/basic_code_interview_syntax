// https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/

func findMaxFish(grid [][]int) int {
    visited := make(map[int]bool)
    var res = 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            tmp := doBfs(i, j, grid, visited)
            if (tmp > res) {
                res = tmp
            }
        }
    }

    return res
}

func doBfs(startX int, startY int, grid [][]int, visited map[int]bool) int {
    moves := [][]int{{1, 0}, {0, 1}, {-1,0}, {0, -1}}
    elem := grid[startX][startY]
    id := startX * len(grid[0]) + startY
    res := 0
    _, hasSeen := visited[id]
    if (elem > 0 && !hasSeen) {
        visited[id] = true
        res = grid[startX][startY]
        for _, move := range moves {
            updatedX := startX + move[0]
            updatedY := startY + move[1]
            if (updatedX >= 0 && updatedX < len(grid) && updatedY >= 0 && updatedY < len(grid[0])) {
                res = res + doBfs(updatedX, updatedY, grid, visited)
            }
        }
    }
    
    return res
}
