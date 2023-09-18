import sys

sys.setrecursionlimit(100001)

N = int(sys.stdin.readline())
parents = [0 for _ in range(N + 1)]
conn = [[] for _ in range(N + 1)]
visited = [False for _ in range(N + 1)]
for _ in range(N-1):
    a, b = map(int, sys.stdin.readline().split())
    conn[a].append(b)
    conn[b].append(a)


def dfs(node):
    for nxt in conn[node]:
        if visited[nxt]:
            continue
        visited[nxt] = True
        parents[nxt] = node
        dfs(nxt)


dfs(1)
for i in range(2, N+1):
    print(parents[i])
