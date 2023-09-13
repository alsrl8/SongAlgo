import sys

n, m = map(int, sys.stdin.readline().split())
parents = [num for num in range(n + 1)]


def union(a, b):
    pa, pb = find(a), find(b)
    if pa == pb:
        return
    else:
        if pa < pb:
            parents[pa] = pb
        else:
            parents[pb] = pa


def find(num):
    if parents[num] == num:
        return num
    else:
        parents[num] = find(parents[num])
        return parents[num]


for _ in range(m):
    cmd, a, b = map(int, sys.stdin.readline().split())
    if cmd == 0:
        union(a, b)
    elif cmd == 1:
        pa, pb = find(a), find(b)
        if pa == pb:
            print('YES')
        else:
            print('NO')
