import sys

n = int(sys.stdin.readline().rstrip())
mtt = []
for i in range(n):
    mtt.append(int(sys.stdin.readline().rstrip()))

print(sum(mtt) - len(mtt) + 1)