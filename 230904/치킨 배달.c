#include <stdio.h>
#include <stdlib.h>
#define MAX 100
#define MIN(a, b) ((a < b) ? a : b)

typedef struct{
    int row, col;
}Location;

int N, M;
int map[50][50];
Location cr[13]; // Chicken Restaurant
Location house[100];
int cntCR, cntHouse;
int minSum = 1300;
int picked[13];

int get_sum_of_chicken_distance(int *picked, int cntPicked){
    int temp, sum = 0;
    for (int i = 0; i < cntHouse; i++){
        temp = MAX;
        for (int j = 0; j < cntPicked; j++)
            temp = MIN(temp, abs(cr[picked[j]].row - house[i].row) + abs(cr[picked[j]].col - house[i].col));
        sum += temp;
    }
    return sum;
}

void pick(int *picked, int cnt, int idx){
    int temp;
    if (cnt == M){
        temp = get_sum_of_chicken_distance(picked, cnt);
        if (temp < minSum) minSum = temp;
        return;
    }
    if (idx == cntCR) return;

    picked[cnt] = idx;
    pick(picked, cnt + 1, idx + 1);

    pick(picked, cnt, idx + 1);
}

int main(){
    scanf("%d %d", &N, &M);
    for (int i = 0; i < N; i++)
        for (int j = 0; j < N; j++){
            scanf("%d", &map[i][j]);
            if (map[i][j] == 0) continue;
            else if (map[i][j] == 1){
                house[cntHouse].row = i;
                house[cntHouse].col = j;
                cntHouse++;
            }
            else{
                cr[cntCR].row = i;
                cr[cntCR].col = j;
                cntCR++;
            }
        }
    pick(picked, 0, 0);
    printf("%d", minSum);
    return 0;    
}