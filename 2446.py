n = int(input())

def star(n, total):
    for i in range(n):
        print(" ", end='')
    for i in range(total*2-1-2*n):
        print("*", end='')
    print()

for i in range(n):
    star(i, n)
for i in reversed(range(n-1)):
    star(i, n)