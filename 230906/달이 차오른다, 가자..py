import sys
from collections import deque

# 입력
N, M = map(int, sys.stdin.readline().split())
grid = [list(sys.stdin.readline().rstrip()) for _ in range(N)]


def find_start_point():
    for r in range(N):
        for c in range(M):
            if grid[r][c] == '0':
                grid[r][c] = '.'
                return r, c


sr, sc = find_start_point()
keys_bit = 0
step = 0

# BFS
dr, dc = [1, -1, 0, 0], [0, 0, 1, -1]
visited = [[[False] * (2 ** 6) for _ in range(M)] for _ in range(N)]
visited[sr][sc][0] = True
q = deque()
q.append((sr, sc, keys_bit, step))
answer = -1
while q:
    r, c, kb, st = q.popleft()
    if grid[r][c] == '1':
        answer = st
        break
    for d in range(4):
        nr = r + dr[d]
        nc = c + dc[d]
        nkb = kb
        if nr < 0 or nc < 0 or nr >= N or nc >= M:
            continue
        elif grid[nr][nc] == '#':
            continue
        elif visited[nr][nc][kb]:
            continue

        if grid[nr][nc] in ['a', 'b', 'c', 'd', 'e', 'f']:
            nkb = kb | (1 << (ord(grid[nr][nc]) - ord('a')))  # 열쇠 획득
        elif grid[nr][nc] in ['A', 'B', 'C', 'D', 'E', 'F']:
            if nkb & (1 << (ord(grid[nr][nc]) - ord('A'))) == 0:  # 해당하는 열쇠가 없는 경우
                continue
        q.append((nr, nc, nkb, st + 1))
        visited[nr][nc][nkb] = True

print(answer)
