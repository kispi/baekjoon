import sys

def mat_sum(mat, pos):
    result = []
    for i in range(pos[0], pos[2]+1):
        result.append(sum(mat[i][pos[1]:pos[3]+1]))
    print(sum(result))

l = [int(x) for x in sys.stdin.readline().rstrip().split(" ")]
n = l[0]
m = l[1]
matrix = []
for i in range(n):
    matrix.append([int(x) for x in sys.stdin.readline().rstrip().split(" ")])

k = int(sys.stdin.readline().rstrip())
pos = []
for i in range(k):
    pos.append([int(x)-1 for x in sys.stdin.readline().rstrip().split(" ")])

for i in range(k):
    mat_sum(matrix, pos[i])