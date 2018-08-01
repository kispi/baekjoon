n = int(input())

twochins = [0, 1, 1]

# def is2chin(binStr):
#     if binStr[0] == '0' or "11" in binStr:
#         return False
#     return True            

# def int2bin(num):
#     return "{0:b}".format(num)

# def get2chins(n):
#     count = 0
#     for i in range(pow(2, n)):
#         if i >= pow(2, n-1):
#             if is2chin(int2bin(i)):
#                 print(int2bin(i))
#                 count += 1
#     return count

# for i in range(10):
#     twochins.append(get2chins(i))

# print(twochins)

# Oh it's fucking fibonaccccccci

for i in range(100):
    if i > 2:
        twochins.append(twochins[i-2] + twochins[i-1])

print(twochins[n])