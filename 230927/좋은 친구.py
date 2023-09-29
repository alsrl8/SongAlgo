import sys
from collections import defaultdict
from bisect import bisect_right

N, K = map(int, sys.stdin.readline().split())

students = defaultdict(list)
for i in range(N):
    name = sys.stdin.readline().rstrip()
    students[len(name)].append(i)

answer = 0
for _len in range(2, 21):
    if len(students[_len]) < 2:
        continue

    rank = students[_len]

    for i, r in enumerate(rank):
        answer += bisect_right(rank, r + K) - 1 - i

print(answer)
