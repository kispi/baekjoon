import math

n = input()

trial = []
answers = [None] * 1000

def do(n, depth, num):
    global trial
    if n == 1:
        trial.append(depth)
        return
    elif answers[num] is not None:
        return answers[num]
    else:
        depth += 1
        if n % 3 == 0:
            do(n/3, depth, num)

        if n % 2 == 0:
            do(n/2, depth, num)

        do(n-1, depth, num)

def wrapper(n):
    if answers[n] is not None:
        return answers[n]
    else:
        do(n, 0, n)
        return sorted(trial)[0]

print(wrapper(int(n)))