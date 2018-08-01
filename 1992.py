n = int(input())
lines = []
for i in range(n):
    lines.append(list(input()))

def comp(square):
    if len(square) == 1:
        if square[0][0] == '1':
            return '1'
        return '0'

    dots = len(square) * len(square)
    ones = 0
    zeros = 0
    for line in square:
        for ch in line:
            if ch == '1':
                ones += 1
            else:
                zeros += 1    
    if dots == ones:
        return '1'
    elif dots == zeros:
        return '0'
    else:
        half = int(len(square) / 2)
        TL = []
        TR = []
        BL = []
        BR = []
        for line in square[:half]:
            TL.append(line[:half])
        for line in square[:half]:
            TR.append(line[half:])
        for line in square[half:]:
            BL.append(line[:half])
        for line in square[half:]:
            BR.append(line[half:])
        return "(" + comp(TL) + comp(TR) + comp(BL) + comp(BR) + ")"

print(comp(lines))