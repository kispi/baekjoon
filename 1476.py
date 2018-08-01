l = [int(x) for x in input().split(" ")]
e = l[0]
s = l[1]
m = l[2]

def ESM(e, s, m, year):
    while True:
        if e == 1 and s == 1 and m == 1:
            return year
        e -= 1
        s -= 1
        m -= 1
        year -= 1
        if e == 0:
            e = 15
        if s == 0:
            s = 28
        if m == 0:
            m = 19

print(-ESM(e, s, m, 0) + 1)