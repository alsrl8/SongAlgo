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