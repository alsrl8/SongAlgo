from collections import defaultdict

answer = 1


def solution(info, edges):
    n = len(info)
    children = defaultdict(list)
    for p, c in edges:


    visited = [False] * (1 << 17)
    visited[0] = True

    def search(cur, cnt_wolf, cnt_sheep):
        global answer
        answer = max(answer, cnt_sheep)
        for i in range(n):
            if cur & (1 << i) == 0:
                continue
            for child in children[i]:
                if cur & (1 << child):
                    continue
                elif cnt_wolf + 1 >= cnt_sheep and info[child] == 1:  # 늑대를 방문한 경우
                    continue
                if info[child] == 1:
                    search(cur | (1 << child), cnt_wolf + 1, cnt_sheep)
                else:
                    search(cur | (1 << child), cnt_wolf, cnt_sheep + 1)

    search(1, 0, 1)
    return answer
