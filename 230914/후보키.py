def solution(relation):
    R, C = len(relation), len(relation[0])

    answer = 0

    def is_valid(_bit):  # 서로 다른 r이 같은 튜플을 가졌는지 조사(join을 사용하는 방법도 있음)
        _bit = (C - len(_bit)) * '0' + _bit
        cnt_one = _bit.count('1')
        for r1 in range(R):
            for r2 in range(R):
                if r1 == r2:
                    continue
                cnt = 0
                for i in range(C):
                    if _bit[i] == '0':


                    if relation[r1][i] == relation[r2][i]:
                        cnt += 1
                if cnt == cnt_one:
                    return False
        return True

    fr, to = 1 << 0, (1 << C) - 1
    visited = [False] * (1 << C)
    for n in range(fr, to):
        bit = bin(n)[2:]
        if visited[n]:
            continue
        if not is_valid(bit):
            continue

        answer += 1
        visited[n] = True
        for m in range(n + 1, to):
            if m & n == n:
                visited[m] = True

    return max(answer, 1)