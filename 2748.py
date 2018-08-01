f = [0, 1, 1]
for i in range(100):
    if i > 2:
        f.append(f[i-2] + f[i-1])

print(f[int(input())])