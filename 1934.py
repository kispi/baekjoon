n = int(input())

numbers = []
for i in range(n):
    numbers.append([int(x) for x in input().split(" ")])

def GCDREC(a, b, gcd):
    if a == b:
        return a
        
    for i in range(2, max(a, b)):
        if a % i == 0 and b % i == 0:
            return GCDREC(int(a / i), int(b / i), gcd * i)

    return gcd

def GCD(a, b):
    return GCDREC(a, b, 1)

def LCM(a, b):
    gcd = GCD(a, b)
    return gcd * int((a / gcd) * (b / gcd))

for number in numbers:
    print(LCM(number[0], number[1]))