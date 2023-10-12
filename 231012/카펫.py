from math import sqrt, ceil


def solution(brown, yellow):
    for w in range(ceil(sqrt(yellow)), yellow + 1):
        if yellow % w != 0:
            continue
        h = yellow // w

            return [w + 2, h + 2]