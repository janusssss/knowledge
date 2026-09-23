from typing import List


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        d = dict.fromkeys(nums, True)
        longest = 0
        for number in d.keys():
            if number - 1 in d:
                continue
            long = 0
            while number in d.keys():
                number += 1
                long += 1
            else:
                longest = long if long > longest else longest
        return longest


if __name__ == "__main__":
    examples = [
        ([9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6], 7),
        ([100, 4, 200, 1, 3, 2], 4),
    ]
    solution = Solution()
    for nums, answer in examples:
        print(f"{solution.longestConsecutive(nums)}\n{answer}\n")
