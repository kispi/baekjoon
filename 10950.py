n = int(input())
pairs = []

for i in range(n):
    pairs.append([int(x) for x in input().split(" ")])

for p in pairs:
    print(sum(p))