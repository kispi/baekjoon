input()
numbers = [int(x) for x in input().split(" ")]
numbers.sort()
zigzag = []
for i in range(int(len(numbers) / 2)):
    zigzag.append(numbers[len(numbers)-1])
    numbers = numbers[:-1]
    zigzag.append(numbers[0])
    numbers = numbers[1:]

print(zigzag)