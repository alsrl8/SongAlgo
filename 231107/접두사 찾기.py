import sys

input = sys.stdin.readline

n, m = map(int, input().split())
words = []
prefix = []

for _ in range(n):
    words.append(input().rstrip())

for _ in range(m):
    prefix.append(input().rstrip())

count = 0

comb = set()
for word in words:
    for i in range(0, len(word)):
        comb.add(word[:i+1])

for item in prefix:
    if item in comb:
        count += 1

print(count)
