#include <stdio.h>

int n;

int main(int argc, char const *argv[])
{

    for (int i = 0; i < 10; i++)
    {
        if (i == 5)
        {
            break;
        }

        printf("i: %d\n", i);
    }

    return 0;
}
