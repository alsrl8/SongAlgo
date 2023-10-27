import sys
from collections import deque

R, C = map(int, sys.stdin.readline().split())
grid = [list(map(int, sys.stdin.readline().split())) for _ in range(R)]


def find_outside_cheese_list():
    cheese_list = []
    dr, dc = [0, 0, 1, -1], [1, -1, 0, 0]
    visited = [[False] * C for _ in range(R)]

    air = deque()
    for r in range(R):
        air.append((r, 0))
        air.append((r, C - 1))
        visited[r][0] = True
        visited[r][C - 1] = True
    for c in range(C):
        air.append((0, c))
        air.append((R - 1, c))
        visited[0][c] = True
        visited[R - 1][c] = True

    while air:
        r, c = air.popleft()
        for d in range(4):
            nr, nc = r + dr[d], c + dc[d]
            if nr < 0 or nr >= R or nc < 0 or nc >= C:
                continue
            elif visited[nr][nc]:
                continue

            visited[nr][nc] = True
            if grid[nr][nc] == 1:
                cheese_list.append((nr, nc))
            else:
                air.append((nr, nc))

    return cheese_list


cheese_num = 0
for r in range(R):
    for c in range(C):
        if grid[r][c] == 1:
            cheese_num += 1

if cheese_num == 0:
    print(0)
    print(0)
    exit(0)

cheese_list = find_outside_cheese_list()

hour = 0
while True:
    hour += 1
    cheese_list = find_outside_cheese_list()
    if len(cheese_list) == cheese_num:
        print(hour)
        print(cheese_num)
        break

    cheese_num -= len(cheese_list)
    for r, c in cheese_list:
        grid[r][c] = 0
