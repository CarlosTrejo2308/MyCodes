// 1. Dividir el valor entre 2, y alamacenar el reminder a un array
// 2. Repetir hasta que ya no se pueda
// 3. Imprimir el array

#include <stdio.h>
#include <stdlib.h>

int a[12], n, i;

int main(int argc, char const *argv[])
{

    system("clear");
    printf("Ingrese un numero: ");
    scanf("%d", &n);

    for (i = 0; n > 0; i++)
    {
        a[i] = n % 2;
        n = n / 2;
    }
    printf("El numero en binario es: ");

    for (i = i - 1; i >= 0; i--)
    {
        printf("%d", a[i]);
    }

    printf("\n");
    return 0;
}
