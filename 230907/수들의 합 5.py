# 수들의 합 5

# 입력
N = int(input())

if N == 1:
    print(1)
    exit(0)

l, r = N, N
total = N
answer = 0

while l > 0:
    if total == N:
        answer += 1
        l -= 1
        total += l
    elif total > N:
        total -= r
        r -= 1
    elif total < N:
        l -= 1
        total += l

print(answer)
