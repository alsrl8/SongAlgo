import sys

N, M = map(int, sys.stdin.readline().split())
costs = [int(sys.stdin.readline()) for _ in range(N)]


def is_enough(money: int) -> bool:
    '''
    주어진 money만큼 M번 인출해서 costs를 감당할 수 있는지 조사하여 bool type을 반환한다.
    :param money: money는 반드시 costs의 최댓값 이상이어야 한다.
    :return:
    '''
    global M, costs

    cur = money
    withdrawal_num = 1  # 인출 횟수
    for cost in costs:
        if cur >= cost:
            cur -= cost
        else:
            if withdrawal_num == M:  # 이미 인출 횟수가 M번이 됐다면 더 이상 인출할 수 없음
                return False
            withdrawal_num += 1
            cur = money - cost
    return True


# 이분 탐색
answer = float('inf')
MAX_COST = max(costs)
L, R = MAX_COST, MAX_COST * N
while L <= R:
    money = (L + R) // 2
    if is_enough(money):
        R = money - 1
        answer = min(answer, money)
    else:
        L = money + 1

print(answer)
