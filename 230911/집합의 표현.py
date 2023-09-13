import sys

# Python 재귀의 최대 깊이를 지정한다.
sys.setrecursionlimit(100000)

n, m = map(int, sys.stdin.readline().split())

# 처음 부모는 자기 자신으로 정한다.
parents = [i for i in range(n + 1)]


def find_parent(num: int) -> int:
    if parents[num] != num:
        ancestor = find_parent(parents[num])
        parents[num] = ancestor
    return parents[num]


# Union-find
for _ in range(m):
    cmd, a, b = map(int, sys.stdin.readline().split())
    if cmd == 0:
        pa, pb = find_parent(a), find_parent(b)
        pa, pb = max(pa, pb), min(pa, pb)
        parents[pa] = pb
    elif cmd == 1:
        pa, pb = find_parent(a), find_parent(b)
        if pa != pb:
            print('NO')
        else:
            print('YES')
