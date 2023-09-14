def solution(answers):
    points = [0, 0, 0]
    pattern_2 = [1, 3, 4, 5]
    pattern_3 = [3, 1, 2, 4, 5]

    for i, ans in enumerate(answers):
        # 1번 수포자가 점수를 얻는 경우
        if (i % 5) + 1 == ans:


        # 2번 수포자가 점수를 얻는 경우
        if i % 2 == 0:
            if ans == 2:
                points[1] += 1
        else:
            index = (i % 8) // 2
            if pattern_2[index] == ans:
                points[1] += 1

        # 3번 수포자가 점수를 얻는 경우
        index = (i % 10) // 2
        if pattern_3[index] == ans:
            points[2] += 1

    print(f'{points=}')

    max_point = max(points)
    return [i + 1 for i in range(3) if points[i] == max_point]