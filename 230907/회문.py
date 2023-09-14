# 회문

import sys

# 입력
T = int(sys.stdin.readline())
words = [sys.stdin.readline().rstrip() for _ in range(T)]


def is_palindrome(_word, start, end):
    l, r = start, end
    while l < r:
        if _word[l] != _word[r]:
            return False
        l += 1
        r -= 1
    return True


def is_similar_palindrome(_word):
    l, r = 0, len(_word) - 1
    while l < r:
        if _word[l] != _word[r]:
            if is_palindrome(_word, l + 1, r):
                return True
            elif is_palindrome(_word, l, r - 1):
                return True
            return False
        l += 1
        r -= 1


for word in words:
    # Case 1. 그냥 회문인지
    if is_palindrome(word, 0, len(word) - 1):
        print(0)
    # Case 2. 유사 회문인지
    elif is_similar_palindrome(word):
        print(1)
    else:
        print(2)
