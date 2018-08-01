X = [int(x) for x in str(input()).split(' ')][1]
for x in [p for p in [int(x) for x in str(input()).split(' ')] if p < X]:
    print(str(x) + ' ', end='')

