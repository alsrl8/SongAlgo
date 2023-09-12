A, B = map(int, input().split())
i, j, arr = 1, 0, []
while len(arr) <= B:
    arr.append(i)
    j += 1
    if j == i:
        j = 0
        i += 1

print(sum(arr[A-1:-1]))
