from typing import List


class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        answer = []
        start = 0
        for i, n in enumerate(nums):
            if i != len(nums) - 1 and nums[i + 1] - 1 == n:
                continue
            if i != start:
               answer.append(f"{nums[start]}->{n}")
            else:
               answer.append(str(n))
            start = i + 1
        return answer


if __name__ == "__main__":
    examples = [([0, 1, 2, 4, 5, 7], ["0->2", "4->5", "7"])]
    solution = Solution()
    for nums, answer in examples:
        print(f"{solution.summaryRanges(nums)}\n{answer}\n{"-" * 10}\n")
