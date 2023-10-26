import sys

H, W = map(int, sys.stdin.readline().split())
blocks = list(map(int, sys.stdin.readline().split()))

start = 0
while start < W and blocks[start] == 0:
    start += 1
end = W - 1
while end >= 0 and blocks[end] == 0:
    end -= 1
if end < 0 or start == end:
    print(0)
    exit(0)

# 탐색 범위: [start, end]

answer = 0
i = start
while i <= end:
    h1 = blocks[i]

    stack = []
    j = i + 1
    while j <= end and blocks[j] < h1:
        stack.append(blocks[j])
        j += 1
    if j <= end:
        stack.append(blocks[j])

    # 역으로 stack을 탐색
    while stack:
        h2 = stack.pop()
        min_h = min(h1, h2)
        while stack and stack[-1] < h2:
            h = stack.pop()
            answer += min_h - h
    i = j

print(answer)
