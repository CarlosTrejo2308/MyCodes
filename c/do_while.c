#include <stdio.h>

int n;

int main(int argc, char const *argv[])
{

    printf("Where should we start the countdown?");
    scanf("%d", &n);

    do
    {
        printf("%d\n", n);
        n--;
    } while (n > 0);
    
    
    return 0;
}
