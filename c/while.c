#include <stdio.h>

int n;

int main(int argc, char const *argv[])
{

    printf("Where should we start the countdown?");
    scanf("%d", &n);

    while (n >= 0)
    {
        printf("n: %d\n", n);
        n--;
    }
    
    return 0;
}
