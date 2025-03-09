class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        word = ''
        while word1 or word2:
            if word1:
                word += word1[:1]
                word1 = word1[1:]
            if word2:
                word += word2[:1]
                word2 = word2[1:]
        return word

solution = Solution()
print(solution.mergeAlternately(word1="ab", word2="pqrs"))