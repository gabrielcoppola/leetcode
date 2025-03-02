from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(len(nums)):
                if i != j and (nums[i] + nums[j]) == target:
                    return [i, j]

    # Hash Tables - Uses more Memory
    def twoSum_optimzed(self, nums: List[int], target: int) -> List[int]:
        num_map = {}
        for i, num in enumerate(nums, start=0):
            complement = target - num
            if complement in num_map:
                return [num_map[complement], i]
            num_map[num] = i
        return []

teste = Solution()
print(teste.twoSum_optimzed(nums=[3,3], target=6))