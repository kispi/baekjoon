results = []
for i in range(5):
    results.append(sum([int(x) for x in input().split(" ")]))
print(results.index(max(results)) + 1, max(results))