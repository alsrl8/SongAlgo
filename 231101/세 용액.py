import sys
from bisect import bisect_left

N = int(sys.stdin.readline())
arr = list(map(int, sys.stdin.readline().split()))
arr.sort()

answer = [arr[0], arr[1], arr[2]]
abs_answer = abs(sum(answer))

for i in range(N):
    j = i + 1
    k = N - 1
    while j < k:
        _sum = arr[i] + arr[j] + arr[k]
        if abs(_sum) < abs_answer:
            abs_answer = abs(_sum)
            answer = [arr[i], arr[j], arr[k]]

        if _sum < 0:
            j += 1
        elif _sum > 0:
            k -= 1
        else:
            break

answer.sort()
print(answer[0], answer[1], answer[2])
