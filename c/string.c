#include <stdio.h>
#include <string.h>

char* strrev(char *str) {
      if (! str || ! *str)
            return str;
      for (char *p1_str = str, *p2_str = str + strlen(str) - 1; p2_str > p1_str; ++p1_str, --p2_str) {
            *p1_str ^= *p2_str;
            *p2_str ^= *p1_str;
            *p1_str  ^= *p2_str;
      }
      return str;
}

int main(int argc, char const *argv[])
{
    char string1[60];

    printf("Enter a string: ");
    // scanf("%s", string1); // se come los espacios
    fgets(string1, 60, stdin); // se lee hasta el enter

    printf("The string entered is: %s\n", string1);

    strrev(string1);
    printf("The reversed string is: %s\n", string1);
    return 0;
}
