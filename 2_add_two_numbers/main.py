from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        head = ListNode(val=None)
        current = head
        a = 0
        while l1 or l2 or a:
            x = l1.val if l1 is not None else 0
            y = l2.val if l2 is not None else 0
            result = x + y + a
            a = result // 10

            if current.val is None:
                current.val = result % 10
            else:
                current.next = ListNode(result % 10)
                current = current.next

            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        return head

def create_linked_list(arr):
    dummy = ListNode()
    current = dummy

    for num in arr:
        current.next = ListNode(num)
        current = current.next

    return dummy.next

l1 = create_linked_list([9,9,9,9,9,9,9])
l2 = create_linked_list([9,9,9,9])

solution = Solution()
result = solution.addTwoNumbers(l1, l2)

while result:
    print(result.val, end=" ")
    result = result.next