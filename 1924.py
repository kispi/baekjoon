l = [int(x) for x in input().split(" ")]
x = l[0]
y = l[1]

day = ['MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT', 'SUN']
month = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

sum = 0
for i in range(x-1):
    sum += month[i]
sum += y - 1

print(day[sum%7])