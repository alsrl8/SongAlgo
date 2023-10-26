import sys
from collections import defaultdict

N = int(sys.stdin.readline())
dic = defaultdict(int)
for _ in range(N):
    num = sys.stdin.readline().rstrip()
    dic[num] += 1

max_cnt = 0
max_cnt_num = ''
for num, cnt in dic.items():
    if max_cnt < cnt:
        max_cnt = cnt
        max_cnt_num = num
    elif max_cnt == cnt:
        if int(num) < int(max_cnt_num):
            max_cnt_num = num
print(max_cnt_num)
