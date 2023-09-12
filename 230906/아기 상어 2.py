import sys
from collections import deque

# 입력
N, M = map(int, sys.stdin.readline().split())
grid = [list(map(int, sys.stdin.readline().split())) for _ in range(N)]

# 안전 거리 최댓값
MAX = N * M
# 안전 거리를 저장하는 list, safe_dist_grip[r][c]는 (r,c) 좌표의 안전 거리를 나타낸다.
safe_dist_grip = [[MAX] * M for _ in range(N)]

# 상하좌우, 대각선 방향의 증감을 관리
dr, dc = [0, -1, -1, -1, 0, 1, 1, 1], [1, 1, 0, -1, -1, -1, 0, 1]

# 모든 좌표를 순회하면서 아기 상어를 탐색
for r in range(N):
    for c in range(M):
        if grid[r][c] == 0:
            continue
        # 아기 상어의 상하 좌우로 BFS를 수행하면서 안전 거리를 최솟값으로 갱신
        safe_dist = 0
        q = deque()
        q.append((r, c, safe_dist))
        safe_dist_grip[r][c] = 0
        while q:
            _r, _c, ds = q.popleft()
            for d in range(8):
                nr, nc = _r + dr[d], _c + dc[d]
                if nr < 0 or nc < 0 or nr >= N or nc >= M:
                    continue
                elif safe_dist_grip[nr][nc] <= ds + 1:
                    continue
                safe_dist_grip[nr][nc] = ds + 1
                q.append((nr, nc, ds + 1))

answer = max([max(row) for row in safe_dist_grip])
print(answer)
