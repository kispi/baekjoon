#include <stdio.h>

void star(int num, int total) {
    for(int i=0; i<(total-num) / 2; i++)
        printf(" ");

    for(int i=0; i<num; i++) {
        printf("*");
    }
    printf("\n");
}

int main() {
    int n;
    scanf("%d", &n);

    for(int i=1; i<=n; i++) {
        star(i * 2 - 1, n * 2 - 1);
    }
}