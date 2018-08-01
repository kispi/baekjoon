import math

n = str(math.factorial(int(input())))

count = 0
for ch in reversed(list(n)):
    if ch == '0':
        count += 1
    else:
        break

print(count)