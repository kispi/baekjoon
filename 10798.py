lines = []
for i in range(5):
    lines.append(input())

for i in range(15):
    for line in lines:
        try: 
            if line[i] is not None:
                print(line[i], end='')
        except:
            pass