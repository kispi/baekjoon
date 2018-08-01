n = int(input())

fib = [0, 1, 1]

for i in range(100):
    if i > 2:
        fib.append(fib[i-2] + fib[i-1])

print(fib[n])