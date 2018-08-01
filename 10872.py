n = int(input())
if n is 0:
    print(1)
else:        
    for i in reversed(range(1, n)):
        n = n * i
    print(n)