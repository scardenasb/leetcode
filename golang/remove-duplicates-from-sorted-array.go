// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
    count := 1
    if len(nums) == 0 {
        return 0
    }
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[count] = nums[i]
            count++
        }
    }
    return count
}
