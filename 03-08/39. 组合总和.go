package main

func combinationSum(candidates []int, target int) (ans [][]int) {
	var backtrace func([]int, int, int, []int)
	backtrace = func(candidates []int, target int, index int, current []int) {
		if target <= 0 || index == len(candidates) {
			if target == 0 {
				ans = append(ans, append([]int{}, current...))
			}
			return
		}
		// 三种情况 第二种情况重复了
		// 1 选当前的数 index不加1
		backtrace(candidates, target-candidates[index], index, append(current, candidates[index]))
		// 2 选当前的数 index加1
		// backtrace(candidates,target-candidates[index],index+1,append(current,candidates[index]))
		// 3 不选当前的数  index+1
		backtrace(candidates, target, index+1, current)
	}
	backtrace(candidates, target, 0, []int{})
	return
}
