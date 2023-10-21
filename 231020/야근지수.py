import heapq


def solution(n, works):
    s = sum(works) - n
    if s < 0:  # 야근을 하지 않는 경우
        return 0

    pq = []  # max heap
    for w in works:
        heapq.heappush(pq, -w)

    for _ in range(n):
        w = -heapq.heappop(pq)
        heapq.heappush(pq, -(w - 1))

