import sys
from collections import defaultdict

# 입력
N, K = map(int, sys.stdin.readline().split())
num = sys.stdin.readline().rstrip()

# 각 숫자마다 인덱스를 역순으로 저장
indices = defaultdict(list)
for i in range(N - 1, -1, -1):
    indices[int(num[i])].append(i)


def get_index_of_biggest(start, end):
    biggest_num, biggest_num_index = -1, -1
    for n in range(9, -1, -1):
        while indices[n] and indices[n][-1] < start:
            indices[n].pop()
        if not indices[n]:
            continue
        elif n <= biggest_num:
            continue
        elif indices[n][-1] > end:
            continue
        biggest_num = n
        biggest_num_index = indices[n][-1]
    return biggest_num_index


answer = []
index_of_biggest = -1
for i in range(N - K):
    index_of_biggest = get_index_of_biggest(index_of_biggest + 1, K + i)
    answer.append(num[index_of_biggest])

print(''.join(answer))
