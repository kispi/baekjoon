import sys

n = int(input())
cards = [int(x) for x in sys.stdin.readline().rstrip().split(" ")]
cards.sort()

m = int(input())
cards_to_determine = [int(x) for x in sys.stdin.readline().rstrip().split(" ")]

def bin_find(target, lst):
    if len(lst) == 1:
        if target == lst[0]:
            return True
        else:
            return False

    half = int(len(lst) / 2)
    if lst[half] == target:
        return True
    elif lst[half] > target:
        return bin_find(target, lst[:half])
    else:
        return bin_find(target, lst[half:])

for card in cards_to_determine:
    if bin_find(card, cards) is True:
        print("1", end=' ')
    else:
        print("0", end=' ')