import sys
from collections import Counter

N = int(sys.stdin.readline())
A = list(map(int, sys.stdin.readline().split()))

if N < 3:
    print(0)
    exit(0)

counter = Counter(A)  # A의 원소별 등장 횟수를 Dictionary에 저장

answer = 0
for i, tgt in enumerate(A):
    for j, x in enumerate(A):
        if i == j:
            continue
        y = tgt - x
        if y in counter:
            if y == x:
                if tgt == x:
                    if counter[x] == 2:
                        continue