N = int(input())
maze = list(map(int, input().split()))

if maze[0] == 0 and N > 1:
    print(-1)
    exit(0)

dp = [float('inf') for _ in range(N)]
dp[0] = 0
for i, num in enumerate(maze):
    for j in range(i + 1, min(i + num + 1, N)):
        dp[j] = min(dp[j], dp[i] + 1)

print(dp[-1] if dp[-1] < float('inf') else -1)
