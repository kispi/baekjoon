n = int(input())

stair = []
result = [[0] * 5 for i in range(0, n+1)]

for i in range(0, n):
    stair.append(int(input()))

for i in range(1, n+1):
    result[i][2] = max(result[i-1][1], result[i-1][0])
    result[i][1] = result[i-1][2] + stair[i-1]
    result[i][0] = result[i-1][1] + stair[i-1]

print(max(result[i][1], result[i][0]))

