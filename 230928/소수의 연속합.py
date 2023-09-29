from bisect import bisect_left

N = int(input())

if N == 1:
    print(0)
    exit(0)

prime_number_flag = [True] * (N + 1)
prime_number_flag[1] = False

end = int(N ** 0.5)
for num in range(2, end + 1):
    if not prime_number_flag[num]:
        continue
    i = 2
    while i * num <= N:
        prime_number_flag[i * num] = False
        i += 1

prime_numbers = [num for num in range(2, N + 1) if prime_number_flag[num]]
l, r = 0, 0
_sum = prime_numbers[0]
answer = 0
while l < len(prime_numbers):
    if _sum < N:
        if r == len(prime_numbers) - 1:
            break
        r += 1
        _sum += prime_numbers[r]
    elif _sum > N:
        _sum -= prime_numbers[l]
        l += 1
    else:
        answer += 1
        _sum -= prime_numbers[l]
        l += 1
print(answer)
