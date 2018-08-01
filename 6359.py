t = int(input())

cases = []
for i in range(t):
    cases.append(int(input()))

def room(n):
    rooms = [int(x) for x in list('1' * 100)]
    if n == 1:
        return rooms

    for a in range(n):
        for i, room in enumerate(rooms):
            if (i+1) % (a+1) == 0:
                if rooms[i] == 1:
                    rooms[i] = 0
                else:
                    rooms[i] = 1

    zeros = 0
    for room in rooms[:n]:
        if room == 0:
            zeros += 1
    
    return zeros

for case in cases:
    print(room(case))