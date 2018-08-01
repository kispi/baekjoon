numbers = []
for i in range(7):
    a = int(input())
    numbers.append(a) if a % 2 == 1 else None
print(sum(numbers), '\n', min(numbers), sep='') if len(numbers) > 0 else print(-1)