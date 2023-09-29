import sys
from typing import List

# 입력
n = int(sys.stdin.readline())
m = int(sys.stdin.readline())
MAX = 10000001  # 비용의 최댓값: 모든 도시를 거쳐 최대 비용으로 끝에서 끝까지 이동하는 비용의 합보다 커야 한다.
conn = [[MAX] * (n + 1) for _ in range(n + 1)]
for i in range(1, n + 1):
    conn[i][i] = 0
for _ in range(m):
    a, b, c = map(int, sys.stdin.readline().split())
    conn[a][b] = min(conn[a][b], c)

for k in range(1, n + 1):
    for i in range(1, n + 1):
        for j in range(1, n + 1):
            conn[i][j] = min(conn[i][j], conn[i][k] + conn[k][j])

for i in range(1, n+1):
    for j in range(1, n+1):
        print(conn[i][j] if conn[i][j] != MAX else 0, end=' ')
    print()
