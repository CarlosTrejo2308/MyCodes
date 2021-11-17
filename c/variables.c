#include <stdio.h>

// Variable Declarations
extern int a, b, c; // x is an external variable, to be used in all the project
float f, g, h;      // y is a local variable, to be used only in this file

// Variable definitions
int a, b, c;

int main(void)
{

    a = 10;
    b = 20;

    c = a + b;
    printf("The sum is: %d\n", c);

    f = 1.0;
    g = 2.0;
    h = f + g;
    printf("The sum is: %f\n", h);
    return 0;
}