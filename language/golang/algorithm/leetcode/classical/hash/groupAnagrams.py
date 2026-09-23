from collections import defaultdict
from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # 使用 defaultdict 来自动处理不存在的键
        anagrams = defaultdict(list)
        for s in strs:
            sorted_str = ''.join(sorted(s))
            anagrams[sorted_str].append(s)
        return list(anagrams.values())


if __name__ == "__main__":
    s = [(["eat", "tea", "tan", "ate", "nat", "bat"], [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]])]

    for strs, answer in s:
        solution = Solution()
        print(f"{solution.groupAnagrams(strs)}\n{answer}")
