#include <stdio.h>

int main()
{

    int coca;

    printf("Elije una platzi cola (1-4): ");
    scanf("%d", &coca);

    switch (coca)
    {
    case 1:
        printf("Elejiste una platzi cola zero.\n");
        break;
    case 2:
        printf("Elejiste una platzi cola light.\n");
        break;
    case 3:
        printf("Elejiste una platzi cola regular.\n");
        break;
    case 4:
        printf("Elejiste una platzi cola premium.\n");
        break;

    default:
        printf("Elejiste una platzi que no existe. %d\n", coca);
        break;
    }

    return 0;
}