#include <stdio.h>

void demo();
void add(int a, int b);

int main(int argc, char const *argv[])
{
    demo();
    add(1, 2);
    return 0;
}

void demo()
{
    int a, b, sum;
    a = 100;
    b = 200;

    sum = a + b;
    printf("sum = %d\n", sum);
    printf("Hello World!\n");
}

void add(int a, int b)
{
    int sum;
    sum = a + b;
    printf("sum = %d\n", sum);
}