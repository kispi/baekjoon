#include <stdio.h>

void printStars(int cur, int length) {
    for (int i=1; i<=length; i++) {
        if (i < cur) {
            printf(" ");
        } else {
            printf("*");
        }
    }
    printf("\n");
}

int main() {
    int n;
    scanf("%d", &n);

    for(int i=n; i>=1; i--) {
        printStars(i, n);
    }
}