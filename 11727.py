n = int(input())

answers = [1, 3, 5]

for i in range(3, n):
    answers.append(answers[i-1] + 2 * answers[i-2])

print(answers[n-1] % 10007)