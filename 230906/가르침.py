import sys
import re
import string

# 입력
N, K = map(int, sys.stdin.readline().split())
words = [sys.stdin.readline().rstrip() for _ in range(N)]

must_learn_characters = list('acint')  # 시간복잡도와 관련지어 설명할 것

# 최소 단어 개수보다 K가 작으면 어떤 단어도 배울 수 없다.
if K < len(must_learn_characters):
    print(0)
    exit(0)

K -= len(must_learn_characters)

lower_alphabets = string.ascii_lowercase
remain_characters = re.findall(r'[^acint]', lower_alphabets)
remain_characters_set = set(remain_characters)

ch_index_dict = dict()
for i, ch in enumerate(remain_characters):
    ch_index_dict[ch] = i

words_bit = []
for word in words:
    bit = ['0'] * 21
    for ch in word:
        if ch not in ch_index_dict:
            continue
        index = ch_index_dict[ch]
        bit[index] = '1'
    words_bit.append(int(''.join(bit), 2))

answer = 0
for num in range(2 ** 21):
    bit = bin(num)[2:]
    if bit.count('1') > K:
        continue
    bit = '0' * (len(bit) - 21) + bit
    bit = int(bit, 2)

    cnt = 0
    for wb in words_bit:
        if wb & bit < wb:
            continue
        cnt += 1
    answer = max(answer, cnt)

print(answer)