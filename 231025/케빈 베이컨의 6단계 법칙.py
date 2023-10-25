import sys
from collections import deque

N, M = map(int, sys.stdin.readline().split())
conn = [[False] * N for _ in range(N)]
for _ in range(M):
    A, B = map(int, sys.stdin.readline().split())
    conn[A - 1][B - 1] = True
    conn[B - 1][A - 1] = True

answer = -1
min_kevin_num = (N + 1) ** 2

for i in range(N):
    q = deque()
    q.append((i, 0))
    visited = [False] * N
    visited[i] = True
    kevin_num = 0
    while q:
        person, num = q.popleft()
        kevin_num += num
        for nxt in range(N):
            if not conn[person][nxt]:
                continue
            elif visited[nxt]:
                continue
            visited[nxt] = True
            q.append((nxt, num + 1))
    if kevin_num < min_kevin_num:
        min_kevin_num = kevin_num
        answer = i

print(answer + 1)
