import math

def comb(n, r):
    return math.factorial(n) / (math.factorial(n-r) * math.factorial(r))

n = input()
inputs = []
for i in range(int(n)):
    inputs.append(input())

for line in inputs:
    splitted = str.split(line, " ")
    r = splitted[0]
    n = splitted[1]
    print(int(comb(int(n), int(r))))