chess = []

for i in range(8):
    chess.append(input())

count = 0
for i in range(8):
    for j in range(8):
        if chess[i][j] == 'F' and (i + j) % 2 == 0:
            count += 1

print(count)