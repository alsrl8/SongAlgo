#include <stdio.h>

int N; // 큐의 크기
int M; // 뽑으려는 원소의 개수
int queue[51];
int currentIdx;
int cntMove;

void initQueue(){
    for (int i = 1; i <= N; i++)
        queue[i] = i;
    currentIdx = 1;
}

void goRight(){
    if (cntMove == N)
        return;
    if (currentIdx > N)
        currentIdx = 1;
    if (queue[currentIdx] == 0){
        currentIdx++;
        cntMove++;
        goRight();
    }
}

void pop(){
    queue[currentIdx] = 0;
    cntMove = 0;
    goRight();
}

int cntLeft(int target, int idx, int cnt){
    if (target == queue[idx])
        return cnt;
    
    if (queue[idx] != 0)
        cnt++;
    idx--;
    if (idx <= 0) idx = N;
    return cntLeft(target, idx, cnt);
}

int cntRight(int target, int idx, int cnt){
    if (target == queue[idx])
        return cnt;

    if (queue[idx] != 0)
        cnt++;
    idx++;
    if (idx > N) idx = 0;
    return cntRight(target, idx, cnt);
}

int main(){
    int left, right, target;
    int arrTarget[50];
    int total = 0;

    scanf("%d %d", &N, &M);
    initQueue();

    for (int i = 0; i < M; i++)
        scanf("%d", &arrTarget[i]);

    for (int i = 0; i < M; i++){
        target = arrTarget[i];
        left = cntLeft(target, currentIdx, 0);
        right = cntRight(target, currentIdx, 0);
        total += (left < right ? left : right);

        currentIdx = target;
        pop();
    }
    printf("%d", total);
    return 0;    
}