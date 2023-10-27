import sys

N, M = map(int, sys.stdin.readline().split())
order = list(map(int, sys.stdin.readline().split()))

cur = 0
visited = [False] * N


def go_right(pos):
    visited[pos] = True
    while True:
        pos = (pos + 1) % N
        if not visited[pos]:
            break
    return pos


def count_left(pos, tgt):
    cnt = 0
    while pos != tgt:
        cnt += 1
        while True:
            pos = (pos - 1) % N
            if not visited[pos]:
                break
    return cnt


def count_right(pos, tgt):
    cnt = 0
    while pos != tgt:
        cnt += 1
        while True:
            pos = (pos + 1) % N
            if not visited[pos]:
                break
    return cnt


answer = 0
for i, tgt in enumerate(order):
    tgt -= 1
    if cur == tgt:
        if i == M - 1:
            break
        cur = go_right(cur)
        continue

    cl = count_left(cur, tgt)
    cr = count_right(cur, tgt)
    answer += min(cl, cr)
    cur = go_right(tgt)

print(answer)
