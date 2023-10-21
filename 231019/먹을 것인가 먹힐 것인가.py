import sys

T = int(sys.stdin.readline())
for _ in range(T):
    N, M = map(int, sys.stdin.readline().split())
    A = list(map(int, sys.stdin.readline().split()))
    B = list(map(int, sys.stdin.readline().split()))

    A.sort()
    B.sort()
    a, b = 0, 0

    answer = 0
    while a < N:
        while b < M and B[b] < A[a]:
            b += 1
        answer += b
        a += 1
    print(answer)
