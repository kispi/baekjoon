n = int(input())
results = []
for i in range(n):
    results.append(sum([int(x) for x in input().split(" ")]))

for (i, summation) in enumerate(results):
    print("Case #", str(i+1), ": ", summation, sep='')