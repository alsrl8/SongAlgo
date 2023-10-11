import sys
from collections import deque

R, C = map(int, sys.stdin.readline().split())
grid = [sys.stdin.readline().rstrip() for _ in range(R)]
visited = [[False] * C for _ in range(R)]

answer = [0, 0]
dr, dc = [0, 0, 1, -1], [1, -1, 0, 0]
for r in range(R):
    for c in range(C):
        if visited[r][c]:
            continue
        elif grid[r][c] not in ['o', 'v']:
            continue

        q = deque()
        q.append((r, c))
        visited[r][c] = True
        wolf_num, sheep_num = 0, 0
        while q:
            _r, _c = q.popleft()
            if grid[_r][_c] == 'v':
                wolf_num += 1
            elif grid[_r][_c] == 'o':
                sheep_num += 1
            for d in range(4):
                nr, nc = _r + dr[d], _c + dc[d]
                if nr < 0 or nc < 0 or nr >= R or nc >= C:
                    continue
                elif grid[nr][nc] == '#':
                    continue
                elif visited[nr][nc]:
                    continue
                q.append((nr, nc))
                visited[nr][nc] = True

        if sheep_num > wolf_num:
            answer[0] += sheep_num
        else:
            answer[1] += wolf_num

print(answer[0], answer[1])
