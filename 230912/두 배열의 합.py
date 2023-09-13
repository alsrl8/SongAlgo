import sys
from bisect import bisect_left, bisect_right

# 입력
T = int(sys.stdin.readline())
n = int(sys.stdin.readline())
A = list(map(int, sys.stdin.readline().split()))
m = int(sys.stdin.readline())
B = list(map(int, sys.stdin.readline().split()))

# A와 B의 구간합(Prefix Sum)
ps_A = [a for a in A]
ps_B = [b for b in B]


def calculate_prefix_sum(arr):
    for i in range(1, len(arr)):
        arr[i] += arr[i - 1]


calculate_prefix_sum(ps_A)
calculate_prefix_sum(ps_B)

# 모든 경우의 수를 미리 구한다.
cases_A = []
cases_B = []

for i in range(n):
    for j in range(i, n):
        cases_A.append(ps_A[j] - ps_A[i-1] if i > 0 else ps_A[j])

for i in range(m):
    for j in range(i, m):
        cases_B.append(ps_B[j] - ps_B[i-1] if i > 0 else ps_B[j])

answer = 0
cases_B.sort()
for ca in cases_A:
    target = T - ca
    answer += (bisect_right(cases_B, target) - bisect_left(cases_B, target))
print(answer)
