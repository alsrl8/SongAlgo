import sys
from collections import deque

M, N = map(int, sys.stdin.readline().split())
grid = [list(sys.stdin.readline().rstrip()) for _ in range(N)]

visited = [[False] * M for _ in range(N)]
dr, dc = [0, 0, 1, -1], [1, -1, 0, 0]

power = [0, 0]

for r in range(N):
    for c in range(M):
        if visited[r][c]:
            continue

        cnt = 0
        visited[r][c] = True
        q = deque()
        q.append((r, c))
        while q:
            _r, _c = q.popleft()
            cnt += 1
            for d in range(4):
                nr, nc = _r + dr[d], _c + dc[d]
                if nr < 0 or nc < 0 or nr >= N or nc >= M:
                    continue
                elif visited[nr][nc]:
                    continue
                elif grid[nr][nc] != grid[r][c]:
                    continue
                q.append((nr, nc))
                visited[nr][nc] = True

        if grid[r][c] == 'W':
            power[0] += cnt ** 2
        else:
            power[1] += cnt ** 2

print(power[0], power[1])
