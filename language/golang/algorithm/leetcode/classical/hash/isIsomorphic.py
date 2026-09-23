class ExampleIsIsomorphic:
    def __init__(self, s: str, t: str, answer: bool):
        self.s = s
        self.t = t
        self.answer = answer

def is_isomorphic(s: str, t: str) -> bool:
    mapping_s_to_t = {}
    mapping_t_to_s = {}

    for char_s, char_t in zip(s, t):
        if (char_s in mapping_s_to_t and mapping_s_to_t[char_s] != char_t) or \
                (char_t in mapping_t_to_s and mapping_t_to_s[char_t] != char_s):
            return False
        mapping_s_to_t[char_s] = char_t
        mapping_t_to_s[char_t] = char_s

    return True

if __name__ == "__main__":
    examples = [
        ExampleIsIsomorphic("egg", "add", True),
        ExampleIsIsomorphic("badc", "baba", False),
    ]

    for example in examples:
        result = is_isomorphic(example.s, example.t)
        print(f"{result}\n{example.answer}\n---------------")