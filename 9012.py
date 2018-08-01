def isVPS(par):
    shouldBeClosed = 0
    for ch in par:
        if ch == '(':
            shouldBeClosed += 1
        else:
            if shouldBeClosed == 0:
                return False
            shouldBeClosed -= 1
    
    if shouldBeClosed == 0:
        return True
    else:
        return False

n = int(input())
cases = []
for i in range(n):
    cases.append(input())

for case in cases:
    print("YES") if isVPS(case) else print("NO")