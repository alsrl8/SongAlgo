def solution(numbers, target):
    answer = 0
    
    def dfs(i, cur):
        nonlocal answer
        if i == len(numbers):
            if cur == target:
                answer += 1


        dfs(i + 1, cur + numbers[i])
        dfs(i + 1, cur - numbers[i])

    dfs(0, 0)

    return answer
