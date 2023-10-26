import sys

N = int(sys.stdin.readline())
M = int(sys.stdin.readline())

can_go = [[0] * (N + 1) for _ in range(N + 1)]
# can_go[i][j]: i에서 j까지 큰 순서임이 보장되는지
for _ in range(M):
    x, y = map(int, sys.stdin.readline().split())
    can_go[x][y] = 1

for k in range(1, N + 1):
    for i in range(1, N + 1):
        for j in range(1, N + 1):
            can_go[i][j] = can_go[i][j] | (can_go[i][k] & can_go[k][j])

for i in range(1, N + 1):
    known_cnt = 1  # 자기 자신
    for j in range(1, N + 1):
        if (can_go[i][j] | can_go[j][i]) > 0:
            known_cnt += 1
    print(N - known_cnt)
