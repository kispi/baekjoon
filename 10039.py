numbers = []
for i in range(5):
    num = int(input())
    numbers.append(40 if num < 40 else num)

print(int(sum(numbers)/5))