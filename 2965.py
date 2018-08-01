l = [int(x) for x in input().split(" ")]

def swap(l, count):
    if l[2] - l[0] == 2:
        return count
    
    if l[2] - l[1] >= l[1] - l[0]:
        l[0], l[1] = l[1], l[1] + 1
        return swap(l, count + 1)
    else:
        l[1], l[2] = l[1] - 1, l[1]
        return swap(l, count + 1)
        

print(swap(l, 0))