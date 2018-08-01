class Num():
    def __init__(self, val):
        self.value = val
        self.count = 0

numbers = [Num(x) for x in range(10)]

A = int(input())
B = int(input())
C = int(input())
for ch in str(A*B*C):
    numbers[int(ch)].count += 1

for number in numbers:
    print(number.count)