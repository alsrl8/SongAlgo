def solution(skill, skill_trees):
    pre_skill_cnt = dict()
    for i, ch in enumerate(skill):
        pre_skill_cnt[ch] = i

    answer = len(skill_trees)
    for s in skill_trees:
        acquired_skill_num = 0
        for ch in s:


            elif acquired_skill_num < pre_skill_cnt[ch]:
                answer -= 1
                break
            acquired_skill_num += 1

    return answer
