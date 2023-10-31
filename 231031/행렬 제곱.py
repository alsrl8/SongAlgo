import sys
from math import log2

N, B = map(int, sys.stdin.readline().split())
A = [list(map(int, sys.stdin.readline().split())) for _ in range(N)]

dic = dict()
dic[1] = A


def multiple(m1, m2):
    result = [[0] * N for _ in range(N)]
    for r in range(N):
        for c in range(N):
            x = 0
            for i in range(N):
                x += m1[r][i] * m2[i][c]
                x %= 1000
            result[r][c] = x
    return result


def divide(num):
    if num in dic:
        return dic[num]

    a = 1 << int(log2(num))
    b = num - a
    if b == 0:
        a >>= 1
        b = a
    ret = multiple(divide(a), divide(b))
    dic[num] = ret
    return dic[num]


divide(B)
for r in range(N):
    for c in range(N):
        print(dic[B][r][c] % 1000, end=' ')
    print('')
