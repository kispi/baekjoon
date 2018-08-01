n = int(input())

def star(n):
    for i in range(n):
        print("*", end='')
    print()

for i in range(1, n+1):
    star(i)
for i in reversed(range(1, n)):
    star(i)