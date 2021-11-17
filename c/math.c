#include <stdio.h>
#include <math.h>

int main(int argc, char const *argv[])
{
    float x;

    printf("Ingrese el valor de x: ");
    scanf("%f", &x);
    printf("El valor de x es: %f\n", x);

    // sin
    printf("El valor de sin(x) es: %f\n", sin(x));

    // cos
    printf("El valor de cos(x) es: %f\n", cos(x));
    return 0;
}
