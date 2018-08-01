def rec(plates, score):
    if len(plates) == 1:
        return score
    score = score + 10 if plates[0] != plates[1] else score + 5
    return rec(plates[1:], score)

print(rec(input(), 10))