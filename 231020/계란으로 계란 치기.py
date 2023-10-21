import sys

N = int(sys.stdin.readline())
eggs = [list(map(int, sys.stdin.readline().split())) for _ in range(N)]

# eggs[i][0]: i번째 계란의 내구도
# eggs[i][1]: i번째 계란의 무게


answer = 0


def dfs(i):
    global answer
    if i == N:  # 모든 계란을 들어본 경우(dfs가 종료된 경우)
        cnt = len([j for j in range(N) if eggs[j][0] <= 0])
        answer = max(answer, cnt)
        return
    elif eggs[i][0] <= 0:  # 내가 들어올린 계란의 내구도가 0이면 바로 내려놓음
        dfs(i + 1)
        return

    # 내가 지금 들어올린 계란 이외에 모든 계란이 깨져있는 경우 -> 탐색 종료
    candidates = [j for j in range(N) if i != j and eggs[j][0] > 0]
    if len(candidates) == 0:
        dfs(N)
        return