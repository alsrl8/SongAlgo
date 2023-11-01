import sys
import heapq

N, M = map(int, sys.stdin.readline().split())
classes = [sorted(list(map(int, sys.stdin.readline().split()))) for _ in range(N)]
indices = [0] * N

pq = []

max_val = 0
for i, cls in enumerate(classes):
    heapq.heappush(pq, (cls[0], i))
    max_val = max(max_val, cls[0])

min_val = pq[0][0]
answer = max_val - min_val

while pq:
    v, i = heapq.heappop(pq)

    if indices[i] == M - 1:
        break

    indices[i] += 1
    heapq.heappush(pq, (classes[i][indices[i]], i))
    min_val = pq[0][0]
    if max_val < classes[i][indices[i]]:
        max_val = classes[i][indices[i]]

    answer = min(answer, max_val - min_val)

print(answer)
