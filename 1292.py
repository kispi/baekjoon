line = [int(x) for x in input().split(" ")]
a = line[0]
b = line[1]

list_of_numbers = []
for i in range(50):
    for j in range(i+1):
        list_of_numbers.append(i+1)

print(sum(list_of_numbers[a-1:b]))