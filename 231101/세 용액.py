import sys
from bisect import bisect_left

N = int(sys.stdin.readline())
arr = list(map(int, sys.stdin.readline().split()))
arr.sort()

answer = [arr[0], arr[1], arr[2]]
abs_answer = abs(sum(answer))

for i in range(N):