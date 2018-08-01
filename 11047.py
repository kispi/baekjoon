l = [int(x) for x in input().split(" ")]
n = l[0]
k = l[1]

coins = []
for i in range(n):
    coins.append(int(input()))

count = 0
for coin in reversed(coins):
    while k >= coin:
        num_of_coin = int(k / coin)
        k -= num_of_coin * coin
        count += num_of_coin

print(count)