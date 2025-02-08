// https://leetcode.com/problems/design-a-number-container-system

type NumberContainers struct {
    Indexes map[int]int
    Reversed map[int]map[int]bool
    FindCache map[int]int
}

// https://stackoverflow.com/questions/18125625/constructors-in-go
func Constructor() NumberContainers {
    return NumberContainers{make(map[int]int), make(map[int]map[int]bool), make(map[int]int)}
}


func (this *NumberContainers) Change(index int, number int)  {
    prevValue, wasValue := this.Indexes[index]
    this.Indexes[index] = number
    curPositions, hasNumberInReverse := this.Reversed[number]
    if wasValue {
        oldReversed, _ := this.Reversed[prevValue]
        delete(oldReversed, index) 
        delete(this.FindCache, prevValue)
    }
    if hasNumberInReverse {
        curPositions[index] = true
    } else {
        tmp := make(map[int]bool)
        tmp[index] = true
        this.Reversed[number] = tmp
    }
    
    delete(this.FindCache, number)
}


func (this *NumberContainers) Find(number int) int {
    cached, isInCache := this.FindCache[number]
    if isInCache {
        return cached
    }
    curPositions, hasNumber := this.Reversed[number]
    if hasNumber {
        if len(curPositions) == 0 {
            return -1
        }
        res := math.MaxInt32
        for k, _ := range curPositions {
            if k < res {
                res = k
            }
        }
        this.FindCache[number] = res
        return res
    } else {
        return -1
    }
}


/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
