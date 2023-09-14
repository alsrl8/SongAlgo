import sys

T = int(sys.stdin.readline())

fibonacci = [0] * 41
fibonacci[1] = 1
for i in range(2, 41):
    fibonacci[i] = fibonacci[i - 1] + fibonacci[i - 2]

for _ in range(T):
    N = int(sys.stdin.readline())
    if N == 0:
        print(1, 0)
    elif N == 1:
        print(0, 1)
    else:
        print(fibonacci[N - 1], fibonacci[N])
