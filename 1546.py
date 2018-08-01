input()
numbers = [int(x) for x in input().split(" ")]
numbers.sort()
highest = numbers[len(numbers) - 1]
print(sum([x / highest * 100 for x in numbers]) / len(numbers))