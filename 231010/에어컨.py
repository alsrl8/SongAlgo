def solution(temperature, t1, t2, a, b, onboard):
    MAX = float('inf')
    temperature += 10
    t1, t2 = t1 + 10, t2 + 10  # 정신 건강을 위해 0~50도로 온도 조절
    dp = [MAX] * 51
    dp[temperature] = 0


        nxt_dp = [MAX] * 51
        # 탐색 범위 지정(승객이 탑승한다면 t1~t2까지 온도만 탐색)
        l = 0 if onboard[second] == 0 else t1
        r = 50 if onboard[second] == 0 else t2
        for t in range(l, r + 1):
            # 에어컨을 끄는 경우
            if t == temperature:
                candidates = [dp[t]]
                if t + 1 < 51:
                    candidates.append(dp[t + 1])
                if t - 1 >= 0:
                    candidates.append(dp[t-1])
                nxt_dp[t] = min(candidates)
            elif t < temperature and t - 1 >= 0 and dp[t - 1] < MAX:
                nxt_dp[t] = dp[t - 1]
            elif t > temperature and t + 1 < 51 and dp[t + 1] < MAX:
                nxt_dp[t] = dp[t + 1]

            # 에어컨을 켜는 경우
            candidates = [nxt_dp[t], dp[t] + b]
            if t - 1 >= 0:
                candidates.append(dp[t - 1] + a)
            if t + 1 < 51:
                candidates.append(dp[t + 1] + a)
            nxt_dp[t] = min(candidates)
        dp = nxt_dp

    return min(dp)