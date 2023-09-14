import sys
from collections import deque

# 입력
N, M = map(int, sys.stdin.readline().split())
R, C = M, N
grid = [list(map(int, sys.stdin.readline().split())) for _ in range(R)]

# 1) 벽을 제거하지 않은 상태에서 방의 개수와 넓이 구하기
dr, dc = [0, -1, 0, 1], [-1, 0, 1, 0]
visited = [[False] * C for _ in range(R)]
room_num = 0
max_room_size = 0
for r in range(R):
    for c in range(C):
        if visited[r][c]:
            continue

        room_num += 1
        room_size = 1

        q = deque()
        q.append((r, c))
        visited[r][c] = True
        while q:
            _r, _c = q.popleft()
            for d in range(4):
                nr, nc = _r + dr[d], _c + dc[d]
                if nr < 0 or nc < 0 or nr >= R or nc >= C:
                    continue
                elif visited[nr][nc]:
                    continue
                elif (1 << d) & grid[_r][_c] > 0:
                    continue
                q.append((nr, nc))
                visited[nr][nc] = True
                room_size += 1
        max_room_size = max(max_room_size, room_size)

print(room_num)
print(max_room_size)

# 벽을 하나 제거했을 때 방의 최대 넓이
max_room_size = 0
for r in range(R):
    for c in range(C):
        for w in range(4):
            if grid[r][c] & (1 << w) == 0:
                continue
            grid[r][c] = grid[r][c] ^ (1 << w)

            visited = [[False] * C for _ in range(R)]
            for _r in range(R):
                for _c in range(C):
                    if visited[_r][_c]:
                        continue

                    room_size = 1

                    q = deque()
                    q.append((_r, _c))
                    visited[_r][_c] = True
                    while q:
                        __r, __c = q.popleft()
                        for d in range(4):
                            nr, nc = __r + dr[d], __c + dc[d]
                            if nr < 0 or nc < 0 or nr >= R or nc >= C:
                                continue
                            elif visited[nr][nc]:
                                continue
                            elif (1 << d) & grid[__r][__c] > 0:
                                continue
                            q.append((nr, nc))
                            visited[nr][nc] = True
                            room_size += 1
                    max_room_size = max(max_room_size, room_size)

            grid[r][c] = grid[r][c] | (1 << w)
print(max_room_size)
