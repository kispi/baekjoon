n = int(input())

test_cases = []

for i in range(n):
    test_cases.append(input().split(" "))

for c in test_cases:
    avg = sum([int(x) for x in c[1:]]) / len(c[1:])
    
    count = 0
    for p in c[1:]:
        if int(p) > avg:
            count += 1
    val = count / len(c[1:]) * 100
    print('%.3f' % val + '%')