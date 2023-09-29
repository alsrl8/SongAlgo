S = input()
answer = []

bracket_open = False
i = 0
while i < len(S):
    ch = S[i]
    if ch == '<':
        bracket_open = True
        answer.append(ch)
        i += 1
        continue
    elif ch == '>':
        bracket_open = False
        answer.append(ch)
        i += 1
        continue
    elif bracket_open:
        answer.append(ch)
        i += 1
        continue
    elif ch == ' ':
        answer.append(ch)
        i += 1
        continue

    j = i
    stack = []
    while j < len(S) and S[j] != '<' and S[j] != ' ':
        stack.append(S[j])
        j += 1
    while stack:
        answer.append(stack.pop())
    i = j

print(''.join(answer))
