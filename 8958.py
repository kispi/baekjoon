def get_score(OX):
    score = 0
    cont = 0
    for ch in OX:
        if ch == 'O':
            cont += 1
            score += cont
        else:
            cont = 0
    return score

n = int(input())

test_cases = []

for i in range(n):
    test_cases.append(input())

for case in test_cases:
    print(get_score(case))