def repeat(n, word):
    for ch in word:
        print(ch * int(n), end='')
    print()
    
n = int(input())
test_cases = []
for i in range(n):
    test_cases.append(input().split(" "))

for t in test_cases:
    repeat(t[0], t[1])