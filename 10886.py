n = int(input())

cute = 0
for i in range(n):
    if input() == '1':
        cute += 1
connector = "not " if cute <= (n/2) else ""
print("Junhee is ", connector, "cute!", sep='')