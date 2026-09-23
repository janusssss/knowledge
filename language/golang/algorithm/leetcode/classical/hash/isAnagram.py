from collections import Counter


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return Counter(s) == Counter(t)
     

if __name__ == "__main__":
    ss = [("anagram", "nagaram", True)]
    solution = Solution()
    for s, t, b in ss:
        print(f"{solution.isAnagram(s, t)}\n{b}")
