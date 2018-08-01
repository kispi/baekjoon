results = []
while True:
    aPb = sum([int(x) for x in input().split(" ")])
    if aPb == 0:
        break
    results.append(aPb)

for summation in results:
    print(summation)