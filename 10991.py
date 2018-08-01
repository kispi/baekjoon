n = int(input())

def star(num, lim):
    for i in range(lim-num):
        print(" ", end='')

    for i in range(num * 2 - 1):
        if i % 2 == 0:
            print("*", end='')
        else:
            print(" ", end='')
    print()

for i in range(1, n+1):
    star(i, n)