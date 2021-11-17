#include <stdio.h>

int main(int argc, char const *argv[])
{

    printf("Contando del 1 al 10:\n");

    int i = 1;

    contar:
    printf("%d\n", i);
    i++;
    if (i <= 10)
        goto contar;
    
    printf("Terminamos\n");

    return 0;
}
