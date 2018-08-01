n = int(input())

def star(n, total):
    for i in range(n):
        print("*", end='')
    for i in  range((total-n-1)*2):
        print(" ", end='')
    for i in range(n):
        print("*", end='')
    print()

for i in range(n):
    star(i+1, n+1)
for i in reversed(range(n-1)):
    star(i+1, n+1)