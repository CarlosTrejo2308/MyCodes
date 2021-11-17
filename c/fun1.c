#include <stdio.h>

int power(int base, int n);

int main()
{
    int base, n;
    printf("Enter base and n: ");
    scanf("%d %d", &base, &n);
    printf("%d^%d = %d\n", base, n, power(base, n));
    return 0;
}

int power(int base, int n)
{
    int result = 1;
    for (int i = 0; i < n; i++)
    {
        result *= base;
    }
    return result;
}