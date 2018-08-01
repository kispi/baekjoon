l = [int(x) for x in input().split(" ")]
if l[0] == 1 and l[1] == 2 and l[2] == 3 and l[3] == 4 and l[4] == 5 and l[5] == 6 and l[6] == 7 and l[7] == 8:
    print("ascending")
elif l[0] == 8 and l[1] == 7 and l[2] == 6 and l[3] == 5 and l[4] == 4 and l[5] == 3 and l[6] == 2 and l[7] == 1:
    print("descending")
else:
    print("mixed")
