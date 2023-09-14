def solution(numbers, hand):
    left_hands = [3, 0]
    right_hands = [3, 2]

    answer = []
    for number in numbers:
        if number in [1, 4, 7]:
            left_hands = [[1, 4, 7].index(number), 0]


            right_hands = [[3, 6, 9].index(number), 2]
            answer.append('R')
        else:
            tgt = [[2, 5, 8, 0].index(number), 1]
            dis_left = abs(left_hands[0] - tgt[0]) + abs(left_hands[1] - tgt[1])
            dis_right = abs(right_hands[0] - tgt[0]) + abs(right_hands[1] - tgt[1])
            if dis_left < dis_right:
                left_hands = tgt
                answer.append('L')
            elif dis_left > dis_right:
                right_hands = tgt
                answer.append('R')
            else:
                if hand == 'left':
                    left_hands = tgt
                    answer.append('L')
                else:
                    right_hands = tgt
                    answer.append('R')

    return ''.join(answer)