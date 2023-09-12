import sys

N, M = map(int, sys.stdin.readline().split())
front = [[] for _ in range(N + 1)]  # 나보다 작은 학생의 번호
back = [[] for _ in range(N + 1)]
for _ in range(M):
    a, b = map(int, sys.stdin.readline().split())
    front[b].append(a)
    back[a].append(b)


def search(n: int, arr: list, set):
    for e in arr[n]:
        if e in set:
            continue
        set.add(e)
        search(e, arr, set)


answer = 0
for n in range(1, N + 1):
    set_f, set_b = set(), set()
    search(n, front, set_f)
    search(n, back, set_b)

    if len(set_f) + len(set_b) == N - 1:
        answer += 1

print(answer)
