class Solution:
    # def romanToInt(self, s: str) -> int:
    #     result = 0
    #     minus = 0
    #     teste = []
    #     romans = {
    #         "I": 1,
    #         "V": 5,
    #         "X": 10,
    #         "L": 50,
    #         "C": 100,
    #         "D": 500,
    #         "M": 1000,
    #         "IV": 4,
    #         "IX": 9,
    #         "XL": 40,
    #         "XC": 90,
    #         "CD": 400,
    #         "CM": 900       
    #     }
    #     for character in range(len(s)):
    #         for roman in romans:
    #             if roman == s[character]:
    #                 teste.append(romans.get(roman))
    #     for num in range(len(teste)-1):
    #         if teste[num + 1] > teste[num]:
    #             print(teste[num]) 
    #             minus += teste[num]
    #         else:
    #             result += teste[num]
    #     return (result - minus) + teste[-1]

    def romanToInt(self, s: str) -> int:
        numbers = []
        romans = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000
        }
        for index in range(len(s)):
            for roman in romans:
                if s[index] == roman:
                    numbers.append(romans.get(roman))
        print(numbers)






solution = Solution()
print(solution.romanToInt(s="LVIII"))