T = int(input())

cases = []
answers = [1, 1, 2]

for i in range(100):
    if i >= 3:
        answers.append(answers[i-3] + answers[i-2] + answers[i-1])

for i in range(T):
    cases.append(int(input()))

for case in cases:
    print(answers[case])