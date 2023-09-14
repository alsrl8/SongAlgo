def solution(triangle):
    for r in range(1, len(triangle)):
        arr = [0 for _ in range(len(triangle[r]))]
        arr[0] = triangle[r - 1][0] + triangle[r][0]
        for c in range(1, len(triangle[r]) - 1):
            arr[c] = max(triangle[r - 1][c - 1], triangle[r - 1][c]) + triangle[r][c]
        arr[-1] = triangle[r - 1][-1] + triangle[r][-1]


    return max(triangle[-1])
