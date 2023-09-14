from collections import deque

def is_safe(place, r, c):
    q = deque()
    q.append((r, c, 0))
    visited = [[False] * 5 for _ in range(5)]
    visited[r][c] = True
    


    while q:
        _r, _c, step = q.popleft()
        for d in range(4):
            nr, nc = _r + dr[d], _c + dc[d]
            if nr < 0 or nc < 0 or nr >= 5 or nc >= 5:
                continue
            elif place[nr][nc] == 'X':
                continue
            elif visited[nr][nc]:
                continue
            elif place[nr][nc] == 'P':
                return False
            visited[nr][nc] = True
            if step + 1 == 1:
                q.append((nr, nc, step + 1))
    return True

def is_safe_place(place):
    for r in range(5):
        for c in range(5):
            if place[r][c] != 'P':
                continue
            if not is_safe(place, r, c):
                return False
    return True
                    

def solution(places):
    answer = [0] * len(places)
    for i, place in enumerate(places):
        if is_safe_place(place):
            answer[i] = 1
    return answer