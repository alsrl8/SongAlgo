N = int(input())

binary = bin(N)[2:]
answer = 0
for i in range(len(binary) - 1, -1, -1):
    if binary[i] == '0':
        continue
    j = len(binary) - i
    answer += 3 ** (j - 1)
print(answer)
