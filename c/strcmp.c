#include <stdio.h>
#include <string.h>

int main(int argc, char const *argv[])
{
    char str1[5], str2[5];
    printf("Enter the first string: ");
    scanf("%s", str1);

    printf("Enter the second string: ");
    scanf("%s", str2);

    if (strcmp(str1, str2) == 0)
        printf("Both strings are equal\n");
    else {
        printf("Both strings are not equal\n");

        strcat(str1, str2);
        printf("If we join them, the result will be: %s\n", str1);
    }

    return 0;
}
