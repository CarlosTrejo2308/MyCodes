#include <stdio.h>
#include <time.h>

int main(int argc, char const *argv[])
{
    time_t begin, end;
    long diff;

    begin = time(NULL);

    for (int i = 0; i < 6000000000; i++)
    {
        /* code */
    }
    
    end = time(NULL);
    diff = difftime(end, begin);
    printf("Tiempo de ejecucion: %ld\n", diff);

    return 0;
}
