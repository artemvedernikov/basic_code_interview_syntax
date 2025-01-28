// https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
    //res := twoSumNestedLoop(nums, target)
    res := twoSumHashMap(nums, target)
    return res 
}

func twoSumNestedLoop(nums []int, target int) []int {
    for i:=0; i < len(nums); i++ {
        for j:=i+1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                res:= []int{i, j}
                return res
            }
        }
    }
    res:= []int{}
    return res
}

func twoSumHashMap(nums []int, target int) []int {
    m := make(map[int]int)
    for i:= 0; i < len(nums); i++ {
        m[nums[i]] = i
    }

    for i:= 0; i < len(nums); i++ {
        j, ok := m[target - nums[i]]
        if ok && i != j {
            return []int{i, j}
        }
    } 
    return []int{}
}    
