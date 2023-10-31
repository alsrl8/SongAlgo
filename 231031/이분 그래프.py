import sys

sys.setrecursionlimit(200000)

K = int(sys.stdin.readline())
for _ in range(K):
    V, E = map(int, sys.stdin.readline().split())
    conn = [[] for _ in range(V + 1)]
    for _ in range(E):
        u, v = map(int, sys.stdin.readline().split())
        conn[u].append(v)
        conn[v].append(u)

    color = [0] * (V + 1)


    def dfs(node):  # node와 연결된 다른 node들의 색이 node와 달라야 한다.
        result = True
        for nxt_node in conn[node]:
            if color[nxt_node] == color[node]:
                return False
            elif color[nxt_node] != 0:
                continue

            color[nxt_node] = color[node] * -1
            result &= dfs(nxt_node)

        return result


    result = True
    for node in range(1, V + 1):
        if color[node] != 0:
            continue
        color[node] = 1
        result &= dfs(node)

    print('YES' if result else 'NO')
