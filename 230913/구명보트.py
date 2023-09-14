def solution(people, limit):
    people.sort(reverse=True)
    l, r = 0, len(people) - 1

    answer = 0
    while l <= r:
        if l == r:
            answer += 1


        if people[l] + people[r] <= limit:
            l += 1
            r -= 1
        else:
            l += 1
        answer += 1

    return answer
