from typing import List


class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        if not points:
            return 0
        points.sort(key=lambda p: p[1])
        right = points[0][1]
        ans = 1
        for p in points:
            if p[0] > right:
                ans += 1
                right = p[1]
        return ans
