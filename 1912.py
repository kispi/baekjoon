def sumsOfNNumbers(numbers, n):
    result = []
    for i in range(len(numbers) - n + 1):
        result.append(sum(numbers[i:i+n]))
    return result

input()
numbers = [int(x) for x in input().split(" ")]
candidates = []

for i in range(len(numbers)):
    candidates.append(max(sumsOfNNumbers(numbers, i+1)))

print(max(candidates))