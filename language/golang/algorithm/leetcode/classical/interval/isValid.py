class Solution:
    def isValid(self, s: str) -> bool:
        if not len(s) % 2 == 0:
            return False
        symbol_dict = {')': '(', '}': '{', ']': '['}
        symbols = {}
        for symbol in s:
            if symbol in ['(', '{', '[']:
                symbols[len(symbols) + 1] = symbol
            else:
                if len(symbols) == 0:
                    return False
                if symbols.pop(len(symbols)) != symbol_dict[symbol]:
                    return False
        if len(symbols) > 0:
            return False
        return True


if __name__ == "__main__":
    solution = Solution()
    print(solution.isValid("()"))
