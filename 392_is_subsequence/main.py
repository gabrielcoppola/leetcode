
class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        sequence = ''
        start_t, start_s = 0, 0
        for idx_s in range(start_s, len(s)):
            for idx_t in range(start_t, len(t)):
                if t[idx_t] == s[idx_s]:
                    start_s += 1
                    start_t = idx_t
                    sequence += s[idx_s]
                    break
            start_t += 1
        return s == sequence
        
solution = Solution()
print(solution.isSubsequence(s = "aaaaaa", t = "bbaaaa"))