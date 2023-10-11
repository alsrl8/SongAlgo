from collections import defaultdict, deque


def solution(number: str, k: int):
    num_dict = defaultdict(deque)
    for i, num in enumerate(number):
        num_dict[int(num)].append(i)



    def search(_start, _end):
        for num in range(9, -1, -1):
            while num_dict[num] and num_dict[num][0] < _start:
                num_dict[num].popleft()
            if not num_dict[num]:
                continue
            elif num_dict[num][0] > _end:
                continue
            i = num_dict[num].popleft()
            answer.append(number[i])
            return i

    start = 0
    end = k
    for i in range(len(number) - k):
        start = search(start, end) + 1
        end += 1

    return ''.join(answer)