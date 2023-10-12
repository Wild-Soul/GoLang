/**
* Approach -> if we're including any number say nums[i], then it's best to include all the instances of it in nums, for this create a count map.
* From this point approach is quite simple, use Priciple of inclusion and exclusion (PIE) with memoization.
*/
func deleteAndEarn(nums []int) int {
	cnt := make(map[int]int)
	for i := range nums {
		if _, ok := cnt[nums[i]]; ok {
			cnt[nums[i]]++
		} else {
			cnt[nums[i]] = 1
		}
	}

	nums = make([]int, len(cnt))
	i := 0
	for k := range cnt {
		nums[i] = k
		i++
	}
	sort.Ints(nums)
	mem := make([][]int, len(nums))
	for i := range mem {
		mem[i] = make([]int, len(nums))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var solve func(int, int) int

	solve = func(i , pi int) int {
		if i >= len(nums) {
			return 0
		}

        // if already encountered this state before, simply return precomputed result.
		if mem[i][pi] != -1 {
			return mem[i][pi]
		}

        // end state for recursion to stop.
		if i == len(nums)-1 {
			if pi == i || nums[pi]+1 != nums[i] {
				return cnt[nums[i]] * nums[i]
			}
			return 0
		}

		ans := 0
		if pi == i || nums[pi]+1 != nums[i] {
			ans = (cnt[nums[i]] * nums[i]) + solve(i+1, i)
		}

		na := solve(i+1, i+1)
		if na > ans {
			mem[i][pi] = na
			return na
		}
		mem[i][pi] = ans
		return ans
	}

	return solve(0, 0)
}

