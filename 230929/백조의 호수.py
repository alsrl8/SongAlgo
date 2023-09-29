import sys
from collections import deque

R, C = map(int, sys.stdin.readline().split())
grid = [list(sys.stdin.readline().rstrip()) for i in range(R)]

swans = []
for r in range(R):
    for c in range(C):
        if grid[r][c] == 'L':
            grid[r][c] = '.'
            swans.append((r, c))            

dRow, dCol = [0, 0, 1, -1], [1, -1, 0, 0]

q = deque()
next_q = deque()
for r in range(R):
    for c in range(C):
        if grid[r][c] == '.':
            q.append((r, c))
            grid[r][c] = 0

date = 0
while q:
    r, c = q.popleft()
    for d in range(4):
        newR, newC = r + dRow[d], c + dCol[d]
        if newR < 0 or newC < 0 or newR >= R or newC >= C:
            continue
        elif grid[newR][newC] == '.':
            q.append((newR, newC))
            grid[newR][newC] = date
        elif grid[newR][newC] == 'X':
            next_q.append((newR, newC))
            grid[newR][newC] = date+  1

    if not q:
        q = next_q
        next_q = deque()
        date += 1

q = deque()
visited = [[-1 for c in range(C)] for r in range(R)]
q.append((swans[0][0], swans[0][1]))
visited[swans[0][0]][swans[0][1]] = 0

date = 0
next_q = deque()
while q:
    r, c = q.popleft()
    for d in range(4):
        newR, newC = r + dRow[d], c + dCol[d]
        if newR < 0 or newC < 0 or newR >= R or newC >= C:
            continue
        elif visited[newR][newC] != -1:
            continue
        elif grid[newR][newC] <= date:
            q.append((newR, newC))
            visited[newR][newC] = date
        else:
            next_q.append((newR, newC))
            visited[newR][newC] = date + 1

    if not q:
        q = next_q
        next_q = deque()
        date += 1

answer = visited[swans[1][0]][swans[1][1]]
print(answer)
