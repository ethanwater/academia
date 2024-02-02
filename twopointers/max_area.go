func maxContainerArea(height []int) int {
    var min func(a, b int) int
    min = func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }

    pointer1, pointer2 := 0, len(height)-1
    containerMax := 0

    for pointer1 < pointer2 {
        containerTemp := min(height[pointer1], height[pointer2]) * (pointer2-pointer1)
        if containerTemp > containerMax {
            containerMax = containerTemp
        }
        if height[pointer1] < height[pointer2]{
            pointer1++
        } else {
            pointer2--
        }
    }
    return containerMax
}
