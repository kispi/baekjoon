#include <stdio.h>

void printStars(int n) {
    for (int i=1; i<=n; i++) {
        printf("*");
    }
    printf("\n");
}

int main() {
    int n;
    scanf("%d", &n);

    for(int i=n; i>=1; i--) {
        printStars(i);
    }
}