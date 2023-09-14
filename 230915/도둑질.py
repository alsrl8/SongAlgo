def solution(money):
    answer = 0
    house = len(money)

    # 0번 집을 터는 경우
    dp = [[0, 0] for _ in range(house)]  # [해당 집을 털었을 경우, 털지 않았을 경우]
    dp[0][0] = money[0]


        dp[i][1] = max(dp[i-1])

    answer = max(answer, max(dp[-2]))


    # 0번 집을 털지 않는 경우
    dp[0][0] = 0
    for i in range(1, house):
        dp[i][0] = dp[i-1][1] + money[i]
        dp[i][1] = max(dp[i-1])

    answer = max(answer, max(dp[-1]))

    return answer