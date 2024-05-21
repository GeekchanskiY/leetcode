class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        rng = len(nums)
        for i in range(rng):
            for z in range(rng):
                if nums[i] + nums[z] == target and i != z:
                    return [i,z]