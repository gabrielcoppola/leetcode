from typing import List

class Solution:
    def findClosestNumber(self, nums: List[int]) -> int:
        close = nums[0]
        for num in nums:
            if abs(num) < abs(close):
                close = num
        for num in nums:
            if abs(close) == abs(num) and close < num:
                return num
        return close

solution = Solution()
print(solution.findClosestNumber(nums=[-4,-2,1,4,8]))