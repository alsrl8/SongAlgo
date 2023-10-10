import sys
from collections import defaultdict

N, M = map(int, sys.stdin.readline().split())
dic = defaultdict(int)  # 단어가 등장한 횟수를 저장하는 Dictionary
for _ in range(N):
    word = sys.stdin.readline().rstrip()
    if len(word) < M:
        continue
    dic[word] += 1

result = [(word, cnt) for word, cnt in dic.items()]
result.sort(key=lambda x: (-x[1], -len(x[0]), x[0]))
for word, cnt in result:
    print(word)
